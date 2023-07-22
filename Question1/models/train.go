package models

type TimePoint struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

type Train struct {
	TrainNumber    int            `json:"trainNumber"`
	TrainName      string         `json:"trainName"`
	DepartureTime  TimePoint      `json:"departureTime"`
	SeatsAvailable map[string]int `json:"seatsAvailable"` // Key: coach type (e.g., "sleeper" or "ac"), Value: number of available seats
	Price          map[string]int `json:"price"`
}
