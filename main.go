// package main

// import (
//     "fmt"
//     "net/http"
// 	"github.com/gorilla/mux"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
//     http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
//     http.Handle("/static/", http.StripPrefix("/static/", fs))

//     http.ListenAndServe(":80", nil)
// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/sivasath16/Go-X-Kubernetes/details"
)

func healthHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("checking application health")
	response := map[string]string {
		"STATUS": "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)
}

func routeHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("serving the home page")
	fmt.Fprintf(w, "Application is up and running")
}

func detailsHandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching handle")
	hostname, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	IP, _ := details.GetIp()
	fmt.Println(hostname)
	fmt.Println(IP)
	response := map[string]string {
		"hostname": hostname,
		"ip": IP.String(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
    r := mux.NewRouter()
	r.HandleFunc("/health", healthHandle)
	r.HandleFunc("/", routeHandle)
	r.HandleFunc("/details", detailsHandle)
	log.Println("server has startedg")
	log.Fatal(http.ListenAndServe(":80", r))
}