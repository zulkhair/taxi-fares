package console

import (
	"github.com/zulkhair/taxi-fares/controller/console/handler"
	"github.com/zulkhair/taxi-fares/usecase/fares"
)

// Console is the struct for console. It contains the console instance.
type Console struct {
	handler *handler.Handler
}

// New creates a new console object.
func New() (*Console, error) {

	// Instantiate usecase
	rpsUsecase := fares.New()

	// Instantiate handler
	h := handler.New(rpsUsecase)

	// Create console object
	c := &Console{
		handler: h,
	}

	return c, nil
}

func (c *Console) StartCalculateTaxiFares() error {
	return c.handler.CalculateFares()
}
