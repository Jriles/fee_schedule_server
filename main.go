/*
 * CorpFees
 *
 * API for the Corp Fees central.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"database/sql"
	"log"
	"os"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	sw "github.com/Jriles/fee_schedule_server/go"

	//
	_ "github.com/lib/pq"
)

var db_conn = ""

func main() {
	db_conn = os.Getenv("FEE_SCHEDULE_SERVER_DB_CONN")

	db, err := sql.Open("postgres", db_conn)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	router := sw.NewRouter(db)

	log.Fatal(router.Run(":8080"))
}
