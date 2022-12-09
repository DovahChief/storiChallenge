package main

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/DovahChief/storiChallenge/cmd/statement-service/database"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/handler"
	"github.com/DovahChief/storiChallenge/cmd/statement-service/logger"
)

// server

func main() {
	logger.Info(context.Background(), "--Init Application--")

	db := database.New()
	if db == nil {
		logger.Errorf(context.Background(), "Could not connect to DB")
	}

	h := handler.New(db)

	router := mux.NewRouter()
	router.HandleFunc("/statement", h.GenerateStatement).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Errorf(context.Background(), "An error has occurred [%v]", err)
	}

}
