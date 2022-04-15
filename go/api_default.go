/*
 * CorpFees
 *
 * API for the Corp Fees central.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package fee_schedule_server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateAttribute -
func CreateAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateAttributeValue -
func CreateAttributeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateService -
func CreateService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// CreateServiceAttributeValue - create a new service attribute value (not an attribute value.) This only applies to the service listed in the path. This will automatically create a Service attribute line if none exists, that's why we need the attribute Id.
func CreateServiceAttributeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteAttribute -
func DeleteAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteAttributeValue -
func DeleteAttributeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteService -
func DeleteService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// DeleteServiceAttributeValue - Delete a service attribute value. valueId here is the service attribute value id NOT the attribute value id.
func DeleteServiceAttributeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllAttributeValues -
func GetAllAttributeValues(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllAttributes -
func GetAllAttributes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetAllServices -
func GetAllServices(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetFee - Retrieve the fee and other information for a particular service variant, ie. (Amended and Restated Articles in Delaware, 1 Day)
func GetFee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetService -
func GetService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// GetServiceAttrVals - Get all the service attribute values for a particular attribute.
func GetServiceAttrVals(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateAttribute -
func UpdateAttribute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateAttributeValue -
func UpdateAttributeValue(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// UpdateService -
func UpdateService(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
