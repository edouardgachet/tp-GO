package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func SimpleHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		start := time.Now()

		var hour int = start.Hour()
		var minute int = start.Minute()

		fmt.Fprintf(w, "il est : %dh%d", hour, minute)
	}
}

func AddHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}
		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
		writeData(req.PostForm["author"][0], req.PostForm["entry"][0])
	}
}

func EntriesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		data, err := ioutil.ReadFile("test.txt") // lire le fichier
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(data))
	}
}

func writeData(author string, entry string) {

	if author != "" && entry != "" {

		file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		defer file.Close() // on ferme automatiquement à la fin de notre programme

		if err != nil {
			panic(err)
		}

		_, err = file.WriteString(author + ":" + entry + "\n") // écrire dans le fichier
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	http.HandleFunc("/", SimpleHandler)
	http.HandleFunc("/add", AddHandler)
	http.HandleFunc("/entries", EntriesHandler)
	//writeData()
	http.ListenAndServe(":4567", nil)
}
