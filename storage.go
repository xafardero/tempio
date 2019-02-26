package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func save(temperature string, humidity string) {
	v1, err := readConfig(".env", map[string]interface{}{
		"port":     "3306",
		"host":     "localhost",
		"username": "user",
		"password": "secret",
		"db_name":  "tempio",
	})
	if err != nil {
		panic(fmt.Errorf("Error when reading config: %v\n", err))
	}

	port := v1.GetInt("port")
	host := v1.GetString("host")
	user := v1.GetString("username")
	pass := v1.GetString("password")
	dbName := v1.GetString("db_name")

	fmt.Printf("Reading config for port = %d\n", port)
	fmt.Printf("Reading config for hostname = %s\n", host)

	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, pass, host, dbName)
	db, err := sql.Open("mysql", s)

	if err != nil {
		panic(err.Error())
	}

	createDB(db)
	useDB(db)
	createTable(db)

	stmt2, err := db.Prepare("INSERT INTO thermometer (temperature, humidity) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Table created successfully..")
	}
	_, err = stmt2.Exec(temperature, humidity)

	defer db.Close()
}

func createDB(db *sql.DB) {
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS tempio")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func useDB(db *sql.DB) {
	_, err := db.Exec("USE tempio")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func createTable(db *sql.DB) {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS thermometer(id int NOT NULL AUTO_INCREMENT, temperature varchar(50), humidity varchar(30), create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY (id));")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func readConfig(filename string, defaults map[string]interface{}) (*viper.Viper, error) {
	v := viper.New()
	for key, value := range defaults {
		v.SetDefault(key, value)
	}
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()
	return v, err
}
