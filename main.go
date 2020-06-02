package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func init() {
	if err := os.MkdirAll("id", 0700); err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{path}/{file}", getFile)
	r.HandleFunc("/{path}/{file}/{value}", putFile).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8081", r))
}

// ================================================================================================

func getFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "")
	w.Header().Set("Date", "")

	path := strings.TrimSpace(mux.Vars(r)["path"])
	file := strings.TrimSpace(mux.Vars(r)["file"])

	log.Println("getFile", path, file)

	if value, err := Get(path, file); err != nil { // todo analyze err
		_, _ = w.Write(nil)
	} else {
		_, _ = w.Write(value)
	}
}

func putFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "")
	w.Header().Set("Date", "")

	path := strings.TrimSpace(mux.Vars(r)["path"])
	file := strings.TrimSpace(mux.Vars(r)["file"])
	value := strings.TrimSpace(mux.Vars(r)["value"])

	log.Println("putFile", path, file, value)

	if err := Set(path, file, value); err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
	}
}

// ================================================================================================

func Get(path, file string) ([]byte, error) {
	if value, err := ioutil.ReadFile(toAbsolute(path, file)); err != nil {
		return nil, err
	} else {
		return value, nil
	}
}

func Set(path, file, value string) error {
	if err := ioutil.WriteFile(toAbsolute(path, file), []byte(value), 0600); err != nil {
		if err := os.MkdirAll(toAbsolute(path, ""), 0700); err != nil {

			log.Println(err)
			return err
		}
		return ioutil.WriteFile(toAbsolute(path, file), []byte(value), 0600)
	}
	return nil
}

// ================================================================================================

func toAbsolute(path, file string) string {
	return string([]byte{
		// '/',
		'i', 'd', '/',
		path[0], '/',
		path[1], '/',
		path[2], '/',
		path[3], '/',
		path[4], '/',
		path[5], '/',
		path[6], '/',
		path[7], '/',
		path[8], '/',
		path[9], '/',
		path[10], '/',
		path[11], '/',
	}) + file
}
