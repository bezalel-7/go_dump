package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Status: available")
	fmt.Fprintln(w,"Environment: ",app.config.env)
	fmt.Fprintln(w,"Version:",version)
}