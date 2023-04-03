package ConnectDb

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Config configuration

type configuration struct {
	Database struct {
		Dbname string `json:"dbname"`
		Dbuser string `json:"dbuser"`
		Dbpass string `json:"dbpass"`
		Target string `json:"target"`
	} 
}

func LoadConfiguration(filename string) error {

	configFile, err := os.Open("config.json")
	if err != nil {
		log.Println(err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		log.Println(err)
	}
	return err
}

// connect with database
func ConnectDB(dbname, dbuser, dbpass, target string) error {
	// db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+target+")/"+dbname)
	db, err := sql.Open("mysql", "student:admin01@tcp(192.168.180.17:3306)/data")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Config)
	fmt.Println("Succesfully connected!")
	db.Ping()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	return err
}
