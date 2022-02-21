package controllers

import (
	"net/http"
)


func StartMainServer() error {
	http.HandleFunc("/", top)

	return http.ListenAndServe(":8080", nil)

	
}
