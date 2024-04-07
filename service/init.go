package service

import (
	"context"
	"log"

	"github.com/shplume/zhulong/ent"

	_ "github.com/go-sql-driver/mysql"
)

var client *ent.Client

func init() {
	var err error
	client, err = ent.Open("mysql", "root:123456@tcp(10.0.1.253:3306)/zhulong?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
