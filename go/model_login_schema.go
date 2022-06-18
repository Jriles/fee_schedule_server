/*
 * CorpFees
 *
 * API for the Corp Fees central.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package fee_schedule_server

type LoginSchema struct {

	Username string `json:"username"`

	Password string `json:"password"`

	RememberMe bool `json:"remember_me"`
}
