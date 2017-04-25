package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"os"

	"github.com/go-redis/redis"
	"github.com/hnakamur/mastogon/dbmodel"
	"github.com/jrallison/go-workers"
	"github.com/pkg/errors"
)

func findAccountByID(db *sql.DB, accountID int64) (*dbmodel.Account, error) {
	store := dbmodel.NewAccountStore(db)
	q := dbmodel.NewAccountQuery().FindByID(accountID)
	account, err := store.FindOne(q)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return account, nil
}

func findStatusByID(db *sql.DB, accountID int64) (*dbmodel.Status, error) {
	store := dbmodel.NewStatusStore(db)
	q := dbmodel.NewStatusQuery().FindByID(accountID)
	status, err := store.FindOne(q)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return status, nil
}

func pushUpdate(message *workers.Msg) {
	jobID := message.Jid()
	args := message.Args()
	accountID := args.Get("account_id").MustInt64()
	statusID := args.Get("status_id").MustInt64()
	workers.Logger.Printf("pushUpdate, jobID=%s, accountID=%d, statusID=%d", jobID, accountID, statusID)
	account, err := findAccountByID(db, accountID)
	if err != nil {
		panic(err)
	}
	status, err := findStatusByID(db, statusID)
	if err != nil {
		panic(err)
	}
	workers.Logger.Printf("pushUpdate, jobID=%s, account=%+v, status=%+v", message.Jid(), account, status)

	// TODO: Fix payload
	payload := struct {
		Account *dbmodel.Account `json:"account"`
		Status  *dbmodel.Status  `json:"status"`
	}{
		Account: account,
		Status:  status,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	ev := struct {
		Event   string `json:"event"`
		Payload string `json:"payload"`
	}{
		Event:   "update",
		Payload: string(payloadBytes),
	}
	b, err := json.Marshal(ev)
	if err != nil {
		panic(err)
	}
	// TODO: Choose proper timeline queue instead of hardcoding 1
	_, err = redisCli.Publish("timeline:1", string(b)).Result()
	if err != nil {
		panic(err)
	}
	workers.Logger.Printf("pushUpdate, jobID=%s finised", jobID)
}

// TODO: make a struct for app and move these global variables to fields in struct.
var redisCli *redis.Client
var db *sql.DB

func main() {
	var dbURL string
	flag.StringVar(&dbURL, "db-url", os.Getenv("DB_URL"), "postgresql database url")
	flag.Parse()

	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	redisCli = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,
	})
	if err != nil {
		panic(err)
	}
	defer redisCli.Close()

	workers.Configure(map[string]string{
		// location of redis instance
		"server": "localhost:6379",
		// instance of the database
		"database": "1",
		// number of connections to keep open with redis
		"pool": "30",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})

	// pull messages from "myqueue" with concurrency of 10
	workers.Process("myqueue", pushUpdate, 10)

	workers.Enqueue("myqueue", "PushUpdate", struct {
		AccountID int64 `json:"account_id"`
		StatusID  int64 `json:"status_id"`
	}{1, 2})

	// stats will be available at http://localhost:8080/stats
	go workers.StatsServer(8080)

	// Blocks until process is told to exit via unix signal
	workers.Run()
}
