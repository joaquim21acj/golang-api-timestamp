package controllers

import (
	"database/sql"
	"main/database"
	"time"
)

func GetTimestamp(db *sql.DB) ReturnTimestamp {
	now := time.Now()
	timestamp := ReturnTimestamp{
		now.Unix(),
	}

	newRequest := database.TimestampRequestSuccess{
		Timestamp_request: timestamp.Timestamp,
		UserId:            1,
	}

	executeDBInsert(db, newRequest, 3)

	return timestamp
}

func GetAllRequests(db *sql.DB) []database.TimestampSuccessRequisition {
	requests := database.GetTimestampSuccessRequests(db)
	return requests
}

func executeDBInsert(db *sql.DB, newRequest database.TimestampRequestSuccess, retry int) {
	err := database.AddTimestampSuccess(db, newRequest)
	if err != nil {
		time.Sleep(2 * time.Second)
		executeDBInsert(db, newRequest, retry-1)
	}
	if retry <= 0 && err != nil {
		unsuccessRequest := database.TimestampRequestUnsuccess{
			Timestamp_request: newRequest.Timestamp_request,
			UserId:            newRequest.UserId,
			Error_reported:    err,
		}
		database.AddTimestampUnsuccess(db, unsuccessRequest)
	}
}

type ReturnTimestamp struct {
	Timestamp int64
}
