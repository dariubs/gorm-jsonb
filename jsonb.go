package gormjsonb

import (
	"database/sql/driver"
	"encoding/json"
)

type JSONB map[string]interface{}

// Value for save jsonb in postgres
func (v JSONB) Value() (driver.Value, error) {
	mjson, err := json.Marshal(v)

	return string(mjson), err
}

// Scan unmarshal data in JSONB map
func (v *JSONB) Scan(value interface{}) error {
	err := json.Unmarshal(value.([]byte), &v)

	return err
}