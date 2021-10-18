package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "HELLO", "layout", "public_navbar", "top")
}
