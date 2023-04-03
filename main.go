package main

import (
	"fmt"
	"log"
	"os"
	db "password/database"
	gn "password/passwordgen"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	LogError()
	db.LoadConfiguration("config.json")
	db.ConnectDB(db.Config.Database.Dbname, db.Config.Database.Dbuser, db.Config.Database.Dbpass, db.Config.Database.Target)
	username := getUser()
	fmt.Println("Username:", username)
	password := gn.GenPassword(gn.LowerCount, gn.UpperCount, gn.SpecialCount)
	fmt.Println("Password:", password)

}


func LogError() {
	file, err := os.OpenFile("error.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(fmt.Errorf("could not open logfile: %s", err))
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// Give username for log in.
func getUser() string {
	var username string
	fmt.Println("Please give a username")
	fmt.Scanln(&username)
	return username
}
