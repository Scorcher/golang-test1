package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")

	kv, err := NewKVStorageWrapper(KVStorageWrapperTypeInMemory)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	app := &application{kv, mux}
	mux.HandleFunc("/storage/get/", app.HTTPServerStorageGet)
	mux.HandleFunc("/storage/put/", app.HTTPServerStoragePut)
	mux.HandleFunc("/storage/delete/", app.HTTPServerStorageDelete)
	http.Handle("/", mux)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
