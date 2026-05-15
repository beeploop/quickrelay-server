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
	ID         string     `query:"id"`
	Recipient  string     `query:"recipient"`
	Message    string     `query:"message"`
	Status     SmsStatus  `query:"status"`
	RetryCount int        `query:"retry_count"`
	CreatedAt  time.Time  `query:"created_at"`
	UpdatedAt  time.Time  `query:"updated_at"`
	SentAt     *time.Time `query:"sent_at"`
	FailedAt   *time.Time `query:"failed_at"`
	LastError  *string    `query:"last_error"`
}

type NewSmsJob struct {
	ID        string `query:"id"`
	Recipient string `query:"recipient"`
	Message   string `query:"message"`
}
