package taxidata

import "time"

const (
	BlankLineError                         = "blank line at row %d"
	ImproperFormat                         = "improper format at row %d"
	InvalidTimeFormat                      = "invalid time format at row %d"
	IntervalBetweenRecordsMoreThan5Minutes = "interval between records is more than 5 minutes at row %d"
	PastTimeHasSent                        = "past time has been sent at row %d"
	InvalidDistanceFormat                  = "invalid distance format at row %d with values '%s'"
	LessThanTwoLinesData                   = "less than two lines of data %v"
	TotalMileageZero                       = "total mileage is 0.0m"
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
