package model

import "time"

type TransactionModel struct {
	Id          uint8
	Date        time.Time
	Transaction float32
}
