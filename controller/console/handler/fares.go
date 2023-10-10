package handler

import (
	"bufio"
	"fmt"
	"os"
)

func (h *Handler) CalculateFares() (err error) {
	scn := bufio.NewScanner(os.Stdin)

	fmt.Println("(To submit, use '!') Enter Lines :")
	var lines []string

	// Input
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			if line[0] == '!' {
				break
			}
		}
		lines = append(lines, line)
	}

	// Usecase Logic
	taxiData, err := h.fares.CalculateFares(lines)
	if err != nil {
		return
	}

	// Output
	fmt.Printf("\n%d\n", taxiData.Fare)
	for _, data := range taxiData.TaxiData {
		fmt.Printf("%s %.1f %.1f\n", data.Time.Format("15:04:05.000"), data.Distance, data.MileageDifference)
	}

	return
}
