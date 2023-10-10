package fares

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	taxidatadomain "github.com/zulkhair/taxi-fares/domain/taxidata"
)

// Usecase is an interface that handles business logic for fares domain.
type Usecase interface {
	// CalculateFares : is a function to calculate fares and return sorted taxis data
	CalculateFares(lines []string) (fares *taxidatadomain.Fares, err error)
}

// Fares is an object that implements Usecase interface.
type Fares struct {
}

// New inst
func New() Usecase {
	return Fares{}
}

// CalculateFares is a function to calculate fares and return sorted taxis data
func (uc Fares) CalculateFares(lines []string) (fares *taxidatadomain.Fares, err error) {
	var (
		taxiData []taxidatadomain.TaxiData
		prevTime = time.Time{}
	)

	// Check for the number of lines
	if len(lines) < 2 {
		log.Error().Msgf(taxidatadomain.LessThanTwoLinesData, taxiData)
		return nil, fmt.Errorf(taxidatadomain.LessThanTwoLinesData, taxiData)
	}

	// Iterate lines of raw taxi data list
	for rowCount, line := range lines {
		// Check for blank line
		if len(strings.TrimSpace(line)) == 0 {
			log.Error().Msgf(taxidatadomain.BlankLineError, rowCount)
			return nil, fmt.Errorf(taxidatadomain.BlankLineError, rowCount)
		}

		// Parse the input line
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Error().Msgf(taxidatadomain.ImproperFormat, rowCount)
			return nil, fmt.Errorf(taxidatadomain.ImproperFormat, rowCount)
		}

		// Parse elapsed time
		elapsedTime, err := time.Parse("15:04:05.000", parts[0])
		if err != nil {
			log.Error().Msgf(taxidatadomain.InvalidTimeFormat, rowCount)
			return nil, fmt.Errorf(taxidatadomain.InvalidTimeFormat, rowCount)
		}

		// Check for past time
		if !elapsedTime.After(prevTime) {
			log.Error().Msgf(taxidatadomain.PastTimeHasSent, rowCount)
			return nil, fmt.Errorf(taxidatadomain.PastTimeHasSent, rowCount)
		}

		// Check the interval between records
		// TODO : Implement interval between records for 00:00:00
		if prevTime != (time.Time{}) && elapsedTime.Sub(prevTime) > 5*time.Minute {
			log.Error().Msgf(taxidatadomain.IntervalBetweenRecordsMoreThan5Minutes, rowCount)
			return nil, fmt.Errorf(taxidatadomain.IntervalBetweenRecordsMoreThan5Minutes, rowCount)
		}

		// Parse distance
		distance, err := parseDistance(parts[1])
		if err != nil {
			log.Error().Msgf(taxidatadomain.InvalidDistanceFormat, rowCount, parts[1])
			return nil, fmt.Errorf(taxidatadomain.InvalidDistanceFormat, rowCount, parts[1])
		}

		// calculate mileage difference
		mileageDifference := 0.0
		if rowCount > 0 {
			mileageDifference = distance - taxiData[rowCount-1].Distance
		}

		// Add the data to the slice
		t := taxidatadomain.TaxiData{Time: elapsedTime, Distance: distance, MileageDifference: mileageDifference}
		taxiData = append(taxiData, t)

		// Set prevTime
		prevTime = elapsedTime
	}

	// Check if the total mileage is 0.0m
	if taxiData[len(taxiData)-1].Distance == 0.0 {
		log.Error().Msgf(taxidatadomain.TotalMileageZero)
		return nil, fmt.Errorf(taxidatadomain.TotalMileageZero)
	}

	// Sort the taxiData by mileage difference in descending order
	sort.Slice(taxiData[:], func(i, j int) bool {
		return taxiData[i].MileageDifference > taxiData[j].MileageDifference
	})

	// Create return value
	fares = &taxidatadomain.Fares{
		Fare:     calculateFare(taxiData[len(taxiData)-1].Distance),
		TaxiData: taxiData,
	}

	return
}

// parseDistance is function to parse distance from string
func parseDistance(distanceStr string) (float64, error) {
	var distance float64
	_, err := fmt.Sscanf(distanceStr, "%f", &distance)
	if err != nil {
		return 0.0, fmt.Errorf("invalid distance format")
	}
	return distance, nil
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