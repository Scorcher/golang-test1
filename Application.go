package main

import (
	"fmt"
	"net/http"
	"time"
)

type application struct {
	kv  *KVStorageWrapper
	mux *http.ServeMux
}

func (app *application) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.mux.ServeHTTP(w, r)
}

func (app *application) HandlerGet(w http.ResponseWriter, r *http.Request) {
	val, _ := app.kv.Get(r.Context(), "key1")
	fmt.Fprint(w, val)
}

func (app *application) HTTPServerStorageGet(w http.ResponseWriter, r *http.Request) {
	key, err := getKey(r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "Cant parse request path\n")
		return
	}

	start := time.Now()
	result, err := app.kv.Get(r.Context(), key)
	elapsed := int(time.Since(start) / time.Millisecond)
	if err != nil {
		fmt.Fprintf(w, "Get key \"%s\" | Took: %d ms | ERROR: %s\n", key, elapsed, err)
	} else {
		fmt.Fprintf(w, "Get key \"%s\" | Took: %d ms | Value: \"%+v\"\n", key, elapsed, result)
	}
}

func (app *application) HTTPServerStoragePut(w http.ResponseWriter, r *http.Request) {
	key, value, err := getKeyValue(r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "Cant parse request path\n")
		return
	}

	start := time.Now()
	err = app.kv.Put(r.Context(), key, value)
	elapsed := int(time.Since(start) / time.Millisecond)
	if err != nil {
		fmt.Fprintf(w, "Put key \"%s\", value \"%s\" | Took: %d ms | ERROR: %s\n", key, value, elapsed, err)
	} else {
		fmt.Fprintf(w, "Put key \"%s\", value \"%s\" | Took: %d ms | OK\n", key, value, elapsed)
	}
}

func (app *application) HTTPServerStorageDelete(w http.ResponseWriter, r *http.Request) {
	key, err := getKey(r.URL.Path)
	if err != nil {
		fmt.Fprintf(w, "Cant parse request path\n")
		return
	}

	start := time.Now()
	err = app.kv.Delete(r.Context(), key)
	elapsed := int(time.Since(start) / time.Millisecond)
	if err != nil {
		fmt.Fprintf(w, "Delete key \"%s\" | Took: %d ms | ERROR: %s\n", key, elapsed, err)
	} else {
		fmt.Fprintf(w, "Delete key \"%s\" | Took: %d ms | OK\n", key, elapsed)
	}
}
