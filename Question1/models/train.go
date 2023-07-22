package models

import "time"

type Coach struct {
	Type             string  `json:"type"`
	SeatAvailability int     `json:"seatAvailability"`
	Price            float64 `json:"price"`
}

type Train struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Departure time.Time `json:"departure"`
	Coaches   []Coach   `json:"coaches"`
}

var MockTrains = []Train{
	{
		ID:        1,
		Name:      "Express 1",
		Departure: time.Now().Add(time.Minute * 15),
		Coaches: []Coach{
			{
				Type:             "sleeper",
				SeatAvailability: 100,
				Price:            500.0,
			},
			{
				Type:             "ac",
				SeatAvailability: 50,
				Price:            1000.0,
			},
		},
	},
	{
		ID:        2,
		Name:      "Super Express 2",
		Departure: time.Now().Add(time.Minute * 40),
		Coaches: []Coach{
			{
				Type:             "sleeper",
				SeatAvailability: 100,
				Price:            1000.0,
			},
			{
				Type:             "ac",
				SeatAvailability: 30,
				Price:            1500.0,
			},
		},
	},
	{
		ID:        3,
		Name:      "Express 3",
		Departure: time.Now().Add(time.Hour * 10),
		Coaches: []Coach{
			{
				Type:             "sleeper",
				SeatAvailability: 60,
				Price:            500.0,
			},
			{
				Type:             "ac",
				SeatAvailability: 20,
				Price:            1000.0,
			},
		},
	},
}
