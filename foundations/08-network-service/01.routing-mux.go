package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	hello := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Hello!!!")
	}

	goodbye := func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)
		fmt.Fprint(resp, "Goodbye!!!")
	}

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/goodbye", goodbye)

	fmt.Println("Service run on :4040 >>>")
	http.ListenAndServe(":4040", mux)
}
