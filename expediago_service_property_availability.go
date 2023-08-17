package expediago

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
)

const expediaEndpointPropertiesAvailability = "/v3/properties/availability"

func GetPropertyAvailability(params ExpediaGoModelRequestAvailability) ([]*ModelResponseAvailability, []byte, error) {
	urlParams, err := query.Values(params)
	if err != nil {
		return nil, nil, ExpediaGoErrorToString(&ExpediaGoModelError{
			Message: err.Error(),
		})
	}
	result, statusCode := ExpediaGoApiRequest(expediaEndpointPropertiesAvailability+"?"+urlParams.Encode(), "GET", nil)

	if statusCode == 200 {
		var parsed []*ModelResponseAvailability
		err := json.Unmarshal(result, &parsed)
		if err != nil {
			return nil, nil, err
		}
		return parsed, result, nil
	} else {
		var data *ExpediaGoModelError
		err = json.Unmarshal(result, &data)
		if err != nil {
			return nil, nil, ExpediaGoErrorToString(&ExpediaGoModelError{
				Message: err.Error(),
			})
		}
		return nil, nil, ExpediaGoErrorToString(data)
	}

}
