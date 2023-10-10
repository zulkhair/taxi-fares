package fares

import (
	taxidatadomain "github.com/zulkhair/taxi-fares/domain/taxidata"
	"sort"
)

// Usecase is an interface that handles business logic for fares domain.
type Usecase interface {
	// CalculateFares : is a function to calculate fares and return sorted taxis data
	CalculateFares(taxiData []taxidatadomain.TaxiData) *taxidatadomain.Fares
}

// Fares is an object that implements Usecase interface.
type Fares struct {
}

// New inst
func New() Usecase {
	return Fares{}
}

// CalculateFares is a function to calculate fares and return sorted taxis data
func (uc Fares) CalculateFares(taxiData []taxidatadomain.TaxiData) *taxidatadomain.Fares {
	// Create return value
	fares := &taxidatadomain.Fares{
		Fare:     calculateFare(taxiData[len(taxiData)-1].Distance),
		TaxiData: taxiData,
	}

	// Sort the taxiData by mileage difference in descending order
	sort.Slice(taxiData[:], func(i, j int) bool {
		return taxiData[i].MileageDifference > taxiData[j].MileageDifference
	})

	return fares
}

// calculateFare is function to calculate fare from total mileage
func calculateFare(distance float64) int {
	baseFare := 400
	if distance <= 1000 {
		return baseFare
	}

	remainingDistance := distance - 1000
	if remainingDistance <= 9000 {
		return baseFare + int(remainingDistance/400)*40
	}

	remainingDistance -= 9000
	return baseFare + 9000/400*40 + int(remainingDistance/350)*40
}
