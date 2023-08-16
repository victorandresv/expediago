package expediago

import (
	"encoding/json"
	"errors"
)

func ExpediaGoErrorToString(data *ExpediaGoModelError) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return errors.New("")
	}
	return errors.New(string(marshal))
}
