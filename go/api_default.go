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
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	attrValResArrKey      = "attribute_values"
	attrResArrKey         = "attributes"
	serviceResArrKey      = "services"
	serviceVariantsArrKey = "service_variants"
	serviceLinesKey       = "service_lines"
)

// CreateAttribute -
func CreateAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var requestBody Attribute
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		sqlStatement := `
		INSERT INTO attributes (title)
		VALUES ($1)
		RETURNING id
		`
		id := ""
		err := db.QueryRow(sqlStatement, title).Scan(&id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			successfulRes := AttributeResponse{Id: id}
			c.JSON(http.StatusCreated, successfulRes)
		}
	}
}

// CreateAttributeValue -
func CreateAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var requestBody AttributeValue
	attributeId := c.Param("attributeId")
	if err := c.BindJSON(&requestBody); err != nil || attributeId == "" {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		sqlStatement := `
		INSERT INTO attribute_values (title, attribute_id)
		VALUES ($1, $2)
		RETURNING id
		`
		id := ""
		err := db.QueryRow(sqlStatement, title, attributeId).Scan(&id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			successfulRes := AttributeResponse{Id: id}
			c.JSON(http.StatusCreated, successfulRes)
		}
	}
}

// CreateService -
func CreateService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var requestBody CreateServiceSchema
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		sqlStatement := `
		INSERT INTO services (title)
		VALUES ($1)
		RETURNING id
		`
		id := ""
		err := db.QueryRow(sqlStatement, title).Scan(&id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			successfulRes := CreateServiceResponse{Id: id}
			c.JSON(http.StatusCreated, successfulRes)
		}
	}
}

// CreateServiceAttributeValue - create a new service attribute value (not an attribute value.) This only applies to the service listed in the path.
func CreateServiceAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var requestBody CreateServiceAttributeValueSchema
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		lineId := c.Param("lineId")
		attributeValueId := requestBody.AttributeValueId

		sqlStatement := `
		INSERT INTO service_attribute_values (line_id, attribute_value_id)
		VALUES ($1, $2)
		RETURNING id
		`

		id := ""
		err := db.QueryRow(sqlStatement, lineId, attributeValueId).Scan(&id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			successfulRes := CreateServiceResponse{Id: id}
			c.JSON(http.StatusCreated, successfulRes)
		}
	}
}

func CreateServiceAttributeLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
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
	} else {
		successfulRes := AttributeResponse{Id: id}
		c.JSON(http.StatusCreated, successfulRes)
	}
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
	} else {
		fee := requestBody.Fee
		serviceId := requestBody.ServiceId
		serviceAttributeValueIds := requestBody.ServiceAttributeValueIds
		sqlStatement := `
		INSERT INTO service_variants (service_id, fee)
		VALUES ($1, $2)
		RETURNING id
		`
		id := ""
		err := db.QueryRow(sqlStatement, serviceId, fee).Scan(&id)

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
			} else {
				c.JSON(http.StatusCreated, gin.H{})
			}
		}

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		} else {
			successfulRes := VariantCreatedResponse{Id: id}
			c.JSON(http.StatusCreated, successfulRes)
		}
	}
}

// DeleteAttribute -
func DeleteAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	attributeId := c.Param("attributeId")
	stmt, err := db.Prepare("DELETE FROM attributes WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(attributeId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

// DeleteAttributeValue -
func DeleteAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	valueID := c.Param("valueId")
	stmt, err := db.Prepare("DELETE FROM attribute_values WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(valueID)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

// DeleteService -
func DeleteService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	serviceId := c.Param("serviceId")
	stmt, err := db.Prepare("DELETE FROM services WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(serviceId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

// DeleteServiceAttributeValue - Delete a service attribute value. valueId here is the service attribute value id NOT the attribute value id.
func DeleteServiceAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	valueId := c.Param("valueId")
	stmt, err := db.Prepare("DELETE FROM service_attribute_values WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(valueId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func DeleteServiceAttributeLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	lineId := c.Param("lineId")
	stmt, err := db.Prepare("DELETE FROM service_attribute_lines WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(lineId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

func DeleteVariant(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	variantId := c.Param("variantId")
	stmt, err := db.Prepare("DELETE FROM service_variants WHERE id=$1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	_, err = stmt.Exec(variantId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusNoContent, gin.H{})
	}
}

// GetAllAttributeValues -
func GetAllAttributeValues(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	attributeId := c.Param("attributeId")
	rows, err := db.Query(
		"SELECT * FROM attribute_values WHERE attribute_id=$1",
		attributeId)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var attrValResArr []AttributeValueResponse
	for rows.Next() {
		var attrValRes AttributeValueResponse
		var attrId string
		err := rows.Scan(&attrValRes.Id, &attrValRes.Title, &attrId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		attrValResArr = append(attrValResArr, attrValRes)
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			attrValResArrKey: attrValResArr,
		})
	}
}

// GetAllAttributes -
func GetAllAttributes(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	rows, err := db.Query(
		"SELECT * FROM attributes")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var attrResArr []AttributeResponse
	for rows.Next() {
		var attrRes AttributeResponse
		err := rows.Scan(&attrRes.Title, &attrRes.Id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			attrResArr = append(attrResArr, attrRes)
		}
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			attrResArrKey: attrResArr,
		})
	}
}

// GetAllServices -
func GetAllServices(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	rows, err := db.Query(
		"SELECT * FROM services")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	var serviceResArr []ServiceResponse
	for rows.Next() {
		var serviceRes ServiceResponse
		err := rows.Scan(&serviceRes.Title, &serviceRes.Id)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			serviceResArr = append(serviceResArr, serviceRes)
		}
	}

	defer rows.Close()
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, gin.H{
			serviceResArrKey: serviceResArr,
		})
	}
}

// GetFee - Retrieve the fee and other information for a particular service variant, ie. (Amended and Restated Articles in Delaware, 1 Day)
func GetVariants(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	var variantsResponse []VariantResponse
	serviceAttributeValueIds := c.Request.URL.Query()["serviceAttributeValueIds[]"]
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
		}

		variantErr := db.QueryRow(
			"SELECT fee FROM service_variants WHERE id=$1",
			variantResponse.Id).Scan(&variantResponse.Fee)

		if variantErr != nil {
			log.Print(variantErr)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		variantsResponse = append(variantsResponse, variantResponse)
	} else {
		//SELECT ALL VARIANTS
		rows, err := db.Query(
			"SELECT * FROM service_variants")
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}

		for rows.Next() {
			var variantRes VariantResponse
			err := rows.Scan(&variantRes.Id, &variantRes.ServiceId, &variantRes.Fee)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
			} else {
				var serviceAttrVals []string
				var err error
				serviceAttrVals, err = GetServiceVariantAttributeValues(db, variantRes.Id)
				if err != nil {
					log.Print(err)
					c.JSON(http.StatusInternalServerError, gin.H{})
					return
				} else {
					variantRes.ServiceAttributeVals = serviceAttrVals
					variantsResponse = append(variantsResponse, variantRes)
				}
			}
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
	}
	var serviceRes ServiceResponse
	serviceId := c.Param("serviceId")
	err := db.QueryRow(
		"SELECT * FROM services WHERE id=$1",
		serviceId).Scan(&serviceRes.Id, &serviceRes.Title)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, serviceRes)
	}
}

func GetServiceAttrLine(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	serviceId := c.Param("serviceId")
	attributeId := c.Param("attributeId")
	// get the line's ID
	var serviceAttrLineRes ServiceAttributeLineResponse
	err := db.QueryRow(
		"SELECT * FROM service_attribute_lines WHERE service_id=$1 AND attribute_id=$2",
		serviceId,
		attributeId,
	).Scan(&serviceAttrLineRes.Id, &serviceId, &attributeId)

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
	} else {
		c.JSON(http.StatusOK, serviceAttrLineRes)
	}
}

func GetServiceAttrLines(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	serviceId := c.Param("serviceId")
	lines, err := db.Query(
		"SELECT * FROM service_attribute_lines WHERE service_id=$1",
		serviceId,
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
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
		} else {
			err := db.QueryRow(
				"SELECT * FROM attributes WHERE id=$1",
				attrId).Scan(&lineRes.AttributeTitle, &attrId)

			serviceAttrVals, err := db.Query(
				"SELECT attribute_value_id FROM service_attribute_values WHERE line_id=$1",
				lineRes.Id,
			)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
			}

			var attrValIds []string
			for serviceAttrVals.Next() {
				var attrValId string
				err := serviceAttrVals.Scan(&attrValId)
				if err != nil {
					log.Print(err)
					c.JSON(http.StatusInternalServerError, gin.H{})
				} else {
					attrValIds = append(attrValIds, attrValId)
				}
			}

			rows, err := db.Query(
				"SELECT * FROM attribute_values WHERE id = ANY($1)",
				pq.Array(attrValIds),
			)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
			}

			var serviceAttrValResArr []ServiceAttributeValue
			for rows.Next() {
				var attrValRes ServiceAttributeValue
				var attrId string
				err := rows.Scan(&attrValRes.Id, &attrValRes.ValueTitle, &attrId)
				if err != nil {
					log.Print(err)
					c.JSON(http.StatusInternalServerError, gin.H{})
				} else {
					serviceAttrValResArr = append(serviceAttrValResArr, attrValRes)
				}
			}
			lineRes.ServiceAttributeValues = serviceAttrValResArr
			linesRes = append(linesRes, lineRes)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		serviceLinesKey: linesRes,
	})
}

// GetServiceAttrVals - Get all the service attribute values for a particular service attr line.
func GetServiceAttrLineVals(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	lineId := c.Param("lineId")
	serviceAttrValIds, err := db.Query(
		"SELECT id FROM service_attribute_values WHERE line_id=$1",
		lineId,
	)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	var serviceAttrVals []ServiceAttributeValue
	for serviceAttrValIds.Next() {
		var serviceAttrValId string
		var serviceAttrVal ServiceAttributeValue
		err := serviceAttrValIds.Scan(&serviceAttrValId)

		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		} else {
			serviceAttrVal.Id = serviceAttrValId
			serviceAttrVal.ValueTitle, err = GetAttributeValueTitleFromServiceAttrId(db, serviceAttrValId)
			if err != nil {
				log.Print(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			} else {
				serviceAttrVals = append(serviceAttrVals, serviceAttrVal)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		attrValResArrKey: serviceAttrVals,
	})
}

//TODO get individual attribute value using id
// maybe, build frontend first

// UpdateAttribute -
func UpdateAttribute(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	attributeId := c.Param("attributeId")
	var requestBody Attribute
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		stmt, err := db.Prepare(
			"UPDATE attributes SET title = $1 WHERE id = $2",
		)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		_, err = stmt.Exec(title, attributeId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

// UpdateAttributeValue -
func UpdateAttributeValue(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	attrValId := c.Param("valueId")
	var requestBody AttributeValue
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		stmt, err := db.Prepare(
			"UPDATE attribute_values SET title = $1 WHERE id = $2",
		)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		_, err = stmt.Exec(title, attrValId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

// UpdateService -
func UpdateService(c *gin.Context) {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	serviceId := c.Param("serviceId")
	var requestBody Attribute
	if err := c.BindJSON(&requestBody); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{})
	} else {
		title := requestBody.Title
		stmt, err := db.Prepare(
			"UPDATE services SET title = $1 WHERE id = $2",
		)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		}
		_, err = stmt.Exec(title, serviceId)
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{})
		} else {
			c.JSON(http.StatusNoContent, gin.H{})
		}
	}
}

func GetServiceVariantAttributeValues(db *sql.DB, variantId string) (serviceAttrVals []string, err error) {
	rows, err := db.Query(
		"SELECT service_attribute_value_id FROM service_variant_combination WHERE service_variant_id = $1", variantId)
	if err != nil {
		return nil, errors.New("query failed to find variant in combination table")
	}
	for rows.Next() {
		var serviceAttrValId string
		err := rows.Scan(&serviceAttrValId)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		serviceAttrVal, err := GetAttributeValueTitleFromServiceAttrId(db, serviceAttrValId)
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
	} else {
		return attrValTitle, nil
	}
}
