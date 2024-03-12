package main

import (
	db "Assignment-2-Golang/database"
	"Assignment-2-Golang/routers"
)

func main() {
	db.Init()
	routers.ServerOn().Run(":8080")
}
