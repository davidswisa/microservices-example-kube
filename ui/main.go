package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func gui() func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			log.Fatal(err)
		}
		wrt.Write([]byte(content))
	})
}

func reservations() func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case "GET":
			fmt.Println("Time to wake up!")
			// xhr.open("GET", "http://localhost:8081/reservations", true);

		case "POST":
			fmt.Println("Time to go to bed.")
			// xhr.open("POST", "http://localhost:8080/reservations", true);

		case "PUT":
			fmt.Println("Time to wake up!")
			// xhr.open("PUT", "http://localhost:8080/reservations/"+ Id, true);

		case "DELETE":
			fmt.Println("Time to go to bed.")
			// var url = "http://localhost:8080/reservations/"+ Id;

		default:
			fmt.Println("error")
		}

		params := mux.Vars(req)
		id := params["id"]
		fmt.Println(id)

		content := "{}"
		wrt.Write([]byte(content))
	})
}

func imagesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/restaurant-circular.svg")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./images/favicon.ico")
}

func main() {
	r := mux.NewRouter()

	// Add handle func for producer.
	r.HandleFunc("/", gui())
	r.HandleFunc("/reservations", reservations())
	r.HandleFunc("/reservations/{id}", reservations())
	r.HandleFunc("/images/restaurant-circular.svg", imagesHandler)
	r.HandleFunc("/favicon.ico", faviconHandler)

	http.Handle("/", r)
	// Run the web server.
	fmt.Println("starting web server...")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
