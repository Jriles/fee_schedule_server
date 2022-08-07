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
	countryCode := requestBody.IsoCountryCode
	currencyCode := requestBody.IsoCurrencyCode
	sqlStatement := `
	INSERT INTO service_variants (service_id, state_cost, service_attribute_value_ids, country_code, currency_code)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`
	id := ""
	err := db.QueryRow(
		sqlStatement,
		serviceId,
		stateCost,
		pq.Array(serviceAttributeValueIds),
		countryCode,
		currencyCode,
	).Scan(&id)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
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
	defer stmt.Close()

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
	defer stmt.Close()

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
	defer stmt.Close()

	_, err = stmt.Exec(serviceId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// DeleteServiceAttributeValue - Delete a service attribute value. valueId here is the service attribute value id NOT the attribute value id.
func DeleteServiceAttributeValueHandler(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	valueId := c.Param("valueId")

	err := DeleteServiceAttributeValue(db, valueId)
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

	// delete all related service attribute values
	// cascading is not option here (we want to delete assoc. variants as well and there is
	// a many2many relation between service attribute values and service variants).
	rows, err := db.Query(
		"SELECT id FROM service_attribute_values WHERE line_id=$1",
		lineId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer rows.Close()

	for rows.Next() {
		// service attribute value Id
		var serviceAttrValId string
		err = rows.Scan(&serviceAttrValId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		//also deletes assoc. variants
		err := DeleteServiceAttributeValue(db, serviceAttrValId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
	}

	stmt, err := db.Prepare("DELETE FROM service_attribute_lines WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer stmt.Close()

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
	defer stmt.Close()

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
	defer rows.Close()

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
	defer rows.Close()

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
	defer rows.Close()

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

	serviceId := c.Query("service_id")
	attributeValueIds := []string{}
	var variantsResponse []VariantResponse
	// step 1: filter service attribute values for the selected attribute values
	// step 2: filter for service variants with selected attribute values/service id
	// step 3: replace the service attribute value ids with attribute value ids, still in an array
	// step 4: replace the attribute value ids with attribute value titles, also in an array, also replace the service id with service title.
	// step 5: done!

	rows, err := db.Query(
		`WITH selected_service_attribute_values AS (
			SELECT array_agg(id) service_attribute_value_id FROM service_attribute_values WHERE 
			CASE 
				WHEN array_length($2::uuid[], 1) > 0 THEN attribute_value_id=ANY($2::uuid[])
				ELSE true
			END
		),
		filtered_variants AS (
			SELECT * FROM service_variants WHERE
			CASE
				WHEN $1::text != '' AND array_length($2::uuid[], 1) > 0
					THEN service_id=$1::uuid AND (service_attribute_value_ids && (SELECT service_attribute_value_id FROM selected_service_attribute_values))
				WHEN $1::text = '' AND array_length($2::uuid[], 1) > 0
					THEN (service_attribute_value_ids && (SELECT service_attribute_value_id FROM selected_service_attribute_values))
				WHEN $1::text != '' AND array_length($2::uuid[], 1) = 0
					THEN service_id=$1::uuid
				ELSE true
			END
		),
		filtered_variants_w_attribute_value_ids AS (
			SELECT 
				ARRAY_AGG(service_attribute_values.attribute_value_id) attribute_value_ids, 
				filtered_variants.id,
				filtered_variants.service_id,
				filtered_variants.state_cost
			FROM service_attribute_values
			INNER JOIN filtered_variants ON service_attribute_values.id=ANY(filtered_variants.service_attribute_value_ids)
			GROUP BY 
				filtered_variants.id,
				filtered_variants.service_id,
				filtered_variants.state_cost
		)
		SELECT 
			filtered_variants_w_attribute_value_ids.id,
			services.title,
			filtered_variants_w_attribute_value_ids.state_cost,
			ARRAY_AGG(attribute_values.title) attribute_value_titles FROM attribute_values
		INNER JOIN filtered_variants_w_attribute_value_ids ON attribute_values.id=ANY(filtered_variants_w_attribute_value_ids.attribute_value_ids)
		INNER JOIN services ON services.id = filtered_variants_w_attribute_value_ids.service_id
		GROUP BY 
			filtered_variants_w_attribute_value_ids.id,
			services.title,
			filtered_variants_w_attribute_value_ids.state_cost
		`,
		serviceId,
		pq.StringArray(attributeValueIds),
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var variantRes VariantResponse
		err := rows.Scan(
			&variantRes.Id,
			&variantRes.ServiceName,
			&variantRes.StateCost,
			(*pq.StringArray)(&variantRes.ServiceAttributeVals),
		)
		if err != nil {
			log.Print(err)
		}
		variantsResponse = append(variantsResponse, variantRes)
	}
	//
	// 		var serviceAttrValIds []string
	// 		err := rows.Scan(
	// 			&variantRes.Id,
	// 			&variantRes.ServiceId,
	// 			&variantRes.StateCost,
	// 			(*pq.StringArray)(&serviceAttrValIds),
	// 			&variantRes.PerPageStateCost,
	// 			&variantRes.IsoCountryCode,
	// 			&variantRes.IsoCountryCode,
	// 		)
	// 		if err != nil {
	// 			log.Print(err)
	// 			c.JSON(http.StatusInternalServerError, gin.H{})
	// 			return
	// 		}

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
	defer stmt.Close()

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
	defer stmt.Close()

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
	defer stmt.Close()

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
		resp.UserId,
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
	//this needs to be done per variant
	//INNER JOIN?
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

func CalculateVariantStateCost(stateCostSubtotal int32, perPageStateCost int32, pageCount int32) (totalStateCost int32) {
	pagesCost := perPageStateCost * pageCount
	totalStateCost = stateCostSubtotal * pagesCost
	return totalStateCost
}

// deletes a given service attribute value AND assoc. service variants.
func DeleteServiceAttributeValue(db *sql.DB, valueId string) (err error) {
	attrValDelStmt, err := db.Prepare("DELETE FROM service_attribute_values WHERE id=$1")
	if err != nil {
		log.Print(err)
		return err
	}
	defer attrValDelStmt.Close()

	_, err = attrValDelStmt.Exec(valueId)
	if err != nil {
		log.Print(err)
		return err
	}

	var serviceVariantIds []string

	rows, variantQueryErr := db.Query(
		"SELECT service_variant_id FROM service_variant_combination WHERE service_attribute_value_id=$1",
		valueId,
	)
	if variantQueryErr != nil {
		log.Print(variantQueryErr)
		return variantQueryErr
	}
	defer rows.Close()

	for rows.Next() {
		var serviceVariantId string
		err = rows.Scan(&serviceVariantId)
		if err != nil {
			log.Print(err)
			return err
		}

		serviceVariantIds = append(serviceVariantIds, serviceVariantId)
	}

	serviceVarDelStmt, err := db.Prepare("DELETE FROM service_variants WHERE id= ANY($1)")
	if err != nil {
		log.Print(err)
		return err
	}
	defer serviceVarDelStmt.Close()

	_, err = serviceVarDelStmt.Exec(pq.Array(serviceVariantIds))
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}
