package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if (*r).Method == "GET" {
		vars := mux.Vars(r)

		sID, err := strconv.Atoi(vars["secretID"])
		if err != nil {
			log.Printf("Couldn ot parse secretID: %s", err)
		}

		s := GetSecret(sID)

		secret, err := json.Marshal(s)
		if err != nil {
			log.Printf("JSON marshalling secret failed: %s", err)
		}
		w.Write(secret)
	}

	if (*r).Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Couldn't parse body: %s", err)
		}
		secret := Secret{}

		err = json.Unmarshal(b, &secret)
		if err != nil {
			log.Printf("Could not unmarshal secret body: %s", err)
		}

		err = CreateSecret(&secret)
		if err != nil {
			log.Print(err)
		}

		w.WriteHeader(200)
	}
}
