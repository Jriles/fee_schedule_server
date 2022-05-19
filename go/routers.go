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
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// Routes is the list of the generated Route.
type Routes []Route

// ApiMiddleware will add the db connection to the context
func ApiMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

// NewRouter returns a new router.
func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(ApiMiddleware(db))
	// TODO: TIGHTEN THIS UP. THIS IS HIGHLY INSECURE
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// Index is the index handler.
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

var routes = Routes{
	{
		"Index",
		http.MethodGet,
		"/",
		Index,
	},

	{
		"CreateAttribute",
		http.MethodPost,
		"/attributes",
		CreateAttribute,
	},

	{
		"CreateAttributeValue",
		http.MethodPost,
		"/attributes/:attributeId/values",
		CreateAttributeValue,
	},

	{
		"CreateService",
		http.MethodPost,
		"/services",
		CreateService,
	},

	{
		"CreateServiceAttributeValue",
		http.MethodPost,
		"/service_attribute_lines/:lineId/values",
		CreateServiceAttributeValue,
	},

	{
		"CreateServiceAttributeLine",
		http.MethodPost,
		"/services/:serviceId/attributes/:attributeId/lines",
		CreateServiceAttributeLine,
	},

	{
		"CreateVariant",
		http.MethodPost,
		"/service_variants",
		CreateVariant,
	},

	{
		"DeleteAttribute",
		http.MethodDelete,
		"/attributes/:attributeId",
		DeleteAttribute,
	},

	{
		"DeleteAttributeValue",
		http.MethodDelete,
		"/attribute_values/:valueId",
		DeleteAttributeValue,
	},

	{
		"DeleteService",
		http.MethodDelete,
		"/services/:serviceId/",
		DeleteService,
	},

	{
		"DeleteServiceAttributeValue",
		http.MethodDelete,
		"/service_attribute_values/:valueId",
		DeleteServiceAttributeValue,
	},

	{
		"DeleteServiceAttributeLine",
		http.MethodDelete,
		"/service_attribute_lines/:lineId",
		DeleteServiceAttributeLine,
	},

	{
		"DeleteVariant",
		http.MethodDelete,
		"/service_variants/:variantId",
		DeleteVariant,
	},

	{
		"GetAllAttributeValues",
		http.MethodGet,
		"/attributes/:attributeId/values",
		GetAllAttributeValues,
	},

	{
		"GetAllAttributes",
		http.MethodGet,
		"/attributes",
		GetAllAttributes,
	},

	{
		"GetAllServices",
		http.MethodGet,
		"/services",
		GetAllServices,
	},

	{
		"GetVariants",
		http.MethodGet,
		"/service_variants/",
		GetVariants,
	},

	{
		"GetService",
		http.MethodGet,
		"/services/:serviceId/",
		GetService,
	},

	{
		"GetServiceAttrLine",
		http.MethodGet,
		"/services/:serviceId/attributes/:attributeId/lines",
		GetServiceAttrLine,
	},

	{
		"GetServiceAttrVals",
		http.MethodGet,
		"/service_attribute_lines/:lineId/values",
		GetServiceAttrLineVals,
	},

	{
		"UpdateAttribute",
		http.MethodPatch,
		"/attributes/:attributeId",
		UpdateAttribute,
	},

	{
		"UpdateAttributeValue",
		http.MethodPatch,
		"/attribute_values/:valueId",
		UpdateAttributeValue,
	},

	{
		"UpdateService",
		http.MethodPatch,
		"/services/:serviceId/",
		UpdateService,
	},
}
