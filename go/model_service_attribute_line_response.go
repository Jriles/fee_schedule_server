/*
 * CorpFees
 *
 * API for the Corp Fees central.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package fee_schedule_server

type ServiceAttributeLineResponse struct {

	Id string `json:"id"`

	AttributeTitle string `json:"attribute_title,omitempty"`

	ServiceAttributeValues []ServiceAttributeValue `json:"service_attribute_values"`

	AttributeId string `json:"attribute_id,omitempty"`
}
