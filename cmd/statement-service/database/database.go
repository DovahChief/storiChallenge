package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/model"
)

type Database struct {
	db *sql.DB
}

func New() *Database {

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
	db, err := sql.Open("postgres", postgresqlDbInfo)
	if err != nil {
		logger.Errorf(context.Background(), "Error on connection to DB [%v]", err)
		return nil
	}

	return &Database{db: db}
}

const insertReportQuery = `INSERT INTO report (email, balance, avg_debit, avg_credit) VALUES ($1, $2, $3, $4) RETURNING id`

func (db *Database) SaveReport(ctx context.Context, report model.Report, email string) (int, error) {
	reportId := 0
	err := db.db.QueryRowContext(ctx, insertReportQuery, email, report.TotalBalance, report.AverageDebitAmount, report.AverageDebitAmount).
		Scan(&reportId)
	if err != nil {
		logger.Errorf(ctx, "Error inserting report to DB [%v]", err)
		return 0, err
	}

	return reportId, nil
}

const insertTransactionQuery = `INSERT INTO transaction (transaction_id, transaction_date, transaction_amount, report_id) VALUES ($1, $2, $3, $4)`

func (db *Database) SaveTransaction(ctx context.Context, reportId int, tr model.TransactionModel) error {
	_, err := db.db.ExecContext(ctx, insertTransactionQuery, tr.Id, tr.Date, tr.Transaction, reportId)
	if err != nil {
		logger.Errorf(ctx, "Error inserting transaction to  DB [%v]", err)
		return err
	}
	return nil
}
