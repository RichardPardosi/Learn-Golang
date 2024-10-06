package main

import (
	"fmt"

	"github.com/RichardPardosi/try-web-crud-golang/src/basic-golang/web_golang/database"
)

func main() {
	database.InitDatabase()
	fmt.Println("Hello Word")
}
