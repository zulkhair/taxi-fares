package handler

import (
	"bufio"
	"fmt"
	"github.com/rs/zerolog/log"
	taxidatadomain "github.com/zulkhair/taxi-fares/domain/taxidata"
	"os"
	"strings"
	"time"
)

func (h *Handler) CalculateFares(in *os.File) (err error) {
	/*** User Input ***/
	fmt.Println("(To submit, use '!') Enter Lines :")
	var lines []string

	scn := bufio.NewScanner(in)
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			if line[0] == '!' {
				break
			}
		}
		lines = append(lines, line)
	}

	/*** Validating Input ***/
	var (
		taxiData []taxidatadomain.TaxiData
		prevTime time.Time
	)

	// Check for the number of lines
	if len(lines) < 2 {
		log.Error().Msgf(taxidatadomain.LessThanTwoLinesData, lines)
		return fmt.Errorf(taxidatadomain.LessThanTwoLinesData, lines)
	}

	// Iterate lines of raw taxi data list
	for rowCount, line := range lines {
		// Check for blank line
		if len(strings.TrimSpace(line)) == 0 {
			log.Error().Msgf(taxidatadomain.BlankLineError, rowCount, lines)
			return fmt.Errorf(taxidatadomain.BlankLineError, rowCount, lines)
		}

		// Parse the input line
		parts := strings.Fields(line)
		if len(parts) != 2 {
			log.Error().Msgf(taxidatadomain.ImproperFormat, rowCount, lines)
			return fmt.Errorf(taxidatadomain.ImproperFormat, rowCount, lines)
		}

		// Parse elapsed time
		elapsedTime, err := time.Parse("15:04:05.000", parts[0])
		if err != nil {
			log.Error().Msgf(taxidatadomain.InvalidTimeFormat, rowCount, lines)
			return fmt.Errorf(taxidatadomain.InvalidTimeFormat, rowCount, lines)
		}

		// Check for past time
		if prevTime != (time.Time{}) && !elapsedTime.After(prevTime) {
			log.Error().Msgf(taxidatadomain.PastTimeHasSent, rowCount, lines)
			return fmt.Errorf(taxidatadomain.PastTimeHasSent, rowCount, lines)
		}

		// Check the interval between records
		if prevTime == (time.Time{}) {
			// for checking first line
			prevTime, _ = time.Parse("15:04:05.000", "00:00:00.000")
		}
		if elapsedTime.Sub(prevTime) > 5*time.Minute {
			log.Error().Msgf(taxidatadomain.IntervalBetweenRecordsMoreThan5Minutes, rowCount, lines)
			return fmt.Errorf(taxidatadomain.IntervalBetweenRecordsMoreThan5Minutes, rowCount, lines)
		}

		// Parse distance
		distance, err := parseDistance(parts[1])
		if err != nil {
			log.Error().Msgf(taxidatadomain.InvalidDistanceFormat, rowCount, parts[1], lines)
			return fmt.Errorf(taxidatadomain.InvalidDistanceFormat, rowCount, parts[1], lines)
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
		log.Error().Msgf(taxidatadomain.TotalMileageZero, lines)
		return fmt.Errorf(taxidatadomain.TotalMileageZero, lines)
	}

	/*** Usecase Logic ***/
	fares := h.fares.CalculateFares(taxiData)

	/*** Print Output ***/
	fmt.Printf("\n%d\n", fares.Fare)
	for _, data := range fares.TaxiData {
		fmt.Printf("%s %.1f %.1f\n", data.Time.Format("15:04:05.000"), data.Distance, data.MileageDifference)
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
