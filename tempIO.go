package main

import (
	"encoding/json"
	"time"
)

type tempIO struct {
	DateTime    time.Time
	Temperature string
	Humidity    string
}

func (tmpIO tempIO) json() (string, error) {
	b, err := json.Marshal(tmpIO)

	return string(b), err
}
