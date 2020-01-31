package models

import (
	"encoding/json"
	"io"
)

type Student struct {
	Id      string `json:"Id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

func (s *Student) ToJson() ([]byte, error) {
	data, err := json.Marshal(s)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func FromJson(data io.Reader) *Student {
	var student *Student
	json.NewDecoder(data).Decode(&student)
	return student
}

func FromBytes(data []byte) *Student {
	student := &Student{}
	json.Unmarshal(data, student)
	return student
}
