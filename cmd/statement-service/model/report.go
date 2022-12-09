package model

type Report struct {
	TotalBalance         float32
	AverageDebitAmount   float32
	AverageCreditAmount  float32
	TransactionsPerMonth map[int]int
}
