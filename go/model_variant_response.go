/*
 * CorpFees
 *
 * API for the Corp Fees central.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package fee_schedule_server

type VariantResponse struct {

	Id string `json:"id"`

	StateCost int32 `json:"state_cost"`

	ServiceId string `json:"service_id"`

	ServiceAttributeVals []string `json:"service_attribute_vals,omitempty"`

	ServiceName string `json:"service_name,omitempty"`

	PerPageStateCost int32 `json:"per_page_state_cost,omitempty"`

	// The three letter iso code for the currency of the country for which this service variant applies to.
	IsoCurrencyCode string `json:"iso_currency_code"`

	// The two letter iso (alphabet, not numeric) code representing the country for which this service variant applies to.
	IsoCountryCode string `json:"iso_country_code"`
}
