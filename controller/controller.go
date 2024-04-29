package controller

import "net/http"

func GetAlgumacoisa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ta pegando \n"))
}
