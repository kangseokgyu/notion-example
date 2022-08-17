package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dstotijn/go-notion"
)

var (
	NOTION_KEY         string = os.Getenv("NOTION_KEY")
	NOTION_DATABASE_ID string = os.Getenv("NOTION_DATABASE_ID")
)

func main() {
	client := notion.NewClient(NOTION_KEY)

	database, err := client.FindDatabaseByID(context.Background(), NOTION_DATABASE_ID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(database)

	fmt.Println("")

	res, err := client.QueryDatabase(context.Background(), NOTION_DATABASE_ID, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Results)
	for _, r := range res.Results {
		fmt.Println(r.ID)
		fmt.Println(r.Properties)
	}
}
