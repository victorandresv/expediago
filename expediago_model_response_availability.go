package expediago

import "time"

type ExpediaGoModelResponseAvailability struct {
	PropertyId        string `json:"property_id"`
	Status            string `json:"status"`
	Score             int    `json:"score"`
	UnavailableReason struct {
		Code string `json:"code"`
		Data string `json:"data"`
	} `json:"unavailable_reason,omitempty"`
	Rooms []struct {
		Id       string `json:"id"`
		RoomName string `json:"room_name"`
		Rates    []struct {
			Id                  string `json:"id"`
			Status              string `json:"status"`
			AvailableRooms      int    `json:"available_rooms"`
			Refundable          bool   `json:"refundable"`
			MemberDealAvailable bool   `json:"member_deal_available"`
			SaleScenario        struct {
				Package      bool `json:"package"`
				Member       bool `json:"member"`
				Corporate    bool `json:"corporate"`
				Distribution bool `json:"distribution"`
			} `json:"sale_scenario"`
			MerchantOfRecord string `json:"merchant_of_record"`
			Amenities        struct {
				Field1 struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"1234"`
			} `json:"amenities"`
			Links struct {
				PaymentOptions struct {
					Method string `json:"method"`
					Href   string `json:"href"`
				} `json:"payment_options"`
			} `json:"links"`
			BedGroups struct {
				Field1 struct {
					Links struct {
						PriceCheck struct {
							Method  string    `json:"method"`
							Href    string    `json:"href"`
							Expires time.Time `json:"expires"`
						} `json:"price_check"`
					} `json:"links"`
					Id            string `json:"id"`
					Description   string `json:"description"`
					Configuration []struct {
						Type     string `json:"type"`
						Size     string `json:"size"`
						Quantity int    `json:"quantity"`
					} `json:"configuration"`
				} `json:"12345"`
			} `json:"bed_groups"`
			CancelPenalties []struct {
				Currency string    `json:"currency"`
				Start    time.Time `json:"start"`
				End      time.Time `json:"end"`
				Amount   string    `json:"amount"`
			} `json:"cancel_penalties"`
			NonrefundableDateRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"nonrefundable_date_ranges"`
			OccupancyPricing struct {
				Field1 struct {
					Nightly [][]struct {
						Type     string `json:"type"`
						Value    string `json:"value"`
						Currency string `json:"currency"`
					} `json:"nightly"`
					Stay []struct {
						Type     string `json:"type"`
						Value    string `json:"value"`
						Currency string `json:"currency"`
					} `json:"stay"`
					Totals struct {
						Inclusive struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"inclusive"`
						Exclusive struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"exclusive"`
						InclusiveStrikethrough struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"inclusive_strikethrough"`
						Strikethrough struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"strikethrough"`
						MarketingFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"marketing_fee"`
						GrossProfit struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"gross_profit"`
						MinimumSellingPrice struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"minimum_selling_price"`
						PropertyFees struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"property_fees"`
					} `json:"totals"`
					Fees struct {
						MandatoryFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"mandatory_fee"`
						ResortFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"resort_fee"`
						MandatoryTax struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"mandatory_tax"`
					} `json:"fees"`
				} `json:"2"`
				Field2 struct {
					Nightly [][]struct {
						Type     string `json:"type"`
						Value    string `json:"value"`
						Currency string `json:"currency"`
					} `json:"nightly"`
					Stay []struct {
						Type     string `json:"type"`
						Value    string `json:"value"`
						Currency string `json:"currency"`
					} `json:"stay"`
					Totals struct {
						Inclusive struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"inclusive"`
						Exclusive struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"exclusive"`
						InclusiveStrikethrough struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"inclusive_strikethrough"`
						Strikethrough struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"strikethrough"`
						MarketingFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"marketing_fee"`
						GrossProfit struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"gross_profit"`
						MinimumSellingPrice struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"minimum_selling_price"`
						PropertyFees struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"property_fees"`
					} `json:"totals"`
					Fees struct {
						MandatoryFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"mandatory_fee"`
						ResortFee struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"resort_fee"`
						MandatoryTax struct {
							BillableCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"billable_currency"`
							RequestCurrency struct {
								Value    string `json:"value"`
								Currency string `json:"currency"`
							} `json:"request_currency"`
						} `json:"mandatory_tax"`
					} `json:"fees"`
				} `json:"1-7,10"`
			} `json:"occupancy_pricing"`
			Promotions struct {
				ValueAdds struct {
					Abc struct {
						Id          string `json:"id"`
						Description string `json:"description"`
						Category    string `json:"category"`
						OfferType   string `json:"offer_type"`
						Frequency   string `json:"frequency"`
						PersonCount int    `json:"person_count"`
					} `json:"123abc"`
				} `json:"value_adds"`
				Deal struct {
					Id          string `json:"id"`
					Description string `json:"description"`
				} `json:"deal"`
			} `json:"promotions"`
			CardOnFileLimit struct {
				Value    string `json:"value"`
				Currency string `json:"currency"`
			} `json:"card_on_file_limit"`
			RefundableDamageDeposit struct {
				Value    string `json:"value"`
				Currency string `json:"currency"`
			} `json:"refundable_damage_deposit"`
			Deposits []struct {
				Value    string `json:"value"`
				Due      string `json:"due"`
				Currency string `json:"currency"`
			} `json:"deposits"`
		} `json:"rates"`
	} `json:"rooms,omitempty"`
	Links struct {
		AdditionalRates struct {
			Method string `json:"method"`
			Href   string `json:"href"`
		} `json:"additional_rates"`
	} `json:"links,omitempty"`
}
