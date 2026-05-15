package persistence

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func InsertSmsJob(ctx context.Context, db *sqlx.DB, job NewSmsJob) (SmsJob, error) {
	var insertedJob SmsJob

	query, args, err := squirrel.Insert("sms_jobs").
		Columns("id", "recipient", "message").
		Values(job.ID, job.Recipient, job.Message).
		ToSql()
	if err != nil {
		return insertedJob, err
	}

	if _, err := db.ExecContext(ctx, query, args...); err != nil {
		return insertedJob, err
	}

	return FindSmsJobById(ctx, db, job.ID)
}

func FindSmsJobById(ctx context.Context, db *sqlx.DB, jobID string) (SmsJob, error) {
	var job SmsJob

	query, args, err := squirrel.Select("*").
		From("sms_jobs").
		Where(squirrel.Eq{"id": jobID}).
		Limit(1).
		ToSql()
	if err != nil {
		return job, err
	}

	if err := db.GetContext(ctx, &job, query, args...); err != nil {
		return job, err
	}

	return job, nil
}
