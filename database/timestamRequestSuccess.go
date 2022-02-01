package database

import (
	"database/sql"
	"log"
	"time"
)

func AddTimestampSuccess(db *sql.DB, newRequest TimestampRequestSuccess) error {
	stmt, err := db.Prepare("INSERT INTO user_request_success (timestamp_request, user_id) VALUES (?, ?)")

	if err == nil {
		_, err := stmt.Exec(newRequest.Timestamp_request, newRequest.UserId)
		defer stmt.Close()
		return err
	}
	return err
}

func GetTimestampSuccessRequests(db *sql.DB) []TimestampSuccessRequisition {
	rows, err := db.Query("SELECT urs.id, urs.timestamp_request, urs.user_id, u.name FROM user_request_success urs INNER JOIN users u on u.id = urs.user_id")

	defer rows.Close()

	successRequests := make([]TimestampSuccessRequisition, 0)

	err = rows.Err()
	if err != nil {
		log.Panic(err)
		return successRequests
	}

	for rows.Next() {
		requestSuccess := TimestampSuccessRequisition{}
		err = rows.Scan(&requestSuccess.Id, &requestSuccess.Timestamp_request, &requestSuccess.Id, &requestSuccess.UserName)
		if err != nil {
			log.Panic(err)
		}

		successRequests = append(successRequests, requestSuccess)
	}
	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}

	return successRequests

}

type TimestampRequestSuccess struct {
	Id                int
	Timestamp_request int64
	UserId            int
}

type TimestampSuccessRequisition struct {
	Id                int
	Timestamp_request time.Time
	UserId            int
	UserName          string
}
