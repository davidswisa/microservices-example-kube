package main

import (
	"bytes"
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
		content := "{}"
		params := mux.Vars(req)
		id := params["id"]
		fmt.Println(id)

		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		switch req.Method {
		case "GET":

			resp, err := http.Get("http://querier:8081/reservations")
			if err != nil {
				fmt.Println(err.Error())
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err.Error())
			}
			content = string(body)
			fmt.Println(string(body))

		case "POST":
			resp, err := http.Post("http://prod:8080/reservations", "Application/json", bytes.NewReader(payload))
			if err != nil {
				fmt.Println(err.Error())
			}
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
			}

		case "PUT":

			req, err := http.NewRequest(http.MethodPut, "http://prod:8080/reservations/"+id, bytes.NewReader(payload))
			if err != nil {
				panic(err)
			}

			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
			}

		case "DELETE":
			fmt.Println("Time to go to bed.")
			req, err := http.NewRequest(http.MethodDelete, "http://prod:8080/reservations/"+id, bytes.NewReader(payload))
			if err != nil {
				panic(err)
			}

			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			if resp.StatusCode != 200 {
				fmt.Println(resp.StatusCode)
			}

		default:
			fmt.Println("error")
		}

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
