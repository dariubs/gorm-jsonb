package gormjsonb

import (
	"database/sql/driver"
	"encoding/json"
)

// JSONB type
type JSONB map[string]interface{}

// Value for save jsonb in postgres
func (v JSONB) Value() (driver.Value, error) {
	mjson, err := json.Marshal(v)

	return string(mjson), err
}

// Scan unmarshal data in JSONB map
func (m *JSONB) Scan(src interface{}) error {
	err := json.Unmarshal([]byte(value), &v)

	return err


	var source []byte
	_m := make(map[string]interface{})


	switch src.(type) {
	case []uint8:
		source = []byte(src.([]uint8))
	case nil:
		return nil
	default:
		return errors.New("incompatible type for StringInterfaceMap")
	}
	err := json.Unmarshal(source, &_m)
	if err != nil {
		return err
	}
	*m = StringInterfaceMap(_m)
	return nil
}
