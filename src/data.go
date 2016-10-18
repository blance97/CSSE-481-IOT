package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "blance97"
	DB_NAME     = "IOT"
)

func initDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	//defer db.Close()
	log.Println("Successfull opened db")
	return db
}
func createSchema() {
	sql_table := `
  CREATE TABLE userinfo
(
		index SERIAL PRIMARY KEY,
    Temp  INT NOT NULL,
		Absolute_Pressure FLOAT NOT NULL,
		Relative_Pressure FLOAT NOT NULL,
		Humidity FLOAT NOT NULL,
		Rain_Sensor_1 FLOAT NOT NULL,
		Rain_Sensor_2 FLOAT NOT NULL,
		Light_Sensor  FLOAT NOT NULL
)
WITH (OIDS=FALSE);
`
	_, err2 := db.Exec(sql_table)
	if err2 != nil {
		log.Print(err2)
	}
}
func storeCSVData() {
	fmt.Println("# Updating")
	sql_stmt := `COPY userinfo(Temp,Absolute_Pressure,Relative_Pressure,Humidity,Rain_Sensor_1,Rain_Sensor_2,Light_Sensor)
	FROM 'C:\Go\test.csv' DELIMITER ',' CSV HEADER;`
	_, err := db.Exec(sql_stmt)
	checkErr(err)
}

func CopyPlaces(filename string) {
	csvfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {
		fmt.Printf("Absolute_Pressure : %s,  Absolute_Pressure : %s, Relative_Pressure : %s, Humidity : %s, Rain_Sensor_1 : %s, Rain_Sensor_2 : %s, Light_Sensor : %s\n", each[0], each[1], each[2], each[3], each[4], each[5], each[6])

		//	err = db.QueryRow("INSERT INTO userinfo(Absolute_Pressure,Absolute_Pressure,Relative_Pressure,Humidity,Rain_Sensor_1,Rain_Sensor_2,Light_Sensor) VALUES($1,
		//checkErr(err)
	}

}

// 	fmt.Println("# Inserting values")
//
// 	var lastInsertId int
// 	err = db.QueryRow("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) returning uid;", "astaxie", "研发部门", "2012-12-09").Scan(&lastInsertId)
// 	checkErr(err)
// 	fmt.Println("last inserted id =", lastInsertId)
//
// 	fmt.Println("# Updating")
// 	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
// 	checkErr(err)
//
// 	res, err := stmt.Exec("astaxieupdate", lastInsertId)
// 	checkErr(err)
//
// 	affect, err := res.RowsAffected()
// 	checkErr(err)
//
// 	fmt.Println(affect, "rows changed")
//
// 	fmt.Println("# Querying")
// 	rows, err := db.Query("SELECT * FROM userinfo")
// 	checkErr(err)
//
// 	for rows.Next() {
// 		var uid int
// 		var username string
// 		var department string
// 		var created time.Time
// 		err = rows.Scan(&uid, &username, &department, &created)
// 		checkErr(err)
// 		fmt.Println("uid | username | department | created ")
// 		fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
// 	}
//
// 	fmt.Println("# Deleting")
// 	stmt, err = db.Prepare("delete from userinfo where uid=$1")
// 	checkErr(err)
//
// 	res, err = stmt.Exec(lastInsertId)
// 	checkErr(err)
//
// 	affect, err = res.RowsAffected()
// 	checkErr(err)
//
// 	fmt.Println(affect, "rows changed")
// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
