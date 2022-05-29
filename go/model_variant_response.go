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

	Fee float32 `json:"fee"`

	ServiceId string `json:"service_id"`

	ServiceAttributeVals []string `json:"service_attribute_vals,omitempty"`

	ServiceName string `json:"service_name,omitempty"`
}
