package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Repos struct {
	Properties PropertyRepo
}

func NewRepos(db *sql.DB) Repos {
	return Repos{
		Properties: PropertyRepo{DB: db},
	}
}
