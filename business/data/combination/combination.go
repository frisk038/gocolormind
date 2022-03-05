// Package combination contains combination related CRUD functionality.
package combination

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type DailyCombination struct {
	Combi     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DB struct {
	*sql.DB
}

func (db DB) Combination() (DailyCombination, error) {
	var dCombi DailyCombination

	row := db.QueryRow("SELECT * FROM daily_combination LIMIT 1")
	if err := row.Scan(&dCombi.Combi, &dCombi.CreatedAt, &dCombi.UpdatedAt); err != nil {
		return dCombi, errors.New("no combination found")
	}
	return dCombi, nil
}

func (db DB) AddCombination(cmb DailyCombination) error {
	_, err := db.Exec("INSERT INTO daily_combination (combination, date_created, date_updated) VALUES ($1, $2, $3)", cmb.Combi, cmb.CreatedAt, cmb.UpdatedAt)
	if err != nil {
		return fmt.Errorf("addCombination: %v", err)
	}

	return nil
}

func (db DB) UpdateCombination(cmb DailyCombination) error {
	_, err := db.Exec("INSERT INTO daily_combination (combination, date_updated) VALUES ($1, $2)", cmb.Combi, cmb.UpdatedAt)
	if err != nil {
		return fmt.Errorf("UpdateCombination: %v", err)
	}

	return nil
}
