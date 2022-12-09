package csvreader

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/model"
)

type FileParser struct {
	csvData [][]string
}

func NewFileParser(csvData [][]string) *FileParser {
	return &FileParser{
		csvData: csvData,
	}
}

func (fp *FileParser) ParseCsvFile(ctx context.Context) ([]model.TransactionModel, error) {

	var response []model.TransactionModel

	for _, line := range fp.csvData[1:] {

		id, err := parseIdField(ctx, line[0])
		if err != nil {
			return nil, err
		}

		date, err := parseDateField(ctx, line[1])
		if err != nil {
			return nil, err
		}

		amount, err := parseAmountField(ctx, line[2])
		if err != nil {
			return nil, err
		}

		if amount != 0 {
			response = append(response, model.TransactionModel{
				Id:          id,
				Date:        date,
				Transaction: amount,
			})
		}
	}

	return response, nil
}

func parseIdField(ctx context.Context, strId string) (uint8, error) {
	id, err := strconv.Atoi(strId)
	if err != nil {
		logger.Error(ctx, "--Error parsing id from csvData--")
		return 0, err
	}

	if id < 0 {
		logger.Error(ctx, "--Error, invalid ID number --")
		return 0, errors.New("invalid ID number")
	}

	return uint8(id), nil
}

func parseAmountField(ctx context.Context, strAm string) (float32, error) {
	amount, err := strconv.ParseFloat(strAm, 32)
	if err != nil {
		logger.Error(ctx, "--Error parsing amount from csvData--")
		return 0, err
	}

	return float32(amount), err
}

func parseDateField(ctx context.Context, strDate string) (time.Time, error) {
	year := strconv.Itoa(time.Now().Year())

	dateElems := strings.Split(strDate, "/")
	month, err := strconv.Atoi(dateElems[0])
	if err != nil {
		logger.Error(ctx, "--Error parsing date month from csvData--")
		return time.Now(), err
	}

	day, err := strconv.Atoi(dateElems[1])
	if err != nil {
		logger.Error(ctx, "--Error parsing date day from csvData--")
		return time.Now(), err
	}

	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)

	var dateString = year + "/" + monthStr + "/" + dayStr
	date, err := time.Parse("2006/01/02", dateString)
	if err != nil {
		logger.Error(ctx, "--Error parsing date from csvData--")
		return time.Now(), err
	}

	return date, err
}
