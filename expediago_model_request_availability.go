package models

type PropertyAvailabilityRequestModel struct {
	Checkin            string `url:"checkin"`
	Checkout           string `url:"checkout"`
	Currency           string `url:"currency"`
	CountryCode        string `url:"country_code"`
	Language           string `url:"language"`
	Occupancy          string `url:"occupancy"`
	PropertyId         string `url:"property_id"`
	RatePlanCount      string `url:"rate_plan_count"`
	SalesChannel       string `url:"sales_channel"`
	SalesEnvironment   string `url:"sales_environment"`
	AmenityCategory    string `url:"amenity_category,omitempty"`
	Exclusion          string `url:"exclusion,omitempty"`
	Filter             string `url:"filter,omitempty"`
	Include            string `url:"include,omitempty"`
	RateOption         string `url:"rate_option,omitempty"`
	TravelPurpose      string `url:"travel_purpose,omitempty"`
	BillingTerms       string `url:"billing_terms,omitempty"`
	PaymentTerms       string `url:"payment_terms,omitempty"`
	PartnerPointOfSale string `url:"partner_point_of_sale,omitempty"`
	PlatformName       string `url:"platform_name,omitempty"`
}
