package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type CVE struct {
	ID  int    `json:"id"`
	CVE string `json:"cve"`
}

func main() {
	con, err := pgx.Connect(context.Background(), "postgres://postgres:1234qwer@localhost:5432/cve")
	if err != nil {
		fmt.Println("erroris : ", err)
		return
	}

	rows, err := con.Query(context.Background(), "SELECT * FROM cves")
	// Create a slice to store all CVE records
	var cves []CVE

	for rows.Next() {
		var cve CVE

		rows.Scan(&cve.ID, &cve.CVE)
		cves = append(cves, cve)
	}

	// Convert slice into JSON
	jsonData, err := json.MarshalIndent(cves, "", "  ")
	if err != nil {
		log.Fatal("JSON error:", err)
	}

	fmt.Println(string(jsonData))

}
