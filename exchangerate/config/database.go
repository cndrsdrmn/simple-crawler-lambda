package config

import (
	u "exchangerate/utils"
)

var DB *db

type dbDriver string

const MySQL dbDriver = "mysql"

type db struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Driver   dbDriver
}

func init() {
	DB = &db{
		Name:     u.Env("DB_DATABASE"),
		Host:     u.Env("DB_HOST", "localhost"),
		Port:     u.EnvInt("DB_PORT", 3306),
		User:     u.Env("DB_USERNAME"),
		Password: u.Env("DB_PASSWORD"),
		Driver:   dbDriver(u.Env("DB_CONNECTION", string(MySQL))),
	}
}
