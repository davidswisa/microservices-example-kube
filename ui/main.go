package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
		case "POST":
			fmt.Println("Time to go to bed.")
		case "PUT":
			fmt.Println("Time to wake up!")
		case "DELETE":
			fmt.Println("Time to go to bed.")
		default:
			fmt.Println("error")
		}
		// fmt.Println(req.Method)
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

	// Add handle func for producer.
	http.HandleFunc("/", gui())
	http.HandleFunc("/reservations", reservations())
	http.HandleFunc("/images/restaurant-circular.svg", imagesHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	// Run the web server.
	fmt.Println("starting web server...")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
