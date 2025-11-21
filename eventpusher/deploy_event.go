package eventpusher

import (
	"math/rand"
)

type Event struct {
	Status    string `json:"status"`
	ErrorType string `json:"errorType"`
}

func randomEvent() Event {
	if rand.Float64() < 0.7 {
		return Event{Status: "success"}
	}
	errorTypes := []string{"timeout", "network", "db", "unknown"}
	return Event{
		Status:    "failure",
		ErrorType: errorTypes[rand.Intn(len(errorTypes))],
	}
}
