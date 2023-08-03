package models

type ErrorModel struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
	Fields  []struct {
		Name  string `json:"name,omitempty"`
		Type  string `json:"type,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"fields,omitempty"`
	Errors []struct {
		Type    string `json:"type,omitempty"`
		Message string `json:"message,omitempty"`
		Fields  []struct {
			Name  string `json:"name,omitempty"`
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"errors,omitempty"`
}
