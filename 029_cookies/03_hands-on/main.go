package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("visit-counter")
	if err == nil {
		counter, _ := strconv.Atoi(c.Value)
		counter = counter + 1
		http.SetCookie(w, &http.Cookie{
			Name:  "visit-counter",
			Value: strconv.Itoa(counter),
			Path:  "/",
		})
	} else if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "visit-counter",
			Value: "0",
			Path:  "/",
		})
	} else {
		fmt.Println(err)
	}
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("visit-counter")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "YOUR COOKIE:", c)
}
