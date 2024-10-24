package main

import "net/http"

func (app *application) healthCheckhandler(w http.ResponseWriter, f *http.Request) {
	w.Write([]byte("ok"))
}
