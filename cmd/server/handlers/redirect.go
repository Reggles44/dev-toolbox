package handlers

import (
	"net/http"
)


var RedirectHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/success/", http.StatusMovedPermanently)
})

