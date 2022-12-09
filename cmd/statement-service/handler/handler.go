package handler

import (
	"encoding/json"
	"net/http"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/database"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/emailSender"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/fileProcessor"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/model"
)

type Handler struct {
	db *database.Database
}

func New(db *database.Database) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GenerateStatement(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	logger.Info(ctx, "--Init Request for statement Generation--")

	var req model.Request
	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		logger.Error(ctx, "Error Parsing request Body")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Email == "" {
		logger.Error(ctx, "missing email in request body")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	fileProcessor := fileProcessor.New("Test.csv")
	report, err := fileProcessor.ProcessFile(ctx)
	if err != nil {
		logger.Error(ctx, "Error processing file")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info(ctx, "-- File Processed --")

	reportId, err := h.db.SaveReport(ctx, report, req.Email)
	if err != nil {
		logger.Error(ctx, "Error saving report to dB")
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	logger.Info(ctx, "-- Report Saved to DB --", reportId)

	for _, tr := range fileProcessor.GetTransactions() {
		go h.db.SaveTransaction(ctx, reportId, tr)
	}

	emailSender.SendEmail(ctx, req.Email, report)

	logger.Info(ctx, "-- Email Sent --")

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

}
