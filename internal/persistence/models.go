package persistence

import "time"

type SmsStatus string

var (
	SMS_PENDING  SmsStatus = "pending"
	SMS_ASSIGNED SmsStatus = "assigned"
	SMS_SENT     SmsStatus = "sent"
	SMS_FAILED   SmsStatus = "failed"
)

type SmsJob struct {
	ID         string     `db:"id"`
	Recipient  string     `db:"recipient"`
	Message    string     `db:"message"`
	Status     SmsStatus  `db:"status"`
	RetryCount int        `db:"retry_count"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	SentAt     *time.Time `db:"sent_at"`
	FailedAt   *time.Time `db:"failed_at"`
	LastError  *string    `db:"last_error"`
}

type NewSmsJob struct {
	ID        string `db:"id"`
	Recipient string `db:"recipient"`
	Message   string `db:"message"`
}
