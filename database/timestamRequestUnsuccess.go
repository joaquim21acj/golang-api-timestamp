package database

import (
	"database/sql"
)

type TimestampRequestUnsuccess struct {
	Id                int
	Timestamp_request int64
	UserId            int
	Error_reported    error
}

func AddTimestampUnsuccess(db *sql.DB, newRequest TimestampRequestUnsuccess) error {
	stmt, err := db.Prepare("INSERT INTO user_request_unsuccess (timestamp_request, user_id, error_reported) VALUES (?, ?, ?)")
	stmt.Exec(newRequest.Timestamp_request, newRequest.UserId, newRequest.Error_reported.Error())
	defer stmt.Close()
	return err
}
