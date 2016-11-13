package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//var mu sync.Mutex

var db = InitDB("data/data.db")

type Roomier struct {
	Rooms   []string
	Private []bool
}

func main() {
	log.SetFlags(log.Lshortfile)
	createSchema()

	http.Handle("/", http.FileServer(http.Dir("webroot")))

	http.HandleFunc("/getData/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got it")
		f := r.URL.Query()
		data := f.Get("data")
		q := getData(data)
		json.NewEncoder(w).Encode(q)
	})
	http.HandleFunc("/listFile", func(w http.ResponseWriter, r *http.Request) {
		files, _ := ioutil.ReadDir("tmp/")
		var Room []string
		var Priv []bool
		for _, file := range files {
			Room = append(Room, file.Name())
		}
		q := Roomier{Rooms: Room, Private: Priv}
		json.NewEncoder(w).Encode(q)
		return
	})
	http.HandleFunc("/getFile/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got it")
		f := r.URL.Query()
		data := f.Get("data")
		clearData()
		CopyPlaces("C:/Projects/Go/src/github.com/Blance97/MA-495-IOT/src/tmp/" + data)
	})

	http.HandleFunc("/uploadData", func(w http.ResponseWriter, r *http.Request) {
		log.Println("going")
		// the FormFile function takes in the POST input id file

		r.ParseMultipartForm(5 * 1024 * 1024)

		file, header, err := r.FormFile("file")

		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		defer file.Close()
		if _, err := os.Stat("tmp/" + header.Filename); err == nil {
			log.Println("File already exists")
			http.Redirect(w, r, "index.html", 302)
			return
		}
		clearData()
		out, err := os.OpenFile("tmp/"+header.Filename, os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		// write the content from POST to the file
		_, err = io.Copy(out, file)
		if err != nil {
			fmt.Fprintln(w, err)
		}

		fmt.Fprintf(w, "File uploaded successfully : ")
		fmt.Fprintf(w, header.Filename)
		CopyPlaces("C:/Projects/Go/src/github.com/Blance97/MA-495-IOT/src/tmp/" + header.Filename)
		//CopyPlaces("src/tmp/" + header.Filename)
		http.Redirect(w, r, "index.html", 302)
	})
	log.Fatal(http.ListenAndServe(":3333", nil))
}
