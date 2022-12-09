package fileProcessor

import (
	"context"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/fileProcessor/csvreader"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/model"
)

type FileProcessor struct {
	fileName     string
	transactions []model.TransactionModel
}

func New(fn string) *FileProcessor {
	return &FileProcessor{
		fileName: fn,
	}
}

func (fp *FileProcessor) ProcessFile(ctx context.Context) (model.Report, error) {
	fileGetter := csvreader.NewFileGetter(fp.fileName)

	fileRawData, err := fileGetter.ReadFile(ctx)
	if err != nil {
		logger.Error(ctx, "Error reading file")
		return model.Report{}, err
	}

	fileParser := csvreader.NewFileParser(fileRawData)

	fp.transactions, err = fileParser.ParseCsvFile(ctx)
	if err != nil {
		logger.Error(ctx, "Error processing file")
		return model.Report{}, err
	}

	report := fp.generateReport()

	return report, nil
}

func (fp *FileProcessor) GetTransactions() []model.TransactionModel {
	return fp.transactions
}

func (fp *FileProcessor) generateReport() model.Report {

	report := model.Report{
		TotalBalance:         0,
		AverageDebitAmount:   0,
		AverageCreditAmount:  0,
		TransactionsPerMonth: make(map[int]int, 12),
	}

	var totalCredit float32 = 0
	var totalDebit float32 = 0
	debitTransactionCount := 0
	creditTransactionCount := 0

	for _, transaction := range fp.transactions {
		report.TotalBalance += transaction.Transaction

		if transaction.Transaction > 0 {
			totalCredit += transaction.Transaction
			creditTransactionCount++
		} else {
			totalDebit += transaction.Transaction
			debitTransactionCount++
		}

		transactionMonth := int(transaction.Date.Month())
		report.TransactionsPerMonth[transactionMonth]++
	}

	report.AverageCreditAmount = totalCredit / float32(creditTransactionCount)
	report.AverageDebitAmount = totalDebit / float32(debitTransactionCount)

	return report

}
