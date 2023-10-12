package handler

import (
	"bufio"
	"fmt"
	taxidatadomain "github.com/zulkhair/taxi-fares/domain/taxidata"
	"os"
	"strings"
	"time"
)

// CalculateFares is a function to start input, calculate fares, and print output
func (h *Handler) CalculateFares(in *os.File) (err error) {
	/*** User Input ***/
	var lines []string
	var errorMessage string

	scn := bufio.NewScanner(in)
	first := true
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 0 {
			break
		}
		// append taxi data
		lines = append(lines, line)

		// construct error message
		if first {
			first = false
		} else {
			errorMessage += "\n"
		}
		errorMessage += line
	}

	/*** Validating Input ***/
	var (
		taxiData []taxidatadomain.TaxiData
		prevTime time.Time
	)

	// Check for the number of lines
	if len(lines) < 2 {
		return fmt.Errorf(taxidatadomain.LessThanTwoLinesData, errorMessage)
	}

	// Iterate lines of raw taxi data list
	for rowCount, line := range lines {
		// Check for blank line
		if len(strings.TrimSpace(line)) == 0 {
			return fmt.Errorf(taxidatadomain.BlankLineError, rowCount, errorMessage)
		}

		// Parse the input line
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return fmt.Errorf(taxidatadomain.ImproperFormat, rowCount, errorMessage)
		}

		// Parse elapsed time
		elapsedTime, err := time.Parse("15:04:05.000", parts[0])
		if err != nil {
			return fmt.Errorf(taxidatadomain.InvalidTimeFormat, rowCount, errorMessage)
		}

		// Check for past time
		if prevTime != (time.Time{}) && !elapsedTime.After(prevTime) {
			return fmt.Errorf(taxidatadomain.PastTimeHasSent, rowCount, errorMessage)
		}

		// Check the interval between records
		if prevTime == (time.Time{}) {
			// for checking first line
			prevTime, _ = time.Parse("15:04:05.000", "00:00:00.000")
		}
		if elapsedTime.Sub(prevTime) > 5*time.Minute {
			return fmt.Errorf(taxidatadomain.IntervalBetweenRecordsMoreThan5Minutes, rowCount, errorMessage)
		}

		// Parse distance
		distance, err := parseDistance(parts[1])
		if err != nil {
			return fmt.Errorf(taxidatadomain.InvalidDistanceFormat, rowCount, parts[1], errorMessage)
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
		return fmt.Errorf(taxidatadomain.TotalMileageZero, errorMessage)
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
