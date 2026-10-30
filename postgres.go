package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func main() {
	con, err := pgx.Connect(context.Background(), "postgres://postgres:1234qwer@192.168.7.129:5432/cve")
	if err != nil {
		fmt.Println("erroris : ", err)
		return
	}

	rows, err := con.Query(context.Background(), "SELECT * FROM cves")

	for rows.Next() {
		var id int
		var cve string
		rows.Scan(&id, &cve)
		fmt.Println(id, cve)
	}

}
