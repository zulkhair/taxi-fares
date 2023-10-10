package fares

import (
	"testing"
	"time"

	"github.com/efficientgo/core/testutil"
	taxidatadomain "github.com/zulkhair/taxi-fares/domain/taxidata"
)

func createTaxiData(hour, minute, second, nanosecond int, distance, mileageDifference float64) taxidatadomain.TaxiData {
	t := time.Time{}
	return taxidatadomain.TaxiData{
		Time:              time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, nanosecond, t.Location()),
		Distance:          distance,
		MileageDifference: mileageDifference,
	}
}
func TestFares_CalculateFares(t *testing.T) {
	type args struct {
		taxiData []taxidatadomain.TaxiData
	}

	tests := []struct {
		want    *taxidatadomain.Fares
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success above 10km",
			args: args{
				taxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 0, 0, 0, 0, 0),
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 2, 0, 125, 1141.2, 660.3),
					createTaxiData(0, 3, 0, 100, 1800.8, 659.6),
					createTaxiData(0, 4, 0, 100, 2800.8, 1000),
					createTaxiData(0, 5, 0, 100, 3800.8, 1000),
					createTaxiData(0, 6, 0, 100, 7800.0, 4000),
					createTaxiData(0, 7, 0, 100, 10800, 2999.2),
				},
			},
			want: &taxidatadomain.Fares{
				Fare: 1360,
				TaxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 6, 0, 100, 7800.0, 4000),
					createTaxiData(0, 7, 0, 100, 10800, 2999.2),
					createTaxiData(0, 4, 0, 100, 2800.8, 1000),
					createTaxiData(0, 5, 0, 100, 3800.8, 1000),
					createTaxiData(0, 2, 0, 125, 1141.2, 660.3),
					createTaxiData(0, 3, 0, 100, 1800.8, 659.6),
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 0, 0, 0, 0, 0),
				},
			},
			wantErr: false,
		},
		{
			name: "Success above 1km below 10km",
			args: args{
				taxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 0, 0, 0, 0, 0),
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 2, 0, 125, 1141.2, 660.3),
					createTaxiData(0, 3, 0, 100, 1800.8, 659.6),
					createTaxiData(0, 4, 0, 100, 2800.8, 1000),
					createTaxiData(0, 5, 0, 100, 3800.8, 1000),
					createTaxiData(0, 6, 0, 100, 7800.0, 4000),
				},
			},
			want: &taxidatadomain.Fares{
				Fare: 1080,
				TaxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 6, 0, 100, 7800.0, 4000),
					createTaxiData(0, 4, 0, 100, 2800.8, 1000),
					createTaxiData(0, 5, 0, 100, 3800.8, 1000),
					createTaxiData(0, 2, 0, 125, 1141.2, 660.3),
					createTaxiData(0, 3, 0, 100, 1800.8, 659.6),
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 0, 0, 0, 0, 0),
				},
			},
			wantErr: false,
		},
		{
			name: "Success below 1km",
			args: args{
				taxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 0, 0, 0, 0, 0),
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 2, 0, 125, 900.2, 419.3),
				},
			},
			want: &taxidatadomain.Fares{
				Fare: 400,
				TaxiData: []taxidatadomain.TaxiData{
					createTaxiData(0, 1, 0, 123, 480.9, 480.9),
					createTaxiData(0, 2, 0, 125, 900.2, 419.3),
					createTaxiData(0, 0, 0, 0, 0, 0),
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			fares := New()

			got, err := fares.CalculateFares(tt.args.taxiData)
			switch {
			case tt.wantErr:
				testutil.NotOk(t, err)
			default:
				testutil.Ok(t, err)
			}

			testutil.Equals(t, tt.want, got)
		})
	}
}
