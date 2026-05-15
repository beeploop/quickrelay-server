package persistence

import "github.com/jmoiron/sqlx"

func InitializeSchema(db *sqlx.DB) error {
	stmt := `
	CREATE TABLE IF NOT EXISTS sms_jobs (
		id TEXT PRIMARY KEY,
		recipient TEXT NOT NULL,
		message TEXT NOT NULL,

		status TEXT CHECK(status IN('pending', 'assigned', 'sent', 'failed')) DEFAULT 'pending',
		retry_count INTEGER NOT NULL DEFAULT 0,

		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

		sent_at DATETIME,
		failed_at DATETIME,
		last_error TEXT
	);

	CREATE TRIGGER IF NOT EXISTS update_sms_jobs_updated_at
	AFTER UPDATE ON sms_jobs
	FOR EACH ROW
	BEGIN
		UPDATE sms_jobs
		SET updated_at = CURRENT_TIMESTAMP
		WHERE id = NEW.id;
	END;
	`

	_, err := db.Exec(stmt)
	return err
}
