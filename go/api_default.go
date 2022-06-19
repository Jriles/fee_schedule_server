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
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	qualdevlabs_auth_go_client "github.com/Jriles/QualDevLabsAuthGoClient"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	attrValResArrKey      = "attribute_values"
	attrResArrKey         = "attributes"
	serviceResArrKey      = "services"
	serviceVariantsArrKey = "service_variants"
	serviceLinesKey       = "service_lines"
	variantsPerPage       = 50
)

type contextKey struct {
	name string
}

var apiKeyCtxKey = &contextKey{"x-api-key"}

// CreateAttribute -
func CreateAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var requestBody Attribute
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	title := requestBody.Title
	sqlStatement := `
	INSERT INTO attributes (title)
	VALUES ($1)
	RETURNING id
	`
	id := ""
	queryErr := db.QueryRow(sqlStatement, title).Scan(&id)
	if queryErr != nil {
		log.Print(queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	successfulRes := AttributeResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

// CreateAttributeValue -
func CreateAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var requestBody AttributeValue
	attributeId := c.Param("attributeId")
	err := c.BindJSON(&requestBody)
	if err != nil || attributeId == "" {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	title := requestBody.Title
	sqlStatement := `
	INSERT INTO attribute_values (title, attribute_id)
	VALUES ($1, $2)
	RETURNING id
	`
	id := ""
	queryErr := db.QueryRow(sqlStatement, title, attributeId).Scan(&id)
	if queryErr != nil {
		log.Print(queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	successfulRes := AttributeResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

// CreateService -
func CreateService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var requestBody CreateServiceSchema
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	title := requestBody.Title
	sqlStatement := `
	INSERT INTO services (title)
	VALUES ($1)
	RETURNING id
	`
	id := ""
	queryErr := db.QueryRow(sqlStatement, title).Scan(&id)
	if queryErr != nil {
		log.Print(queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	successfulRes := CreateServiceResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

// CreateServiceAttributeValue - create a new service attribute value (not an attribute value.) This only applies to the service listed in the path.
func CreateServiceAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var requestBody CreateServiceAttributeValueSchema
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	lineId := c.Param("lineId")
	attributeValueId := requestBody.AttributeValueId
	sqlStatement := `
	INSERT INTO service_attribute_values (line_id, attribute_value_id)
	VALUES ($1, $2)
	RETURNING id
	`
	id := ""
	queryErr := db.QueryRow(sqlStatement, lineId, attributeValueId).Scan(&id)
	if queryErr != nil {
		log.Print(queryErr)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	successfulRes := CreateServiceResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

func CreateServiceAttributeLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	serviceId := c.Param("serviceId")
	attributeId := c.Param("attributeId")
	sqlStatement := `
	INSERT INTO service_attribute_lines (service_id, attribute_id)
	VALUES ($1, $2)
	RETURNING id
	`
	id := ""
	err := db.QueryRow(sqlStatement, serviceId, attributeId).Scan(&id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	successfulRes := AttributeResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

func CreateVariant(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var requestBody CreateServiceVariantSchema
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	stateCost := requestBody.StateCost
	serviceId := requestBody.ServiceId
	serviceAttributeValueIds := requestBody.ServiceAttributeValueIds
	perPageStateCost := requestBody.PerPageStateCost
	sqlStatement := `
	INSERT INTO service_variants (service_id, state_cost, service_attribute_value_ids, per_page_state_cost)
	VALUES ($1, $2, $3, $4)
	RETURNING id
	`
	id := ""
	err := db.QueryRow(sqlStatement, serviceId, stateCost, pq.Array(serviceAttributeValueIds), perPageStateCost).Scan(&id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	for _, serviceAttrValId := range serviceAttributeValueIds {
		stmt, err := db.Prepare(
			`INSERT INTO service_variant_combination
			(service_variant_id, service_attribute_value_id) 
			VALUES ($1, $2)`,
		)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		_, err = stmt.Exec(id, serviceAttrValId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	successfulRes := VariantCreatedResponse{Id: id}
	c.JSON(http.StatusCreated, successfulRes)
}

// DeleteAttribute -
func DeleteAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	attributeId := c.Param("attributeId")
	stmt, err := db.Prepare("DELETE FROM attributes WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(attributeId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// DeleteAttributeValue -
func DeleteAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	valueID := c.Param("valueId")
	stmt, err := db.Prepare("DELETE FROM attribute_values WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(valueID)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// DeleteService -
func DeleteService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	serviceId := c.Param("serviceId")
	stmt, err := db.Prepare("DELETE FROM services WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(serviceId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// DeleteServiceAttributeValue - Delete a service attribute value. valueId here is the service attribute value id NOT the attribute value id.
func DeleteServiceAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	valueId := c.Param("valueId")
	stmt, err := db.Prepare("DELETE FROM service_attribute_values WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(valueId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteServiceAttributeLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	lineId := c.Param("lineId")
	stmt, err := db.Prepare("DELETE FROM service_attribute_lines WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(lineId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func DeleteVariant(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	variantId := c.Param("variantId")
	stmt, err := db.Prepare("DELETE FROM service_variants WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(variantId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// GetAllAttributeValues -
func GetAllAttributeValues(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	attributeId := c.Param("attributeId")
	rows, err := db.Query(
		"SELECT * FROM attribute_values WHERE attribute_id=$1",
		attributeId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var attrValResArr []AttributeValueResponse
	for rows.Next() {
		var attrValRes AttributeValueResponse
		var attrId string
		err := rows.Scan(&attrValRes.Id, &attrValRes.Title, &attrId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		attrValResArr = append(attrValResArr, attrValRes)
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	c.JSON(http.StatusOK, gin.H{
		attrValResArrKey: attrValResArr,
	})
}

// GetAllAttributes -
func GetAllAttributes(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	rows, err := db.Query(
		"SELECT * FROM attributes")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var attrResArr []AttributeResponse
	for rows.Next() {
		var attrRes AttributeResponse
		err := rows.Scan(&attrRes.Title, &attrRes.Id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		attrResArr = append(attrResArr, attrRes)
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		attrResArrKey: attrResArr,
	})
}

func GetAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	attributeId := c.Param("attributeId")
	var attrRes AttributeResponse
	err := db.QueryRow(
		"SELECT * FROM attributes WHERE id=$1", attributeId).Scan(&attrRes.Title, &attrRes.Id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, attrRes)
}

// GetAllServices -
func GetAllServices(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	rows, err := db.Query(
		"SELECT * FROM services")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var serviceResArr []ServiceResponse
	for rows.Next() {
		var serviceRes ServiceResponse
		err := rows.Scan(&serviceRes.Title, &serviceRes.Id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		serviceResArr = append(serviceResArr, serviceRes)
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		serviceResArrKey: serviceResArr,
	})
}

func GetVariants(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var variantsResponse []VariantResponse
	serviceAttributeValueIds := c.Request.URL.Query()["serviceAttributeValueIds[]"]
	pageCount := c.Request.URL.Query()["filing_page_count"]
	combinationLen := len(serviceAttributeValueIds)
	if combinationLen > 0 {
		//select a specific variant
		var variantResponse VariantResponse
		serviceAttributeValueIdsPqArr := pq.Array(serviceAttributeValueIds)

		combinationErr := db.QueryRow(
			`SELECT service_variant_id
			FROM service_variant_combination 
			WHERE service_attribute_value_id = ANY($1) 
			GROUP BY service_variant_id HAVING COUNT(*) >= $2
			`,
			serviceAttributeValueIdsPqArr, combinationLen,
		).Scan(
			&variantResponse.Id,
		)
		if combinationErr != nil {
			log.Print(combinationErr)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		var pageCountInt int
		var stateCostSubtotal int
		var perPageStateCost int
		variantErr := db.QueryRow(
			"SELECT state_cost, per_page_state_cost FROM service_variants WHERE id=$1",
			variantResponse.Id).Scan(&stateCostSubtotal, &perPageStateCost)
		if variantErr != nil {
			log.Print(variantErr)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		if len(pageCount) > 0 {
			var err error
			pageCountInt, err = strconv.Atoi(pageCount[0])
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		variantResponse.StateCost = CalculateVariantStateCost(stateCostSubtotal, perPageStateCost, pageCountInt)
		variantsResponse = append(variantsResponse, variantResponse)
	} else {
		//SELECT ALL VARIANTS
		pageNumStr := c.Query("page_number")
		pageNum := 1
		if pageNumStr != "" {
			var err error
			pageNum, err = strconv.Atoi(pageNumStr)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
		}

		offset := (pageNum - 1) * variantsPerPage
		rows, err := db.Query(
			`SELECT * FROM service_variants LIMIT $1 OFFSET $2
			`, variantsPerPage, offset)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		for rows.Next() {
			var variantRes VariantResponse
			var serviceAttrValIds []string
			err := rows.Scan(&variantRes.Id, &variantRes.ServiceId, &variantRes.StateCost, (*pq.StringArray)(&serviceAttrValIds), &variantRes.PerPageStateCost)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}

			serviceTitleErr := db.QueryRow(
				"SELECT title FROM services WHERE id=$1",
				variantRes.ServiceId).Scan(&variantRes.ServiceName)
			if serviceTitleErr != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}

			serviceAttrVals, err := GetServiceVariantAttributeValues(db, serviceAttrValIds)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			variantRes.ServiceAttributeVals = serviceAttrVals
			variantsResponse = append(variantsResponse, variantRes)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		serviceVariantsArrKey: variantsResponse,
	})
}

// GetService -
func GetService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var serviceRes ServiceResponse
	serviceId := c.Param("serviceId")
	err := db.QueryRow(
		"SELECT * FROM services WHERE id=$1",
		serviceId).Scan(&serviceRes.Title, &serviceRes.Id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, serviceRes)
}

func GetServiceAttrLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	lineId := c.Param("lineId")
	// get the line's ID
	var serviceAttrLineRes ServiceAttributeLineResponse
	var serviceId string
	err := db.QueryRow(
		"SELECT * FROM service_attribute_lines WHERE id=$1",
		lineId,
	).Scan(&serviceAttrLineRes.Id, &serviceId, &serviceAttrLineRes.AttributeId)

	attrErr := db.QueryRow(
		"SELECT * FROM attributes WHERE id=$1",
		serviceAttrLineRes.AttributeId).Scan(&serviceAttrLineRes.AttributeTitle, &serviceAttrLineRes.AttributeId)
	if attrErr != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	serviceAttrLineRes.ServiceAttributeValues, err = GetServiceAttrLineVals(db, serviceAttrLineRes.Id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusOK, serviceAttrLineRes)
}

func GetServiceAttrLines(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	serviceId := c.Param("serviceId")
	lines, err := db.Query(
		"SELECT * FROM service_attribute_lines WHERE service_id=$1",
		serviceId,
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	var linesRes []ServiceAttributeLineResponse
	for lines.Next() {
		var lineRes ServiceAttributeLineResponse
		var serviceId string
		var attrId string
		err := lines.Scan(&lineRes.Id, &serviceId, &attrId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		attrErr := db.QueryRow(
			"SELECT * FROM attributes WHERE id=$1",
			attrId).Scan(&lineRes.AttributeTitle, &attrId)
		if attrErr != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		lineRes.ServiceAttributeValues, err = GetServiceAttrLineVals(db, lineRes.Id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		linesRes = append(linesRes, lineRes)
	}

	c.JSON(http.StatusOK, gin.H{
		serviceLinesKey: linesRes,
	})
}

//TODO get individual attribute value using id
// maybe, build frontend first

// UpdateAttribute -
func UpdateAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	attributeId := c.Param("attributeId")
	var requestBody Attribute
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	title := requestBody.Title
	stmt, err := db.Prepare(
		"UPDATE attributes SET title = $1 WHERE id = $2",
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(title, attributeId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// UpdateAttributeValue -
func UpdateAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	attrValId := c.Param("valueId")
	var requestBody AttributeValue
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	title := requestBody.Title
	stmt, err := db.Prepare(
		"UPDATE attribute_values SET title = $1 WHERE id = $2",
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(title, attrValId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// UpdateService -
func UpdateService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	serviceId := c.Param("serviceId")
	var requestBody Attribute
	err := c.BindJSON(&requestBody)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	title := requestBody.Title
	stmt, err := db.Prepare(
		"UPDATE services SET title = $1 WHERE id = $2",
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_, err = stmt.Exec(title, serviceId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func LoginUser(c *gin.Context) {
	orgId := os.Getenv("AUTH_ORG_ID") // string | the org's UUID (unique)
	appId := os.Getenv("AUTH_APP_ID") // string | the app's UUID (unique)
	authApiKey := os.Getenv("AUTH_API_KEY")
	authApiKeyStruct := qualdevlabs_auth_go_client.APIKey{
		Key: authApiKey,
	}

	var loginDetails LoginSchema
	err := c.BindJSON(&loginDetails)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	loginSchema := *qualdevlabs_auth_go_client.NewLoginSchema(
		loginDetails.Username,
		loginDetails.Password,
		loginDetails.RememberMe,
	)

	configuration := qualdevlabs_auth_go_client.NewConfiguration()
	api_client := qualdevlabs_auth_go_client.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), qualdevlabs_auth_go_client.ContextAPIKeys, map[string]qualdevlabs_auth_go_client.APIKey{
		"apiKeyHeader": authApiKeyStruct,
	})
	resp, r, err := api_client.DefaultApi.CreateUserSession(ctx, orgId, appId).LoginSchema(loginSchema).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateUserSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	loginRes := SuccessfulLoginResponse{
		resp.Token,
	}
	c.JSON(http.StatusOK, loginRes)
}

func GetServiceVariantAttributeValues(db *sql.DB, serviceAttrValIds []string) (serviceAttrVals []string, err error) {
	for _, serviceAttrValId := range serviceAttrValIds {
		serviceAttrVal, err := GetAttributeValueTitleFromServiceAttrId(db, string(serviceAttrValId))
		if err != nil {
			return nil, errors.New(err.Error())
		}
		serviceAttrVals = append(serviceAttrVals, serviceAttrVal)
	}
	return serviceAttrVals, nil
}

func GetAttributeValueTitleFromServiceAttrId(db *sql.DB, serviceAttrValId string) (attrValTitle string, err error) {
	sqlStatement := `
	WITH attr_value_ids AS (
		SELECT attribute_value_id FROM service_attribute_values WHERE id = $1
	)
	SELECT title FROM attribute_values WHERE id IN (SELECT attribute_value_id FROM attr_value_ids)
	`
	queryErr := db.QueryRow(sqlStatement, serviceAttrValId).Scan(&attrValTitle)
	if queryErr != nil {
		return "", errors.New(err.Error())
	}
	return attrValTitle, nil
}

func GetServiceAttrLineVals(db *sql.DB, lineId string) (serviceAttrVals []ServiceAttributeValue, err error) {
	serviceAttrValIds, err := db.Query(
		"SELECT id FROM service_attribute_values WHERE line_id=$1",
		lineId,
	)
	if err != nil {
		log.Print(err)
		return nil, errors.New(err.Error())
	}
	for serviceAttrValIds.Next() {
		var serviceAttrValId string
		var serviceAttrVal ServiceAttributeValue
		err := serviceAttrValIds.Scan(&serviceAttrValId)
		if err != nil {
			log.Print(err)
			return nil, errors.New(err.Error())
		}

		serviceAttrVal.Id = serviceAttrValId
		serviceAttrVal.ValueTitle, err = GetAttributeValueTitleFromServiceAttrId(db, serviceAttrValId)
		if err != nil {
			log.Print(err)
			return nil, errors.New(err.Error())
		}
		serviceAttrVals = append(serviceAttrVals, serviceAttrVal)
	}
	return serviceAttrVals, nil
}

func CalculateVariantStateCost(stateCostSubtotal int, perPageStateCost int, pageCount int) (totalStateCost int) {
	pagesCost := perPageStateCost * pageCount
	totalStateCost = stateCostSubtotal * pagesCost
	return totalStateCost
}
