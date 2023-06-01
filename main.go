package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/zxc10110/Fruit_Market/Config"
	Routes "github.com/zxc10110/Fruit_Market/Controller"
	"github.com/zxc10110/Fruit_Market/Model"
)

var err error

func main() {
	fmt.Println("Run Swagger : http://localhost:8080/swagger/index.html")
	fmt.Println()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Model.Fruit{})
	r := Routes.SetupRouter()
	//running
	port := ":8000" // Specify the desired port number
	fmt.Printf("Server listening on port %s\n", port)
	r.Run(port)
}
