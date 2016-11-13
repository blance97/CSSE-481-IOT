package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	Temp              string
	Absolute_Pressure string
	Relative_Pressure string
	Humidity          string
	Rain_Sensor_1     string
	Rain_Sensor_2     string
	Light_Sensor      string
}

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	log.Println(filepath)
	if err != nil {
		log.Print(err)
	}
	if db == nil {
		log.Print("db nil")
	}
	log.Println("Successfull opened db")
	return db
}

func createSchema() {

	sql_table := `
	CREATE TABLE IF NOT EXISTS Data(
    Temp  TEXT NOT NULL,
		Absolute_Pressure TEXT NOT NULL,
		Relative_Pressure TEXT NOT NULL,
		Humidity TEXT NOT NULL,
		Rain_Sensor_1 TEXT NOT NULL,
		Rain_Sensor_2 TEXT NOT NULL,
		Light_Sensor  TEXT NOT NULL
);
`
	_, err := db.Exec(sql_table)
	if err != nil {
		log.Print(err)
	}
}
func storeCSVData() {
	fmt.Println("# Updating")
	sql_stmt := `COPY userinfo(Temp,Absolute_Pressure,Relative_Pressure,Humidity,Rain_Sensor_1,Rain_Sensor_2,Light_Sensor)
	FROM 'C:\Go\test.csv' DELIMITER ',' CSV HEADER;`
	_, err := db.Exec(sql_stmt)
	checkErr(err)
}
func clearData() {
	fmt.Println("Clearing Data")
	_, err := db.Exec("delete from Data")
	checkErr(err)
}

func CopyPlaces(filename string) {
	sql_stmt := `
	INSERT OR REPLACE INTO Data(
    Temp,
    Absolute_Pressure,
    Relative_Pressure,
		Humidity,
		Rain_Sensor_1,
		Rain_Sensor_2,
		Light_Sensor
	)values(?, ?, ?, ?, ?,?,?)
	`
	stmt, err := db.Prepare(sql_stmt)
	if err != nil {
		log.Print(err)
	}
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
		c := Data{
			Temp:              each[0],
			Absolute_Pressure: each[1],
			Relative_Pressure: each[2],
			Humidity:          each[3],
			Rain_Sensor_1:     each[4],
			Rain_Sensor_2:     each[5],
			Light_Sensor:      each[6],
		}
		if _, err := stmt.Exec(c.Temp, c.Absolute_Pressure, c.Relative_Pressure, c.Humidity, c.Rain_Sensor_1, c.Rain_Sensor_2, c.Light_Sensor); err != nil {
			log.Println(err)
		}
	}

}
func getData(data string) []Data {
	rows, err := db.Query("SELECT * FROM Data")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var result []Data
	for rows.Next() {
		item := Data{}
		err2 := rows.Scan(&item.Temp, &item.Absolute_Pressure, &item.Relative_Pressure, &item.Humidity, &item.Rain_Sensor_1, &item.Rain_Sensor_2, &item.Light_Sensor)
		if err2 != nil {
			log.Println("Error scanning Time in ApptTypeTime")
		}
		result = append(result, item)
	}
	return result
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func getJSON(r *http.Request) map[string]interface{} {
	var data map[string]interface{}

	//	log.Printf("getJSON:\tBegin execution")
	if r.Body == nil {
		log.Printf("No Request Body")
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("Error Decoding JSON")
	}
	defer r.Body.Close()
	return data
} //decode JSON
