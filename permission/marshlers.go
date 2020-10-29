package permission

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

var ErrInvalidPermission = errors.New("invalid permission")

func stringToType(str string) (Type, error) {
	for _, perm := range All {
		if perm.Code == str {
			return perm, nil
		}
	}

	return Type{}, ErrInvalidPermission
}

// Methods for marshaling and unmarshaling from JSON.
func (perm Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(perm.Code)
}

func (perm *Type) UnmarshalJSON(data []byte) error {
	var code string
	err := json.Unmarshal(data, &code)
	if err != nil {
		return err
	}

	*perm, err = stringToType(code)
	return err
}

// Methods for converting to sql.
func (perm Type) Value() (driver.Value, error) {
	return perm.Code, nil
}

func (perm *Type) Scan(value interface{}) error {
	code, ok := value.(string)
	if !ok {
		return ErrInvalidPermission
	}

	var err error
	*perm, err = stringToType(code)
	return err
}
