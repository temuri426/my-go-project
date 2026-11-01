package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type CVE struct {
	ID  int    `json:"id"`
	CVE string `json:"cve"`
}

func main() {

	// Connection to database
	db_url := "postgres://postgres:1234qwer@test-db:5432/cve"

	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer conn.Close(context.Background())

	// Define POST handler, check id requets is POST
	h1 := func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			// Decode received JSON body. POST request will be in JSON
			var cve CVE

			err := json.NewDecoder(r.Body).Decode(&cve)
			if err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			// Insert values into DB
			commandTag, err := conn.Exec(r.Context(), "INSERT INTO cves (cve) VALUES ($1)", cve.CVE)

			if err != nil {
				http.Error(w, "Database insert error", http.StatusInternalServerError)
				return
			}

			fmt.Println("Rows affected:", commandTag.RowsAffected())

			w.WriteHeader(http.StatusCreated)
			fmt.Fprintln(w, "CVE inserted successfully")

			//return ID if the new entry
			var id int
			var cveValue string
			conn.QueryRow(r.Context(), "SELECT * FROM cves WHERE cve = $1", cve.CVE).Scan(&id, &cveValue)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{
				"id":  id,
				"cve": cveValue,
			})

		}

		if r.Method == "GET" {
			rows, err := conn.Query(context.Background(), "SELECT * FROM cves")
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

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonData)
		}

		//show hostname only
			cmd := exec.Command("hostname")
			stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
		log.Fatal(err)
		}

		h2 := func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, string(stdoutStderr))
		}

	// Register handler
	http.HandleFunc("/cve", h1)    //for inserting and listing all entries
	http.HandleFunc("/", h2)       //for getting hostname
	http.HandleFunc("/cve/id", h3) //for deleting an entrie

	// Start HTTP server
	http.ListenAndServe(":8090", nil)

	}
