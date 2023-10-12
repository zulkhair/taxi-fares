package taxidata

import "time"

const (
	BlankLineError                         = "blank line at row %d, input '%s'"
	ImproperFormat                         = "improper format at row %d, input '%s'"
	InvalidTimeFormat                      = "invalid time format at row %d, input '%s'"
	IntervalBetweenRecordsMoreThan5Minutes = "interval between records is more than 5 minutes at row %d, input '%s'"
	PastTimeHasSent                        = "past time has been sent at row %d, input '%s'"
	InvalidDistanceFormat                  = "invalid distance format at row %d with values '%s', input '%s'"
	LessThanTwoLinesData                   = "less than two lines of data, input '%s'"
	TotalMileageZero                       = "total mileage is 0.0m, input '%s'"
)

type TaxiData struct {
	Time              time.Time
	Distance          float64
	MileageDifference float64
}

type Fares struct {
	Fare     int
	TaxiData []TaxiData
}
