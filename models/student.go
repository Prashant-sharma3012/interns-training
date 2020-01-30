package models

import (
	"encoding/json"
)

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func (s *Student) toJson() ([]byte, error) {
	data, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (s *Student) fromJson(data []byte) error {
	return json.Unmarshal(data, s)
}
