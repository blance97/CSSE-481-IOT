package main

import (
	"log"
	"net/http"
)

//var mu sync.Mutex

var db = initDB()

func main() {
	log.SetFlags(log.Lshortfile)
	createSchema()
	//storeCSVData()
	CopyPlaces("C:/Go/test.csv")

	http.Handle("/", http.FileServer(http.Dir("webroot")))
	//	http.HandleFunc("/upload", signIn)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

/**
Returns User so that it can validate whether or not a message belongs to them.
*/
// func getData(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Path in Users Handler: ", r.URL.Path)
// 	case "/getUser":
// 		q := User{
// 			Username: username,
// 		}
// 		json.NewEncoder(w).Encode(q)
// 	case "/getUserInfo/":
// 		data := r.URL.Query()
// 		Username := data.Get("Username")
// 		json.NewEncoder(w).Encode(getUserInfo(Username))
// 	}
// }
