package config

import "todo/lib"

type Config struct {
	lib.ServerConfig
	ServerUrl string `json:"serverUrl"`
	Token     string `json:"token"`
}
