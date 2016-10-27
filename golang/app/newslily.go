package app

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // Note the sql package provides the namespace

type story struct {
	Url string `json:url`
	Title string `json:title`
}

func showStories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application-json")

	var Stories []*story

	rows, err := db.Query("select url,title from stories")

	checkErr(err)

	for rows.Next() {
		var url string
		var title string
		err = rows.Scan(&url,&title)
		this := &story{Url:url,Title:title}
		Stories = append(Stories,this)
	}
	storyJson, err := json.Marshal(Stories)
	checkErr(err)
	fmt.Fprintf(w, string(storyJson))

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the default response")
}

func checkErr(err error) {
	fmt.Println("some fucking error, lol")
	if err != nil {
		panic(err)
	}
}

func App() {
	var err error
	//not my real db creds you tricksy hobbitses
	db, err = sql.Open("mysql", "notroot:notthispass@/golang")
	if err != nil {
		fmt.Println("Some fucking error, lol")
	}

	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/stories/", showStories)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
