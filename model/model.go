package model

import (
	"database/sql"
	"log"
)

var (
	DB     *sql.DB
	Logger *log.Logger
)
