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

	//delete vaules from DB

	commandTag, err := conn.Exec(context.Background(), "DELETE FROM cves WHERE id = $1", cve.ID)
	if err != nil {
		return
	}

	fmt.Println("data is deleted successfuly")
	fmt.Println("rows affected:", commandTag.RowsAffected())

}
