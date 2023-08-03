package expediago

const expediaApiUrlDevelopment = "https://test.ean.com"
const expediaApiUrlProduction = "https://api.ean.com"

var expediaKey string
var expediaSecret string
var expediaBaseUrl string

func SetExpediaKeyAndSecret(apiKey string, apiSecret string) {
	expediaKey = apiKey
	expediaSecret = apiSecret
}

func EnableForProduction() {
	expediaBaseUrl = expediaApiUrlProduction
}

func GetApiKey() string {
	return expediaKey
}

func GetApiSecret() string {
	return expediaSecret
}

func GetApiBaseUrl() string {
	if expediaBaseUrl != "" {
		return expediaBaseUrl
	} else {
		return expediaApiUrlDevelopment
	}
}
