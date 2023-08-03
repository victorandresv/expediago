package expediago

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/victorandresv/expediago/models"
)

const expediaEndpointPropertiesAvailability = "/v3/properties/availability"

func GetPropertyAvailability(params models.PropertyAvailabilityRequestModel) (interface{}, []byte, *models.ErrorModel) {
	urlParams, err := query.Values(params)
	if err != nil {
		return nil, nil, &models.ErrorModel{
			Message: err.Error(),
		}
	}
	result, statusCode := ApiRequest(expediaEndpointPropertiesAvailability+"?"+urlParams.Encode(), "GET", nil)

	if statusCode == 200 {
		return result, result, nil
	} else {
		var data *models.ErrorModel
		err = json.Unmarshal(result, &data)
		if err != nil {
			return nil, nil, &models.ErrorModel{
				Message: err.Error(),
			}
		}
		return nil, nil, data
	}

}
