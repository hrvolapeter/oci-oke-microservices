package main

import (
	"fmt"
	"net/http"
	"os"
)

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(ct.counter)
	ct.counter++
	fmt.Fprintln(w, "Service B - Counter:", ct.counter)
}

func main() {
	port := "8080"
	if p := os.Getenv("MONGO_PASS"); p != "" {
		port = p
	}
	th := &CounterHandler{counter: 0}
	http.Handle("/", th)
	http.ListenAndServe(":"+port, nil)
}
