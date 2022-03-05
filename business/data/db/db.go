// Package db contains combination related CRUD functionality.
package db

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

type Store struct {
	db *sql.DB
}

// NewStore constructs a data for api access.
func NewStore(db *sql.DB) Store {
	return Store{
		db: db,
	}
}

func (DB Store) Combination() (DailyCombination, error) {
	var dCombi DailyCombination

	row := DB.db.QueryRow("SELECT * FROM daily_combination LIMIT 1")
	if err := row.Scan(&dCombi.Combi, &dCombi.CreatedAt, &dCombi.UpdatedAt); err != nil {
		return dCombi, errors.New("no combination found")
	}
	return dCombi, nil
}

func (DB Store) AddCombination(cmb DailyCombination) (int64, error) {
	const q = `
	INSERT INTO daily_combination (
        combination,
        date_created,
        date_updated
    )
	VALUES ($1, $2, $3)`

	result, err := DB.db.Exec(q, cmb.Combi, cmb.CreatedAt, cmb.UpdatedAt)
	if err != nil {
		return 0, fmt.Errorf("addCombination: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addCombination: %v", err)
	}

	return id, nil
}

func (DB Store) UpdateCombination(cmb DailyCombination, cmbID int64) error {
	const q = `
	UPDATE daily_combination
	SET (combination, date_updated)
	VALUES ($1, $2)
	WHERE id = $3`

	_, err := DB.db.Exec(q, cmb.Combi, cmb.UpdatedAt, cmbID)
	if err != nil {
		return fmt.Errorf("UpdateCombination: %v", err)
	}

	return nil
}
