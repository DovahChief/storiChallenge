package csvreader

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
)

type FileGetter struct {
	fileName string
}

func NewFileGetter(s string) *FileGetter {
	return &FileGetter{fileName: s}
}

func (fg *FileGetter) ReadFile(ctx context.Context) ([][]string, error) {

	content, err := os.Open(fg.fileName)
	if err != nil {
		logger.Error(ctx, "-- Error Opening File --")
		return [][]string{}, err
	}
	defer content.Close()

	lines, err := csv.NewReader(content).ReadAll()
	if err != nil {
		logger.Error(ctx, "-- Error Reading CSV File --")
		return [][]string{}, err
	}
	return lines, nil

}
