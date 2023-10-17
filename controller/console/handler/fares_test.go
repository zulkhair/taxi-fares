package handler

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/zulkhair/taxi-fares/domain/taxidata"
	"os"
	"testing"
	"time"

	"github.com/efficientgo/core/testutil"
	"github.com/rs/zerolog/log"
	mock_fares "github.com/zulkhair/taxi-fares/usecase/fares/mock"
)

func TestHandler_CalculateFares(t *testing.T) {
	// Mock
	type fields struct {
		fares *mock_fares.MockUsecase
	}
	// Input parameters
	type args struct {
		content []byte
	}

	// Test Case
	tests := []struct {
		prepare func(f *fields)
		name    string
		args    args
		want    error
	}{
		{
			prepare: func(f *fields) {
				// mock usecase
				f.fares.EXPECT().CalculateFares(gomock.Any()).Return(&taxidata.Fares{
					Fare: 100,
					TaxiData: []taxidata.TaxiData{
						{
							Time:              time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
							Distance:          100,
							MileageDifference: 100,
						},
					},
				})
			},
			name: "Success",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:01:00.123 480.9\n00:02:00.125 1141.2\n00:03:00.100 1800.8"),
			},
			want: nil,
		},
		{
			prepare: func(f *fields) {},
			name:    "Error less than two lines",
			args: args{
				content: []byte("\n"),
			},
			want: fmt.Errorf("less than two lines of data, input ''"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error less than two lines 2",
			args: args{
				content: []byte("00:00:00.000 0.0"),
			},
			want: fmt.Errorf("less than two lines of data, input '00:00:00.000 0.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error blank line",
			args: args{
				content: []byte("00:00:00.000 0.0\n "),
			},
			want: fmt.Errorf("blank line at row 1, input '00:00:00.000 0.0\n '"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error improper format",
			args: args{
				content: []byte("00:00:00.000 0.0\nimproper-format"),
			},
			want: fmt.Errorf("improper format at row 1, input '00:00:00.000 0.0\nimproper-format'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error invalid time format",
			args: args{
				content: []byte("00:00:00.000 0.0\n1:12:12.12 0.0"),
			},
			want: fmt.Errorf("invalid time format at row 1, input '00:00:00.000 0.0\n1:12:12.12 0.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error past time has been sent",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:01:00.000 100.0\n00:00:02.000 100.0"),
			},
			want: fmt.Errorf("past time has been sent at row 2, input '00:00:00.000 0.0\n00:01:00.000 100.0\n00:00:02.000 100.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error interval between records is more than 5 minutes",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:07:00.000 100.0\n00:00:02.000 100.0"),
			},
			want: fmt.Errorf("interval between records is more than 5 minutes at row 1, input '00:00:00.000 0.0\n00:07:00.000 100.0\n00:00:02.000 100.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error invalid distance format",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:03:00.000 asd\n00:00:02.000 100.0"),
			},
			want: fmt.Errorf("invalid distance format at row 1 with values 'asd', input '00:00:00.000 0.0\n00:03:00.000 asd\n00:00:02.000 100.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error invalid distance format 18a0",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:03:00.000 18a0\n00:05:02.000 100.0"),
			},
			want: fmt.Errorf("invalid distance format at row 1 with values '18a0', input '00:00:00.000 0.0\n00:03:00.000 18a0\n00:05:02.000 100.0'"),
		},
		{
			prepare: func(f *fields) {},
			name:    "Error total mileage is 0.0m",
			args: args{
				content: []byte("00:00:00.000 0.0\n00:03:00.000 0.0\n00:08:00.000 0.0"),
			},
			want: fmt.Errorf("total mileage is 0.0m, input '00:00:00.000 0.0\n00:03:00.000 0.0\n00:08:00.000 0.0'"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// prepare mock
			ctrl := gomock.NewController(t)
			f := &fields{
				fares: mock_fares.NewMockUsecase(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(f)
			}

			// prepare content
			tempFile, err := os.CreateTemp("", "taxi_data")
			if err != nil {
				log.Err(err)
			}
			if _, err := tempFile.Write(tt.args.content); err != nil {
				log.Err(err)
			}
			if _, err := tempFile.Seek(0, 0); err != nil {
				log.Err(err)
			}

			// run test
			handler := New(f.fares)
			got := handler.CalculateFares(tempFile)
			testutil.Equals(t, tt.want, got)

			// cleanup temp file
			err = tempFile.Close()
			if err != nil {
				log.Err(err)
			}
		})
	}
}
