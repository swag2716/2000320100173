package controllers

import (
	"encoding/json"
	"net/http"
	"sort"
	"time"

	"github.com/swapnika/train_data/models"
)

func GetTrainsHanlder(w http.ResponseWriter, r *http.Request) {
	var filteredTrains []models.Train

	now := time.Now()
	next30Mins := now.Add(time.Minute * 30)
	next12Hours := now.Add(time.Hour * 12)
	for _, train := range models.MockTrains {
		if train.Departure.After(next30Mins) && train.Departure.Before(next12Hours) {
			filteredTrains = append(filteredTrains, train)
		}
	}

	sort.SliceStable(filteredTrains, func(i, j int) bool {
		if filteredTrains[i].Coaches[1].Price != filteredTrains[j].Coaches[1].Price {
			return filteredTrains[i].Coaches[1].Price < filteredTrains[j].Coaches[1].Price
		}
		if filteredTrains[i].Coaches[0].SeatAvailability != filteredTrains[j].Coaches[0].SeatAvailability {
			return filteredTrains[i].Coaches[0].SeatAvailability > filteredTrains[j].Coaches[0].SeatAvailability
		}
		return filteredTrains[i].Departure.After(filteredTrains[j].Departure)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredTrains)
}
