package main

import(
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	// "math/rand"
	// "strconv"
	"github.com/gorilla/mux"  // must be installed by "go get -u github.com/gorilla/mux"
)
func getweather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q="+params["city"]+"&units=metric&APPID=67fddc841569cd4432f8dcefdae751c3")
	if err != nil {
		json.NewEncoder(w).Encode("Some thing went wrong")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.NewEncoder(w).Encode(string(data))
	}
}
func getlocation(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?lat="+params["lat"]+"&lon="+params["lon"]+"&units=metric&appid=67fddc841569cd4432f8dcefdae751c3")
	if err != nil {
		json.NewEncoder(w).Encode("some thing went wrong")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.NewEncoder(w).Encode(string(data))
	}
}
func main() {
	route := mux.NewRouter()

	route.HandleFunc("/api/weather/{city}", getweather).Methods("GET")
	route.HandleFunc("/api/weather/{lon}/{lat}", getlocation).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", route))
} 