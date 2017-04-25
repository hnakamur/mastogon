package dbmodel

import (
	"time"

	kallax "gopkg.in/src-d/go-kallax.v1"
)

//go:generate kallax gen
type Account struct {
	kallax.Model          `table:"accounts"`
	ID                    int64      `pk:"autoincr" json:"id"`
	CreatedAt             time.Time  `json:"created_at"`
	UpdatedAt             time.Time  `json:"updated_at"`
	Username              string     `json:"username"`
	Domain                *string    `json:"domain"`
	Secret                string     `json:"secret"`
	PrivateKey            *string    `json:"private_key"`
	PublicKey             string     `json:"public_key"`
	RemoteURL             string     `json:"remote_key"`
	SalmonURL             string     `json:"salmon_url"`
	HubURL                string     `json:"hub_url"`
	Note                  string     `json:"note"`
	DisplayName           string     `json:"display_name"`
	URI                   string     `json:"uri"`
	URL                   *string    `json:"url"`
	AvatarFileName        *string    `json:"avatar_file_name"`
	AvatarContentType     *string    `json:"avatar_content_type"`
	AvatarFileSize        *int       `json:"avatar_file_size"`
	AvatarUpdatedAt       *time.Time `json:"avatar_updated_at"`
	AvatarRemoteURL       *string    `json:"avatar_remote_url"`
	HeaderFileName        *string    `json:"header_file_name"`
	HeaderContentType     *string    `json:"header_content_type"`
	HeaderFileSize        *int       `json:"header_file_size"`
	HeaderUpdatedAt       *time.Time `json:"header_updated_at"`
	HeaderRemoteURL       string     `json:"header_remote_url"`
	SubscriptionExpiresAt *time.Time `json:"subscription_expires_at"`
	Silenced              bool       `json:"silenced"`
	Suspended             bool       `json:"suspended"`
	Locked                bool       `json:"locked"`
	StatusesCount         int        `json:"statuses_count"`
	FollowersCount        int        `json:"followers_count"`
	LastWebfingeredAt     *time.Time `json:"last_webfingered_at"`
}

type User struct {
	kallax.Model           `table:"users"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
	ID                     int64      `pk:"autoincr" json:"id"`
	AccountID              int        `json:"account_id"`
	EncryptedPassword      string     `json:"encrypted_password"`
	ResetPasswordToken     *string    `json:"reset_password_token"`
	ResetPasswordSentAt    *time.Time `json:"reset_password_sent_at"`
	RememberCreatedAt      *time.Time `json:"remember_created_at"`
	SignInCount            int        `json:"sign_in_count"`
	CurrentSignInAt        *time.Time `json:"current_sign_in_at"`
	LastSignInAt           *time.Time `json:"last_sign_in_at"`
	CurrentSignInIP        *string    `json:"current_sign_in_ip"` // column type is inet
	LastSignInIP           *string    `json:"last_sign_in_ip"`    // column type is inet
	Admin                  *bool      `json:"admin"`
	ConfirmationToken      *string    `json:"confirmation_token"`
	ConfirmedAt            *time.Time `json:"confirmed_at"`
	ConfirmationSentAt     *time.Time `json:"confirmation_sent_at"`
	UnconfirmedEmail       *string    `json:"unconfirmed_email"`
	Locale                 *string    `json:"locale"`
	EncryptedOtpSecret     *string    `json:"encrypted_otp_secret"`
	EncryptedOtpSecretSalt *string    `json:"encrypted_otp_secret_salt"`
	ConsumedTimestep       *int       `json:"consumed_timestep"`
	OtpRequiredForLogin    *bool      `json:"otp_required_for_Login"`
	LastEmailedAt          *time.Time `json:"last_emailed_at"`
	OtpBackupCodes         []string   `json:"otp_backup_codes"`
}

type Status struct {
	kallax.Model       `table:"statuses"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ID                 int64     `pk:"autoincr" json:"id"`
	URI                *string   `json:"uri"`
	AccountID          int       `json:"account_id"`
	Text               string    `json:"text"`
	InReplyToID        *int64    `json:"in_reply_to_id"`
	ReblogOfID         *int64    `json:"reblog_of_id"`
	URL                *string   `json:"url"`
	Sensitive          *bool     `json:"sensitive"`
	Visibility         *int      `json:"visibility"`
	InReplyToAccountID *int      `json:"in_reply_to_account_id"`
	ApplicationID      *int      `json:"application_id"`
	SpoilerText        string    `json:"spoiler_text"`
	Reply              *bool     `json:"reply"`
	FavouritesCount    int       `json:"favourites_count"`
	ReblogsCount       int       `json:"reblogs_count"`
	Language           string    `json:"language"`
}
