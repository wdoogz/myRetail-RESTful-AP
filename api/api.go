package api

import (
	"fmt"
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strings"
	"os"

	"github.com/wdoogz/myRetail-RESTful-API/db_connector"
)

var Store map[string]map[string]map[string]map[string]map[string]string

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to myRetail's RESTful API!!!")
	log.Printf("%s Request: %s --> %s%s", r.Method, r.RemoteAddr, r.Host, r.URL)
}

func products(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s Request: %s --> %s%s", r.Method, r.RemoteAddr, r.Host, r.URL)
	pid := strings.TrimPrefix(r.URL.Path, "/products/")

	redsky, rsexists := os.LookupEnv("REDSKY_BASE_URL_KEY")
	if !rsexists {
		log.Fatal("Environment var REDSKY_BASE_URL_KEY does not exist, cannot pull data!!!")
	}

	redskyurl := redsky+"&tcin="+pid
	gr, err := http.Get(redskyurl)
	if err != nil {
		log.Println("here")
		log.Println(err)
	}

	rqbody, rerr := ioutil.ReadAll(gr.Body)
	if rerr != nil {
		log.Println("here")
		log.Println(rerr)
	}

	jserr := json.Unmarshal([]byte(rqbody), &Store)
	if jserr != nil {
		log.Println(jserr)
	}

	var title string
	for k,v := range Store["data"]["product"]["item"]["product_description"] {
		if k == "title" {
			title = string(v)
		}
	}
	fmt.Println("here")
	db_connector.DBConnect(pid)
	fmt.Fprintf(w, "Title: %s", title)

}

func Handle(port string) {
	http.HandleFunc("/", home)
	http.HandleFunc("/products/", products)
	http.ListenAndServe(":"+port, nil)
}
