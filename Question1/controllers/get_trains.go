package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/swapnika/train_data/models"
)

func GetTrainsHanlder(w http.ResponseWriter, r *http.Request) {

	access_token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTAwMDcwODAsImNvbXBhbnlOYW1lIjoiRXhwcmVzcyBUcmFpbnMiLCJjbGllbnRJRCI6IjEzZjU1N2M1LWJmMmMtNGY2NS1iNjdhLWY3YzhjZTUwNmFkZCIsIm93bmVyTmFtZSI6IiIsIm93bmVyRW1haWwiOiIiLCJyb2xsTm8iOiIyMDAwMzIwMTAwMTczIn0.hRgx-u9byxzmVcv_PMhDypmi7xOiTN3SrwC4KXLs4tE"

	apiURL := "http://20.244.56.144/train/trains"
	req, err := http.NewRequest("GET", apiURL, nil)

	if err != nil {
		http.Error(w, "Failed to create API request", http.StatusInternalServerError)
		return
	}
	req.Header.Add("access_token", "Bearer "+access_token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch data from API", http.StatusInternalServerError)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read data from API", http.StatusInternalServerError)
		return
	}

	print(string(body))

	var trains []models.Train
	err = json.Unmarshal(body, &trains)
	if err != nil {
		http.Error(w, "Failed to parse data from API", http.StatusInternalServerError)
		return
	}

	now := time.Now()
	next30Mins := now.Add(time.Minute * 30)
	next12Hours := now.Add(time.Hour * 12)
	var filteredTrains []models.Train
	for _, train := range trains {
		departureTime := time.Date(now.Year(), now.Month(), now.Day(), train.DepartureTime.Hours, train.DepartureTime.Minutes, train.DepartureTime.Seconds, 0, now.Location())
		if departureTime.After(next30Mins) && departureTime.Before(next12Hours) {
			filteredTrains = append(filteredTrains, train)
		}
	}

	sort.SliceStable(filteredTrains, func(i, j int) bool {
		if filteredTrains[i].Price["AC"] != filteredTrains[j].Price["AC"] {
			return filteredTrains[i].Price["AC"] < filteredTrains[j].Price["AC"]
		}
		if filteredTrains[i].SeatsAvailable["sleeper"] != filteredTrains[j].SeatsAvailable["sleeper"] {
			return filteredTrains[i].SeatsAvailable["sleeper"] > filteredTrains[j].SeatsAvailable["sleeper"]
		}
		return compareTimePoints(filteredTrains[i].DepartureTime, filteredTrains[j].DepartureTime)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filteredTrains)
}

func compareTimePoints(tp1, tp2 models.TimePoint) bool {
	if tp1.Hours != tp2.Hours {
		return tp1.Hours < tp2.Hours
	}
	if tp1.Minutes != tp2.Minutes {
		return tp1.Minutes < tp2.Minutes
	}
	return tp1.Seconds < tp2.Seconds
}
