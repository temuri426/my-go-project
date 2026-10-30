package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type CVE struct {
	ID  int    `json:"id"`
	CVE string `json:"cve"`
}

func main() {
	db_url := "postgres://postgres:1234qwer@localhost:5432/cve"
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Println("error is ", err)
	}
	defer conn.Close(context.Background())

	var cve CVE
	fmt.Println("enter CVE ID:")
	fmt.Scan(&cve.ID)
	fmt.Println("enter CVE name: ")
	fmt.Scan(&cve.CVE)

	//incert vaules into DB

	commandTag, err := conn.Exec(context.Background(), "INSERT INTO cves (id, cve) VALUES ($1, $2)", cve.ID, cve.CVE)
	if err != nil {
		return err
	}
	fmt.Println("data is inserted successfuly")
	fmt.Println("rows affected:", commandTag.rows)

}
