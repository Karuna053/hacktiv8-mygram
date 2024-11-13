package main

import (
	"mygram/database"
	"mygram/routers"
)

func main() {
	database.InitDB()
	routers.StartRouter().Run(":8080")
}
