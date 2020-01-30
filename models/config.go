package models

type Config struct {
	DB   string `json: "db"`
	Port string `json: "port"`
	Host string `json: "host"`
}
