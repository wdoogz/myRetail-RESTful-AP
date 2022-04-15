package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/wdoogz/myRetail-RESTful-API/db_connector"
)

var Store map[string]map[string]map[string]map[string]map[string]string
var UpdateStore map[string]interface{}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to myRetail's RESTful API!!!")
	log.Printf("%s Request: %s --> %s%s", r.Method, r.RemoteAddr, r.Host, r.URL)
}

func products(w http.ResponseWriter, r *http.Request) {
	type apioutput struct {
		ID            int         `json:"id"`
		Name          string      `json:"name"`
		Current_Price interface{} `json:"current_price"`
	}
	log.Printf("%s Request: %s --> %s%s", r.Method, r.RemoteAddr, r.Host, r.URL)
	pidstr := strings.TrimPrefix(r.URL.Path, "/products/")
	pidint, converr := strconv.Atoi(pidstr)
	if converr != nil {
		log.Fatal(converr)
	}

	if r.Method == "GET" {

		redsky, rsexists := os.LookupEnv("REDSKY_BASE_URL_KEY")
		if !rsexists {
			log.Fatal("Environment var REDSKY_BASE_URL_KEY does not exist, cannot pull data!!!")
		}

		redskyurl := redsky + "&tcin=" + pidstr
		gr, err := http.Get(redskyurl)
		if err != nil {
			log.Println(err)
		}

		rqbody, rerr := ioutil.ReadAll(gr.Body)
		if rerr != nil {
			log.Println(rerr)
		}

		jserr := json.Unmarshal([]byte(rqbody), &Store)
		if jserr != nil {
			log.Println(jserr.Error())
		}

		var title string
		for k, v := range Store["data"]["product"]["item"]["product_description"] {
			if k == "title" {
				title = string(v)
			}
		}
		db_vals := db_connector.DBConnect(pidint)
		apioutput_build, _ := json.Marshal(apioutput{ID: pidint, Name: title, Current_Price: db_vals})
		fmt.Fprintf(w, "%s", string(apioutput_build))
	}
	if r.Method == "PUT" {
		body, readerr := ioutil.ReadAll(r.Body)
		if readerr != nil {
			log.Println(readerr)
		}

		jsonerr := json.Unmarshal([]byte(body), &UpdateStore)
		if jsonerr != nil {
			log.Println(jsonerr)
		}

		newValRange := UpdateStore["current_price"].(map[string]interface{})
		var newVal float64
		for k, v := range newValRange {
			if k == "value" {
				newVal = v.(float64)
			}
		}
		db_connector.DBUpdate(pidint, newVal)
		fmt.Fprintf(w, "%s updated ID: %d to %v", r.Method, pidint, newVal)
	}
}

func Handle(port string) {
	http.HandleFunc("/", home)
	http.HandleFunc("/products/", products)
	http.ListenAndServe(":"+port, nil)
}
