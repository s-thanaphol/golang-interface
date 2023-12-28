package main

import (
	"fmt"
	"golang-unit-test/database"
	"golang-unit-test/database/filelog"
	"golang-unit-test/input"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	databaseCase := os.Getenv("DB_CASE")
	var db database.Database
	if databaseCase == "File" {
		filelogDb := filelog.FileLog{}
		db = &filelogDb
	} else if databaseCase == "Redis" {
		// redis
	}

	nowCount,err := db.Get()
	if err != nil {
		return
	}
	inPut,err := input.Init()
	if err != nil {
		return
	}
	for {
		new, end, err := inPut.ReadOneInput()
		if err != nil {
			return
		} else if end == true {
			break
		}
		nowCount += new
	}
	fmt.Println("output: ", nowCount)
	db.Post(nowCount)
}
