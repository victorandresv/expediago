package expediago

import "time"

type ModelResponseAvailability struct {
	PropertyId        string            `json:"property_id"`
	Status            string            `json:"status"`
	Score             int               `json:"score"`
	UnavailableReason unavailableReason `json:"unavailable_reason,omitempty"`
	Rooms             []struct {
		Id       string `json:"id"`
		RoomName string `json:"room_name"`
		Rates    []struct {
			Id                  string            `json:"id"`
			Status              string            `json:"status"`
			AvailableRooms      int               `json:"available_rooms"`
			Refundable          bool              `json:"refundable"`
			MemberDealAvailable bool              `json:"member_deal_available"`
			SaleScenario        saleScenario      `json:"sale_scenario"`
			MerchantOfRecord    string            `json:"merchant_of_record"`
			Amenities           map[string]idName `json:"amenities"`
			Links               struct {
				PaymentOptions methodHref `json:"payment_options"`
			} `json:"links"`
			BedGroups               map[string]bedGroups        `json:"bed_groups"`
			CancelPenalties         []currencyAmount            `json:"cancel_penalties"`
			NonrefundableDateRanges []currencyAmount            `json:"nonrefundable_date_ranges"`
			OccupancyPricing        map[string]occupancyPricing `json:"occupancy_pricing"`
			Promotions              struct {
				ValueAdds map[string]valueAdds `json:"value_adds"`
				Deal      struct {
					Id          string `json:"id"`
					Description string `json:"description"`
				} `json:"deal"`
			} `json:"promotions"`
			CardOnFileLimit         valueCurrency   `json:"card_on_file_limit"`
			RefundableDamageDeposit valueCurrency   `json:"refundable_damage_deposit"`
			Deposits                []valueCurrency `json:"deposits"`
		} `json:"rates"`
	} `json:"rooms,omitempty"`
	Links struct {
		AdditionalRates methodHref `json:"additional_rates"`
	} `json:"links,omitempty"`
}

type unavailableReason struct {
	Code string `json:"code"`
	Data string `json:"data"`
}

type saleScenario struct {
	Package      bool `json:"package"`
	Member       bool `json:"member"`
	Corporate    bool `json:"corporate"`
	Distribution bool `json:"distribution"`
}

type idName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type methodHref struct {
	Method  string    `json:"method"`
	Href    string    `json:"href"`
	Expires time.Time `json:"expires,omitempty"`
}

type valueCurrency struct {
	Value    string `json:"value"`
	Due      string `json:"due,omitempty"`
	Currency string `json:"currency"`
	Type     string `json:"type,omitempty"`
}

type currencyBillableRequest struct {
	BillableCurrency valueCurrency `json:"billable_currency"`
	RequestCurrency  valueCurrency `json:"request_currency"`
}

type occupancyPricing struct {
	Nightly [][]valueCurrency `json:"nightly"`
	Stay    []valueCurrency   `json:"stay"`
	Totals  struct {
		Inclusive              currencyBillableRequest `json:"inclusive"`
		Exclusive              currencyBillableRequest `json:"exclusive"`
		InclusiveStrikethrough currencyBillableRequest `json:"inclusive_strikethrough"`
		Strikethrough          currencyBillableRequest `json:"strikethrough"`
		MarketingFee           currencyBillableRequest `json:"marketing_fee"`
		GrossProfit            currencyBillableRequest `json:"gross_profit"`
		MinimumSellingPrice    currencyBillableRequest `json:"minimum_selling_price"`
		PropertyFees           currencyBillableRequest `json:"property_fees"`
	} `json:"totals"`
	Fees struct {
		MandatoryFee currencyBillableRequest `json:"mandatory_fee"`
		ResortFee    currencyBillableRequest `json:"resort_fee"`
		MandatoryTax currencyBillableRequest `json:"mandatory_tax"`
	} `json:"fees"`
}

type configuration struct {
	Type     string `json:"type"`
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type bedGroups struct {
	Links struct {
		PriceCheck methodHref `json:"price_check"`
	} `json:"links"`
	Id            string          `json:"id"`
	Description   string          `json:"description"`
	Configuration []configuration `json:"configuration"`
}

type currencyAmount struct {
	Currency string    `json:"currency,omitempty"`
	Start    time.Time `json:"start,omitempty"`
	End      time.Time `json:"end,omitempty"`
	Amount   string    `json:"amount,omitempty"`
}

type valueAdds struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Category    string `json:"category"`
	OfferType   string `json:"offer_type"`
	Frequency   string `json:"frequency"`
	PersonCount int    `json:"person_count"`
}
