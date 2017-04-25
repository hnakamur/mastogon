// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package dbmodel

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewAccount returns a new instance of Account.
func NewAccount() (record *Account) {
	return new(Account)
}

// GetID returns the primary key of the model.
func (r *Account) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Account) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.CreatedAt, nil
	case "updated_at":
		return &r.UpdatedAt, nil
	case "username":
		return &r.Username, nil
	case "domain":
		return types.Nullable(&r.Domain), nil
	case "secret":
		return &r.Secret, nil
	case "private_key":
		return types.Nullable(&r.PrivateKey), nil
	case "public_key":
		return &r.PublicKey, nil
	case "remote_url":
		return &r.RemoteURL, nil
	case "salmon_url":
		return &r.SalmonURL, nil
	case "hub_url":
		return &r.HubURL, nil
	case "note":
		return &r.Note, nil
	case "display_name":
		return &r.DisplayName, nil
	case "uri":
		return &r.URI, nil
	case "url":
		return types.Nullable(&r.URL), nil
	case "avatar_file_name":
		return types.Nullable(&r.AvatarFileName), nil
	case "avatar_content_type":
		return types.Nullable(&r.AvatarContentType), nil
	case "avatar_file_size":
		return types.Nullable(&r.AvatarFileSize), nil
	case "avatar_updated_at":
		return types.Nullable(&r.AvatarUpdatedAt), nil
	case "avatar_remote_url":
		return types.Nullable(&r.AvatarRemoteURL), nil
	case "header_file_name":
		return types.Nullable(&r.HeaderFileName), nil
	case "header_content_type":
		return types.Nullable(&r.HeaderContentType), nil
	case "header_file_size":
		return types.Nullable(&r.HeaderFileSize), nil
	case "header_updated_at":
		return types.Nullable(&r.HeaderUpdatedAt), nil
	case "header_remote_url":
		return &r.HeaderRemoteURL, nil
	case "subscription_expires_at":
		return types.Nullable(&r.SubscriptionExpiresAt), nil
	case "silenced":
		return &r.Silenced, nil
	case "suspended":
		return &r.Suspended, nil
	case "locked":
		return &r.Locked, nil
	case "statuses_count":
		return &r.StatusesCount, nil
	case "followers_count":
		return &r.FollowersCount, nil
	case "last_webfingered_at":
		return types.Nullable(&r.LastWebfingeredAt), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Account: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Account) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.CreatedAt, nil
	case "updated_at":
		return r.UpdatedAt, nil
	case "username":
		return r.Username, nil
	case "domain":
		if r.Domain == (*string)(nil) {
			return nil, nil
		}
		return r.Domain, nil
	case "secret":
		return r.Secret, nil
	case "private_key":
		if r.PrivateKey == (*string)(nil) {
			return nil, nil
		}
		return r.PrivateKey, nil
	case "public_key":
		return r.PublicKey, nil
	case "remote_url":
		return r.RemoteURL, nil
	case "salmon_url":
		return r.SalmonURL, nil
	case "hub_url":
		return r.HubURL, nil
	case "note":
		return r.Note, nil
	case "display_name":
		return r.DisplayName, nil
	case "uri":
		return r.URI, nil
	case "url":
		if r.URL == (*string)(nil) {
			return nil, nil
		}
		return r.URL, nil
	case "avatar_file_name":
		if r.AvatarFileName == (*string)(nil) {
			return nil, nil
		}
		return r.AvatarFileName, nil
	case "avatar_content_type":
		if r.AvatarContentType == (*string)(nil) {
			return nil, nil
		}
		return r.AvatarContentType, nil
	case "avatar_file_size":
		if r.AvatarFileSize == (*int)(nil) {
			return nil, nil
		}
		return r.AvatarFileSize, nil
	case "avatar_updated_at":
		if r.AvatarUpdatedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.AvatarUpdatedAt, nil
	case "avatar_remote_url":
		if r.AvatarRemoteURL == (*string)(nil) {
			return nil, nil
		}
		return r.AvatarRemoteURL, nil
	case "header_file_name":
		if r.HeaderFileName == (*string)(nil) {
			return nil, nil
		}
		return r.HeaderFileName, nil
	case "header_content_type":
		if r.HeaderContentType == (*string)(nil) {
			return nil, nil
		}
		return r.HeaderContentType, nil
	case "header_file_size":
		if r.HeaderFileSize == (*int)(nil) {
			return nil, nil
		}
		return r.HeaderFileSize, nil
	case "header_updated_at":
		if r.HeaderUpdatedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.HeaderUpdatedAt, nil
	case "header_remote_url":
		return r.HeaderRemoteURL, nil
	case "subscription_expires_at":
		if r.SubscriptionExpiresAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.SubscriptionExpiresAt, nil
	case "silenced":
		return r.Silenced, nil
	case "suspended":
		return r.Suspended, nil
	case "locked":
		return r.Locked, nil
	case "statuses_count":
		return r.StatusesCount, nil
	case "followers_count":
		return r.FollowersCount, nil
	case "last_webfingered_at":
		if r.LastWebfingeredAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.LastWebfingeredAt, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Account: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Account) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Account has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Account) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Account has no relationships")
}

// AccountStore is the entity to access the records of the type Account
// in the database.
type AccountStore struct {
	*kallax.Store
}

// NewAccountStore creates a new instance of AccountStore
// using a SQL database.
func NewAccountStore(db *sql.DB) *AccountStore {
	return &AccountStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *AccountStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *AccountStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Account in the database. A non-persisted object is
// required for this operation.
func (s *AccountStore) Insert(record *Account) error {

	return s.Store.Insert(Schema.Account.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *AccountStore) Update(record *Account, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Account.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *AccountStore) Save(record *Account) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *AccountStore) Delete(record *Account) error {

	return s.Store.Delete(Schema.Account.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *AccountStore) Find(q *AccountQuery) (*AccountResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewAccountResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *AccountStore) MustFind(q *AccountQuery) *AccountResultSet {
	return NewAccountResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *AccountStore) Count(q *AccountQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *AccountStore) MustCount(q *AccountQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *AccountStore) FindOne(q *AccountQuery) (*Account, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *AccountStore) MustFindOne(q *AccountQuery) *Account {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Account with the data in the database and
// makes it writable.
func (s *AccountStore) Reload(record *Account) error {
	return s.Store.Reload(Schema.Account.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *AccountStore) Transaction(callback func(*AccountStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&AccountStore{store})
	})
}

// AccountQuery is the object used to create queries for the Account
// entity.
type AccountQuery struct {
	*kallax.BaseQuery
}

// NewAccountQuery returns a new instance of AccountQuery.
func NewAccountQuery() *AccountQuery {
	return &AccountQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Account.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *AccountQuery) Select(columns ...kallax.SchemaField) *AccountQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *AccountQuery) SelectNot(columns ...kallax.SchemaField) *AccountQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *AccountQuery) Copy() *AccountQuery {
	return &AccountQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *AccountQuery) Order(cols ...kallax.ColumnOrder) *AccountQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *AccountQuery) BatchSize(size uint64) *AccountQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *AccountQuery) Limit(n uint64) *AccountQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *AccountQuery) Offset(n uint64) *AccountQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *AccountQuery) Where(cond kallax.Condition) *AccountQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *AccountQuery) FindByID(v ...int64) *AccountQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Account.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *AccountQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *AccountQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.UpdatedAt, v))
}

// FindByUsername adds a new filter to the query that will require that
// the Username property is equal to the passed value.
func (q *AccountQuery) FindByUsername(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Username, v))
}

// FindBySecret adds a new filter to the query that will require that
// the Secret property is equal to the passed value.
func (q *AccountQuery) FindBySecret(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Secret, v))
}

// FindByPublicKey adds a new filter to the query that will require that
// the PublicKey property is equal to the passed value.
func (q *AccountQuery) FindByPublicKey(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.PublicKey, v))
}

// FindByRemoteURL adds a new filter to the query that will require that
// the RemoteURL property is equal to the passed value.
func (q *AccountQuery) FindByRemoteURL(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.RemoteURL, v))
}

// FindBySalmonURL adds a new filter to the query that will require that
// the SalmonURL property is equal to the passed value.
func (q *AccountQuery) FindBySalmonURL(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.SalmonURL, v))
}

// FindByHubURL adds a new filter to the query that will require that
// the HubURL property is equal to the passed value.
func (q *AccountQuery) FindByHubURL(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.HubURL, v))
}

// FindByNote adds a new filter to the query that will require that
// the Note property is equal to the passed value.
func (q *AccountQuery) FindByNote(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Note, v))
}

// FindByDisplayName adds a new filter to the query that will require that
// the DisplayName property is equal to the passed value.
func (q *AccountQuery) FindByDisplayName(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.DisplayName, v))
}

// FindByURI adds a new filter to the query that will require that
// the URI property is equal to the passed value.
func (q *AccountQuery) FindByURI(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.URI, v))
}

// FindByAvatarUpdatedAt adds a new filter to the query that will require that
// the AvatarUpdatedAt property is equal to the passed value.
func (q *AccountQuery) FindByAvatarUpdatedAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.AvatarUpdatedAt, v))
}

// FindByHeaderUpdatedAt adds a new filter to the query that will require that
// the HeaderUpdatedAt property is equal to the passed value.
func (q *AccountQuery) FindByHeaderUpdatedAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.HeaderUpdatedAt, v))
}

// FindByHeaderRemoteURL adds a new filter to the query that will require that
// the HeaderRemoteURL property is equal to the passed value.
func (q *AccountQuery) FindByHeaderRemoteURL(v string) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.HeaderRemoteURL, v))
}

// FindBySubscriptionExpiresAt adds a new filter to the query that will require that
// the SubscriptionExpiresAt property is equal to the passed value.
func (q *AccountQuery) FindBySubscriptionExpiresAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.SubscriptionExpiresAt, v))
}

// FindBySilenced adds a new filter to the query that will require that
// the Silenced property is equal to the passed value.
func (q *AccountQuery) FindBySilenced(v bool) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Silenced, v))
}

// FindBySuspended adds a new filter to the query that will require that
// the Suspended property is equal to the passed value.
func (q *AccountQuery) FindBySuspended(v bool) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Suspended, v))
}

// FindByLocked adds a new filter to the query that will require that
// the Locked property is equal to the passed value.
func (q *AccountQuery) FindByLocked(v bool) *AccountQuery {
	return q.Where(kallax.Eq(Schema.Account.Locked, v))
}

// FindByStatusesCount adds a new filter to the query that will require that
// the StatusesCount property is equal to the passed value.
func (q *AccountQuery) FindByStatusesCount(cond kallax.ScalarCond, v int) *AccountQuery {
	return q.Where(cond(Schema.Account.StatusesCount, v))
}

// FindByFollowersCount adds a new filter to the query that will require that
// the FollowersCount property is equal to the passed value.
func (q *AccountQuery) FindByFollowersCount(cond kallax.ScalarCond, v int) *AccountQuery {
	return q.Where(cond(Schema.Account.FollowersCount, v))
}

// FindByLastWebfingeredAt adds a new filter to the query that will require that
// the LastWebfingeredAt property is equal to the passed value.
func (q *AccountQuery) FindByLastWebfingeredAt(cond kallax.ScalarCond, v time.Time) *AccountQuery {
	return q.Where(cond(Schema.Account.LastWebfingeredAt, v))
}

// AccountResultSet is the set of results returned by a query to the
// database.
type AccountResultSet struct {
	ResultSet kallax.ResultSet
	last      *Account
	lastErr   error
}

// NewAccountResultSet creates a new result set for rows of the type
// Account.
func NewAccountResultSet(rs kallax.ResultSet) *AccountResultSet {
	return &AccountResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *AccountResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Account.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Account)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Account")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *AccountResultSet) Get() (*Account, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *AccountResultSet) ForEach(fn func(*Account) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *AccountResultSet) All() ([]*Account, error) {
	var result []*Account
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *AccountResultSet) One() (*Account, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *AccountResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *AccountResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewStatus returns a new instance of Status.
func NewStatus() (record *Status) {
	return new(Status)
}

// GetID returns the primary key of the model.
func (r *Status) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Status) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.CreatedAt, nil
	case "updated_at":
		return &r.UpdatedAt, nil
	case "uri":
		return types.Nullable(&r.URI), nil
	case "account_id":
		return &r.AccountID, nil
	case "text":
		return &r.Text, nil
	case "in_reply_to_id":
		return types.Nullable(&r.InReplyToID), nil
	case "reblog_of_id":
		return types.Nullable(&r.ReblogOfID), nil
	case "url":
		return types.Nullable(&r.URL), nil
	case "sensitive":
		return types.Nullable(&r.Sensitive), nil
	case "visibility":
		return types.Nullable(&r.Visibility), nil
	case "in_reply_to_account_id":
		return types.Nullable(&r.InReplyToAccountID), nil
	case "application_id":
		return types.Nullable(&r.ApplicationID), nil
	case "spoiler_text":
		return &r.SpoilerText, nil
	case "reply":
		return types.Nullable(&r.Reply), nil
	case "favourites_count":
		return &r.FavouritesCount, nil
	case "reblogs_count":
		return &r.ReblogsCount, nil
	case "language":
		return &r.Language, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Status: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Status) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.CreatedAt, nil
	case "updated_at":
		return r.UpdatedAt, nil
	case "uri":
		if r.URI == (*string)(nil) {
			return nil, nil
		}
		return r.URI, nil
	case "account_id":
		return r.AccountID, nil
	case "text":
		return r.Text, nil
	case "in_reply_to_id":
		if r.InReplyToID == (*int64)(nil) {
			return nil, nil
		}
		return r.InReplyToID, nil
	case "reblog_of_id":
		if r.ReblogOfID == (*int64)(nil) {
			return nil, nil
		}
		return r.ReblogOfID, nil
	case "url":
		if r.URL == (*string)(nil) {
			return nil, nil
		}
		return r.URL, nil
	case "sensitive":
		if r.Sensitive == (*bool)(nil) {
			return nil, nil
		}
		return r.Sensitive, nil
	case "visibility":
		if r.Visibility == (*int)(nil) {
			return nil, nil
		}
		return r.Visibility, nil
	case "in_reply_to_account_id":
		if r.InReplyToAccountID == (*int)(nil) {
			return nil, nil
		}
		return r.InReplyToAccountID, nil
	case "application_id":
		if r.ApplicationID == (*int)(nil) {
			return nil, nil
		}
		return r.ApplicationID, nil
	case "spoiler_text":
		return r.SpoilerText, nil
	case "reply":
		if r.Reply == (*bool)(nil) {
			return nil, nil
		}
		return r.Reply, nil
	case "favourites_count":
		return r.FavouritesCount, nil
	case "reblogs_count":
		return r.ReblogsCount, nil
	case "language":
		return r.Language, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Status: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Status) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Status has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Status) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Status has no relationships")
}

// StatusStore is the entity to access the records of the type Status
// in the database.
type StatusStore struct {
	*kallax.Store
}

// NewStatusStore creates a new instance of StatusStore
// using a SQL database.
func NewStatusStore(db *sql.DB) *StatusStore {
	return &StatusStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *StatusStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *StatusStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a Status in the database. A non-persisted object is
// required for this operation.
func (s *StatusStore) Insert(record *Status) error {

	return s.Store.Insert(Schema.Status.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *StatusStore) Update(record *Status, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.Status.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *StatusStore) Save(record *Status) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *StatusStore) Delete(record *Status) error {

	return s.Store.Delete(Schema.Status.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *StatusStore) Find(q *StatusQuery) (*StatusResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewStatusResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *StatusStore) MustFind(q *StatusQuery) *StatusResultSet {
	return NewStatusResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *StatusStore) Count(q *StatusQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *StatusStore) MustCount(q *StatusQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *StatusStore) FindOne(q *StatusQuery) (*Status, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *StatusStore) MustFindOne(q *StatusQuery) *Status {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Status with the data in the database and
// makes it writable.
func (s *StatusStore) Reload(record *Status) error {
	return s.Store.Reload(Schema.Status.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *StatusStore) Transaction(callback func(*StatusStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&StatusStore{store})
	})
}

// StatusQuery is the object used to create queries for the Status
// entity.
type StatusQuery struct {
	*kallax.BaseQuery
}

// NewStatusQuery returns a new instance of StatusQuery.
func NewStatusQuery() *StatusQuery {
	return &StatusQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Status.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *StatusQuery) Select(columns ...kallax.SchemaField) *StatusQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *StatusQuery) SelectNot(columns ...kallax.SchemaField) *StatusQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *StatusQuery) Copy() *StatusQuery {
	return &StatusQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *StatusQuery) Order(cols ...kallax.ColumnOrder) *StatusQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *StatusQuery) BatchSize(size uint64) *StatusQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *StatusQuery) Limit(n uint64) *StatusQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *StatusQuery) Offset(n uint64) *StatusQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *StatusQuery) Where(cond kallax.Condition) *StatusQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *StatusQuery) FindByID(v ...int64) *StatusQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Status.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *StatusQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *StatusQuery {
	return q.Where(cond(Schema.Status.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *StatusQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *StatusQuery {
	return q.Where(cond(Schema.Status.UpdatedAt, v))
}

// FindByAccountID adds a new filter to the query that will require that
// the AccountID property is equal to the passed value.
func (q *StatusQuery) FindByAccountID(cond kallax.ScalarCond, v int) *StatusQuery {
	return q.Where(cond(Schema.Status.AccountID, v))
}

// FindByText adds a new filter to the query that will require that
// the Text property is equal to the passed value.
func (q *StatusQuery) FindByText(v string) *StatusQuery {
	return q.Where(kallax.Eq(Schema.Status.Text, v))
}

// FindBySpoilerText adds a new filter to the query that will require that
// the SpoilerText property is equal to the passed value.
func (q *StatusQuery) FindBySpoilerText(v string) *StatusQuery {
	return q.Where(kallax.Eq(Schema.Status.SpoilerText, v))
}

// FindByFavouritesCount adds a new filter to the query that will require that
// the FavouritesCount property is equal to the passed value.
func (q *StatusQuery) FindByFavouritesCount(cond kallax.ScalarCond, v int) *StatusQuery {
	return q.Where(cond(Schema.Status.FavouritesCount, v))
}

// FindByReblogsCount adds a new filter to the query that will require that
// the ReblogsCount property is equal to the passed value.
func (q *StatusQuery) FindByReblogsCount(cond kallax.ScalarCond, v int) *StatusQuery {
	return q.Where(cond(Schema.Status.ReblogsCount, v))
}

// FindByLanguage adds a new filter to the query that will require that
// the Language property is equal to the passed value.
func (q *StatusQuery) FindByLanguage(v string) *StatusQuery {
	return q.Where(kallax.Eq(Schema.Status.Language, v))
}

// StatusResultSet is the set of results returned by a query to the
// database.
type StatusResultSet struct {
	ResultSet kallax.ResultSet
	last      *Status
	lastErr   error
}

// NewStatusResultSet creates a new result set for rows of the type
// Status.
func NewStatusResultSet(rs kallax.ResultSet) *StatusResultSet {
	return &StatusResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *StatusResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Status.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Status)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Status")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *StatusResultSet) Get() (*Status, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *StatusResultSet) ForEach(fn func(*Status) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *StatusResultSet) All() ([]*Status, error) {
	var result []*Status
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *StatusResultSet) One() (*Status, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *StatusResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *StatusResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewUser returns a new instance of User.
func NewUser() (record *User) {
	return new(User)
}

// GetID returns the primary key of the model.
func (r *User) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *User) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.CreatedAt, nil
	case "updated_at":
		return &r.UpdatedAt, nil
	case "account_id":
		return &r.AccountID, nil
	case "encrypted_password":
		return &r.EncryptedPassword, nil
	case "reset_password_token":
		return types.Nullable(&r.ResetPasswordToken), nil
	case "reset_password_sent_at":
		return types.Nullable(&r.ResetPasswordSentAt), nil
	case "remember_created_at":
		return types.Nullable(&r.RememberCreatedAt), nil
	case "sign_in_count":
		return &r.SignInCount, nil
	case "current_sign_in_at":
		return types.Nullable(&r.CurrentSignInAt), nil
	case "last_sign_in_at":
		return types.Nullable(&r.LastSignInAt), nil
	case "current_sign_in_ip":
		return types.Nullable(&r.CurrentSignInIP), nil
	case "last_sign_in_ip":
		return types.Nullable(&r.LastSignInIP), nil
	case "admin":
		return types.Nullable(&r.Admin), nil
	case "confirmation_token":
		return types.Nullable(&r.ConfirmationToken), nil
	case "confirmed_at":
		return types.Nullable(&r.ConfirmedAt), nil
	case "confirmation_sent_at":
		return types.Nullable(&r.ConfirmationSentAt), nil
	case "unconfirmed_email":
		return types.Nullable(&r.UnconfirmedEmail), nil
	case "locale":
		return types.Nullable(&r.Locale), nil
	case "encrypted_otp_secret":
		return types.Nullable(&r.EncryptedOtpSecret), nil
	case "encrypted_otp_secret_salt":
		return types.Nullable(&r.EncryptedOtpSecretSalt), nil
	case "consumed_timestep":
		return types.Nullable(&r.ConsumedTimestep), nil
	case "otp_required_for_login":
		return types.Nullable(&r.OtpRequiredForLogin), nil
	case "last_emailed_at":
		return types.Nullable(&r.LastEmailedAt), nil
	case "otp_backup_codes":
		return types.Slice(&r.OtpBackupCodes), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// Value returns the value of the given column.
func (r *User) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.CreatedAt, nil
	case "updated_at":
		return r.UpdatedAt, nil
	case "account_id":
		return r.AccountID, nil
	case "encrypted_password":
		return r.EncryptedPassword, nil
	case "reset_password_token":
		if r.ResetPasswordToken == (*string)(nil) {
			return nil, nil
		}
		return r.ResetPasswordToken, nil
	case "reset_password_sent_at":
		if r.ResetPasswordSentAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.ResetPasswordSentAt, nil
	case "remember_created_at":
		if r.RememberCreatedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.RememberCreatedAt, nil
	case "sign_in_count":
		return r.SignInCount, nil
	case "current_sign_in_at":
		if r.CurrentSignInAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.CurrentSignInAt, nil
	case "last_sign_in_at":
		if r.LastSignInAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.LastSignInAt, nil
	case "current_sign_in_ip":
		if r.CurrentSignInIP == (*string)(nil) {
			return nil, nil
		}
		return r.CurrentSignInIP, nil
	case "last_sign_in_ip":
		if r.LastSignInIP == (*string)(nil) {
			return nil, nil
		}
		return r.LastSignInIP, nil
	case "admin":
		if r.Admin == (*bool)(nil) {
			return nil, nil
		}
		return r.Admin, nil
	case "confirmation_token":
		if r.ConfirmationToken == (*string)(nil) {
			return nil, nil
		}
		return r.ConfirmationToken, nil
	case "confirmed_at":
		if r.ConfirmedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.ConfirmedAt, nil
	case "confirmation_sent_at":
		if r.ConfirmationSentAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.ConfirmationSentAt, nil
	case "unconfirmed_email":
		if r.UnconfirmedEmail == (*string)(nil) {
			return nil, nil
		}
		return r.UnconfirmedEmail, nil
	case "locale":
		if r.Locale == (*string)(nil) {
			return nil, nil
		}
		return r.Locale, nil
	case "encrypted_otp_secret":
		if r.EncryptedOtpSecret == (*string)(nil) {
			return nil, nil
		}
		return r.EncryptedOtpSecret, nil
	case "encrypted_otp_secret_salt":
		if r.EncryptedOtpSecretSalt == (*string)(nil) {
			return nil, nil
		}
		return r.EncryptedOtpSecretSalt, nil
	case "consumed_timestep":
		if r.ConsumedTimestep == (*int)(nil) {
			return nil, nil
		}
		return r.ConsumedTimestep, nil
	case "otp_required_for_login":
		if r.OtpRequiredForLogin == (*bool)(nil) {
			return nil, nil
		}
		return r.OtpRequiredForLogin, nil
	case "last_emailed_at":
		if r.LastEmailedAt == (*time.Time)(nil) {
			return nil, nil
		}
		return r.LastEmailedAt, nil
	case "otp_backup_codes":
		return types.Slice(r.OtpBackupCodes), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in User: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *User) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model User has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *User) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model User has no relationships")
}

// UserStore is the entity to access the records of the type User
// in the database.
type UserStore struct {
	*kallax.Store
}

// NewUserStore creates a new instance of UserStore
// using a SQL database.
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *UserStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *UserStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Insert inserts a User in the database. A non-persisted object is
// required for this operation.
func (s *UserStore) Insert(record *User) error {

	return s.Store.Insert(Schema.User.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *UserStore) Update(record *User, cols ...kallax.SchemaField) (updated int64, err error) {

	return s.Store.Update(Schema.User.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *UserStore) Save(record *User) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *UserStore) Delete(record *User) error {

	return s.Store.Delete(Schema.User.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *UserStore) Find(q *UserQuery) (*UserResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewUserResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *UserStore) MustFind(q *UserQuery) *UserResultSet {
	return NewUserResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *UserStore) Count(q *UserQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *UserStore) MustCount(q *UserQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *UserStore) FindOne(q *UserQuery) (*User, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *UserStore) MustFindOne(q *UserQuery) *User {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the User with the data in the database and
// makes it writable.
func (s *UserStore) Reload(record *User) error {
	return s.Store.Reload(Schema.User.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *UserStore) Transaction(callback func(*UserStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&UserStore{store})
	})
}

// UserQuery is the object used to create queries for the User
// entity.
type UserQuery struct {
	*kallax.BaseQuery
}

// NewUserQuery returns a new instance of UserQuery.
func NewUserQuery() *UserQuery {
	return &UserQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.User.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *UserQuery) Select(columns ...kallax.SchemaField) *UserQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *UserQuery) SelectNot(columns ...kallax.SchemaField) *UserQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *UserQuery) Copy() *UserQuery {
	return &UserQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *UserQuery) Order(cols ...kallax.ColumnOrder) *UserQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *UserQuery) BatchSize(size uint64) *UserQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *UserQuery) Limit(n uint64) *UserQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *UserQuery) Offset(n uint64) *UserQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *UserQuery) Where(cond kallax.Condition) *UserQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByID(v ...int64) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.User.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *UserQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *UserQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.UpdatedAt, v))
}

// FindByAccountID adds a new filter to the query that will require that
// the AccountID property is equal to the passed value.
func (q *UserQuery) FindByAccountID(cond kallax.ScalarCond, v int) *UserQuery {
	return q.Where(cond(Schema.User.AccountID, v))
}

// FindByEncryptedPassword adds a new filter to the query that will require that
// the EncryptedPassword property is equal to the passed value.
func (q *UserQuery) FindByEncryptedPassword(v string) *UserQuery {
	return q.Where(kallax.Eq(Schema.User.EncryptedPassword, v))
}

// FindByResetPasswordSentAt adds a new filter to the query that will require that
// the ResetPasswordSentAt property is equal to the passed value.
func (q *UserQuery) FindByResetPasswordSentAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.ResetPasswordSentAt, v))
}

// FindByRememberCreatedAt adds a new filter to the query that will require that
// the RememberCreatedAt property is equal to the passed value.
func (q *UserQuery) FindByRememberCreatedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.RememberCreatedAt, v))
}

// FindBySignInCount adds a new filter to the query that will require that
// the SignInCount property is equal to the passed value.
func (q *UserQuery) FindBySignInCount(cond kallax.ScalarCond, v int) *UserQuery {
	return q.Where(cond(Schema.User.SignInCount, v))
}

// FindByCurrentSignInAt adds a new filter to the query that will require that
// the CurrentSignInAt property is equal to the passed value.
func (q *UserQuery) FindByCurrentSignInAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.CurrentSignInAt, v))
}

// FindByLastSignInAt adds a new filter to the query that will require that
// the LastSignInAt property is equal to the passed value.
func (q *UserQuery) FindByLastSignInAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.LastSignInAt, v))
}

// FindByConfirmedAt adds a new filter to the query that will require that
// the ConfirmedAt property is equal to the passed value.
func (q *UserQuery) FindByConfirmedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.ConfirmedAt, v))
}

// FindByConfirmationSentAt adds a new filter to the query that will require that
// the ConfirmationSentAt property is equal to the passed value.
func (q *UserQuery) FindByConfirmationSentAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.ConfirmationSentAt, v))
}

// FindByLastEmailedAt adds a new filter to the query that will require that
// the LastEmailedAt property is equal to the passed value.
func (q *UserQuery) FindByLastEmailedAt(cond kallax.ScalarCond, v time.Time) *UserQuery {
	return q.Where(cond(Schema.User.LastEmailedAt, v))
}

// FindByOtpBackupCodes adds a new filter to the query that will require that
// the OtpBackupCodes property contains all the passed values; if no passed values,
// it will do nothing.
func (q *UserQuery) FindByOtpBackupCodes(v ...string) *UserQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.User.OtpBackupCodes, values...))
}

// UserResultSet is the set of results returned by a query to the
// database.
type UserResultSet struct {
	ResultSet kallax.ResultSet
	last      *User
	lastErr   error
}

// NewUserResultSet creates a new result set for rows of the type
// User.
func NewUserResultSet(rs kallax.ResultSet) *UserResultSet {
	return &UserResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *UserResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.User.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*User)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *User")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *UserResultSet) Get() (*User, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *UserResultSet) ForEach(fn func(*User) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *UserResultSet) All() ([]*User, error) {
	var result []*User
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *UserResultSet) One() (*User, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *UserResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *UserResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Account *schemaAccount
	Status  *schemaStatus
	User    *schemaUser
}

type schemaAccount struct {
	*kallax.BaseSchema
	ID                    kallax.SchemaField
	CreatedAt             kallax.SchemaField
	UpdatedAt             kallax.SchemaField
	Username              kallax.SchemaField
	Domain                kallax.SchemaField
	Secret                kallax.SchemaField
	PrivateKey            kallax.SchemaField
	PublicKey             kallax.SchemaField
	RemoteURL             kallax.SchemaField
	SalmonURL             kallax.SchemaField
	HubURL                kallax.SchemaField
	Note                  kallax.SchemaField
	DisplayName           kallax.SchemaField
	URI                   kallax.SchemaField
	URL                   kallax.SchemaField
	AvatarFileName        kallax.SchemaField
	AvatarContentType     kallax.SchemaField
	AvatarFileSize        kallax.SchemaField
	AvatarUpdatedAt       kallax.SchemaField
	AvatarRemoteURL       kallax.SchemaField
	HeaderFileName        kallax.SchemaField
	HeaderContentType     kallax.SchemaField
	HeaderFileSize        kallax.SchemaField
	HeaderUpdatedAt       kallax.SchemaField
	HeaderRemoteURL       kallax.SchemaField
	SubscriptionExpiresAt kallax.SchemaField
	Silenced              kallax.SchemaField
	Suspended             kallax.SchemaField
	Locked                kallax.SchemaField
	StatusesCount         kallax.SchemaField
	FollowersCount        kallax.SchemaField
	LastWebfingeredAt     kallax.SchemaField
}

type schemaStatus struct {
	*kallax.BaseSchema
	ID                 kallax.SchemaField
	CreatedAt          kallax.SchemaField
	UpdatedAt          kallax.SchemaField
	URI                kallax.SchemaField
	AccountID          kallax.SchemaField
	Text               kallax.SchemaField
	InReplyToID        kallax.SchemaField
	ReblogOfID         kallax.SchemaField
	URL                kallax.SchemaField
	Sensitive          kallax.SchemaField
	Visibility         kallax.SchemaField
	InReplyToAccountID kallax.SchemaField
	ApplicationID      kallax.SchemaField
	SpoilerText        kallax.SchemaField
	Reply              kallax.SchemaField
	FavouritesCount    kallax.SchemaField
	ReblogsCount       kallax.SchemaField
	Language           kallax.SchemaField
}

type schemaUser struct {
	*kallax.BaseSchema
	ID                     kallax.SchemaField
	CreatedAt              kallax.SchemaField
	UpdatedAt              kallax.SchemaField
	AccountID              kallax.SchemaField
	EncryptedPassword      kallax.SchemaField
	ResetPasswordToken     kallax.SchemaField
	ResetPasswordSentAt    kallax.SchemaField
	RememberCreatedAt      kallax.SchemaField
	SignInCount            kallax.SchemaField
	CurrentSignInAt        kallax.SchemaField
	LastSignInAt           kallax.SchemaField
	CurrentSignInIP        kallax.SchemaField
	LastSignInIP           kallax.SchemaField
	Admin                  kallax.SchemaField
	ConfirmationToken      kallax.SchemaField
	ConfirmedAt            kallax.SchemaField
	ConfirmationSentAt     kallax.SchemaField
	UnconfirmedEmail       kallax.SchemaField
	Locale                 kallax.SchemaField
	EncryptedOtpSecret     kallax.SchemaField
	EncryptedOtpSecretSalt kallax.SchemaField
	ConsumedTimestep       kallax.SchemaField
	OtpRequiredForLogin    kallax.SchemaField
	LastEmailedAt          kallax.SchemaField
	OtpBackupCodes         kallax.SchemaField
}

var Schema = &schema{
	Account: &schemaAccount{
		BaseSchema: kallax.NewBaseSchema(
			"accounts",
			"__account",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Account)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("username"),
			kallax.NewSchemaField("domain"),
			kallax.NewSchemaField("secret"),
			kallax.NewSchemaField("private_key"),
			kallax.NewSchemaField("public_key"),
			kallax.NewSchemaField("remote_url"),
			kallax.NewSchemaField("salmon_url"),
			kallax.NewSchemaField("hub_url"),
			kallax.NewSchemaField("note"),
			kallax.NewSchemaField("display_name"),
			kallax.NewSchemaField("uri"),
			kallax.NewSchemaField("url"),
			kallax.NewSchemaField("avatar_file_name"),
			kallax.NewSchemaField("avatar_content_type"),
			kallax.NewSchemaField("avatar_file_size"),
			kallax.NewSchemaField("avatar_updated_at"),
			kallax.NewSchemaField("avatar_remote_url"),
			kallax.NewSchemaField("header_file_name"),
			kallax.NewSchemaField("header_content_type"),
			kallax.NewSchemaField("header_file_size"),
			kallax.NewSchemaField("header_updated_at"),
			kallax.NewSchemaField("header_remote_url"),
			kallax.NewSchemaField("subscription_expires_at"),
			kallax.NewSchemaField("silenced"),
			kallax.NewSchemaField("suspended"),
			kallax.NewSchemaField("locked"),
			kallax.NewSchemaField("statuses_count"),
			kallax.NewSchemaField("followers_count"),
			kallax.NewSchemaField("last_webfingered_at"),
		),
		ID:                    kallax.NewSchemaField("id"),
		CreatedAt:             kallax.NewSchemaField("created_at"),
		UpdatedAt:             kallax.NewSchemaField("updated_at"),
		Username:              kallax.NewSchemaField("username"),
		Domain:                kallax.NewSchemaField("domain"),
		Secret:                kallax.NewSchemaField("secret"),
		PrivateKey:            kallax.NewSchemaField("private_key"),
		PublicKey:             kallax.NewSchemaField("public_key"),
		RemoteURL:             kallax.NewSchemaField("remote_url"),
		SalmonURL:             kallax.NewSchemaField("salmon_url"),
		HubURL:                kallax.NewSchemaField("hub_url"),
		Note:                  kallax.NewSchemaField("note"),
		DisplayName:           kallax.NewSchemaField("display_name"),
		URI:                   kallax.NewSchemaField("uri"),
		URL:                   kallax.NewSchemaField("url"),
		AvatarFileName:        kallax.NewSchemaField("avatar_file_name"),
		AvatarContentType:     kallax.NewSchemaField("avatar_content_type"),
		AvatarFileSize:        kallax.NewSchemaField("avatar_file_size"),
		AvatarUpdatedAt:       kallax.NewSchemaField("avatar_updated_at"),
		AvatarRemoteURL:       kallax.NewSchemaField("avatar_remote_url"),
		HeaderFileName:        kallax.NewSchemaField("header_file_name"),
		HeaderContentType:     kallax.NewSchemaField("header_content_type"),
		HeaderFileSize:        kallax.NewSchemaField("header_file_size"),
		HeaderUpdatedAt:       kallax.NewSchemaField("header_updated_at"),
		HeaderRemoteURL:       kallax.NewSchemaField("header_remote_url"),
		SubscriptionExpiresAt: kallax.NewSchemaField("subscription_expires_at"),
		Silenced:              kallax.NewSchemaField("silenced"),
		Suspended:             kallax.NewSchemaField("suspended"),
		Locked:                kallax.NewSchemaField("locked"),
		StatusesCount:         kallax.NewSchemaField("statuses_count"),
		FollowersCount:        kallax.NewSchemaField("followers_count"),
		LastWebfingeredAt:     kallax.NewSchemaField("last_webfingered_at"),
	},
	Status: &schemaStatus{
		BaseSchema: kallax.NewBaseSchema(
			"statuses",
			"__status",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Status)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("uri"),
			kallax.NewSchemaField("account_id"),
			kallax.NewSchemaField("text"),
			kallax.NewSchemaField("in_reply_to_id"),
			kallax.NewSchemaField("reblog_of_id"),
			kallax.NewSchemaField("url"),
			kallax.NewSchemaField("sensitive"),
			kallax.NewSchemaField("visibility"),
			kallax.NewSchemaField("in_reply_to_account_id"),
			kallax.NewSchemaField("application_id"),
			kallax.NewSchemaField("spoiler_text"),
			kallax.NewSchemaField("reply"),
			kallax.NewSchemaField("favourites_count"),
			kallax.NewSchemaField("reblogs_count"),
			kallax.NewSchemaField("language"),
		),
		ID:                 kallax.NewSchemaField("id"),
		CreatedAt:          kallax.NewSchemaField("created_at"),
		UpdatedAt:          kallax.NewSchemaField("updated_at"),
		URI:                kallax.NewSchemaField("uri"),
		AccountID:          kallax.NewSchemaField("account_id"),
		Text:               kallax.NewSchemaField("text"),
		InReplyToID:        kallax.NewSchemaField("in_reply_to_id"),
		ReblogOfID:         kallax.NewSchemaField("reblog_of_id"),
		URL:                kallax.NewSchemaField("url"),
		Sensitive:          kallax.NewSchemaField("sensitive"),
		Visibility:         kallax.NewSchemaField("visibility"),
		InReplyToAccountID: kallax.NewSchemaField("in_reply_to_account_id"),
		ApplicationID:      kallax.NewSchemaField("application_id"),
		SpoilerText:        kallax.NewSchemaField("spoiler_text"),
		Reply:              kallax.NewSchemaField("reply"),
		FavouritesCount:    kallax.NewSchemaField("favourites_count"),
		ReblogsCount:       kallax.NewSchemaField("reblogs_count"),
		Language:           kallax.NewSchemaField("language"),
	},
	User: &schemaUser{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__user",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(User)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("account_id"),
			kallax.NewSchemaField("encrypted_password"),
			kallax.NewSchemaField("reset_password_token"),
			kallax.NewSchemaField("reset_password_sent_at"),
			kallax.NewSchemaField("remember_created_at"),
			kallax.NewSchemaField("sign_in_count"),
			kallax.NewSchemaField("current_sign_in_at"),
			kallax.NewSchemaField("last_sign_in_at"),
			kallax.NewSchemaField("current_sign_in_ip"),
			kallax.NewSchemaField("last_sign_in_ip"),
			kallax.NewSchemaField("admin"),
			kallax.NewSchemaField("confirmation_token"),
			kallax.NewSchemaField("confirmed_at"),
			kallax.NewSchemaField("confirmation_sent_at"),
			kallax.NewSchemaField("unconfirmed_email"),
			kallax.NewSchemaField("locale"),
			kallax.NewSchemaField("encrypted_otp_secret"),
			kallax.NewSchemaField("encrypted_otp_secret_salt"),
			kallax.NewSchemaField("consumed_timestep"),
			kallax.NewSchemaField("otp_required_for_login"),
			kallax.NewSchemaField("last_emailed_at"),
			kallax.NewSchemaField("otp_backup_codes"),
		),
		ID:                     kallax.NewSchemaField("id"),
		CreatedAt:              kallax.NewSchemaField("created_at"),
		UpdatedAt:              kallax.NewSchemaField("updated_at"),
		AccountID:              kallax.NewSchemaField("account_id"),
		EncryptedPassword:      kallax.NewSchemaField("encrypted_password"),
		ResetPasswordToken:     kallax.NewSchemaField("reset_password_token"),
		ResetPasswordSentAt:    kallax.NewSchemaField("reset_password_sent_at"),
		RememberCreatedAt:      kallax.NewSchemaField("remember_created_at"),
		SignInCount:            kallax.NewSchemaField("sign_in_count"),
		CurrentSignInAt:        kallax.NewSchemaField("current_sign_in_at"),
		LastSignInAt:           kallax.NewSchemaField("last_sign_in_at"),
		CurrentSignInIP:        kallax.NewSchemaField("current_sign_in_ip"),
		LastSignInIP:           kallax.NewSchemaField("last_sign_in_ip"),
		Admin:                  kallax.NewSchemaField("admin"),
		ConfirmationToken:      kallax.NewSchemaField("confirmation_token"),
		ConfirmedAt:            kallax.NewSchemaField("confirmed_at"),
		ConfirmationSentAt:     kallax.NewSchemaField("confirmation_sent_at"),
		UnconfirmedEmail:       kallax.NewSchemaField("unconfirmed_email"),
		Locale:                 kallax.NewSchemaField("locale"),
		EncryptedOtpSecret:     kallax.NewSchemaField("encrypted_otp_secret"),
		EncryptedOtpSecretSalt: kallax.NewSchemaField("encrypted_otp_secret_salt"),
		ConsumedTimestep:       kallax.NewSchemaField("consumed_timestep"),
		OtpRequiredForLogin:    kallax.NewSchemaField("otp_required_for_login"),
		LastEmailedAt:          kallax.NewSchemaField("last_emailed_at"),
		OtpBackupCodes:         kallax.NewSchemaField("otp_backup_codes"),
	},
}
