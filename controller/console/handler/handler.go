package handler

import (
	"github.com/zulkhair/taxi-fares/usecase/fares"
)

// Handler is struct for console handler. It contains the necessary usecases for handling business logic.
type Handler struct {
	fares fares.Usecase
}

// New creates a new console job handler.
func New(fares fares.Usecase) *Handler {
	return &Handler{
		fares: fares,
	}
}
