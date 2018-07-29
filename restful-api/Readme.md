## Creating a RESTful API With Golang

##### Getting Started with A Basic API
To get started we will have to create a very simple server which can handle HTTP requests. To do this we’ll create a new file called main.go. Within this main.go file we’ll want to define 3 distinct functions. A homePage function that will handle all requests to our root URL, a handleRequests function that will match the URL path hit with a defined function and a main function which will kick off our API.

```
package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
```

If we run this on our machine now, we should see our very simple API start up on port 8081 if it’s not already been taken by another process. If we now navigate to http://localhost:8081/ in our local browser we should see Welcome to the HomePage! print out on our screen. This means we have successfully created the base from which we’ll build our REST API.

---
##### Router with Gorilla Mux
Package `gorilla/mux` implements a request router and dispatcher for matching incoming requests to their respective handler.

Before use `gorilla/mux` you will be must install gorilla/mux package by

```
go get -u github.com/gorilla/mux
```

##### Building our Router
```
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Other Functions
...

func handleRequestsV2() {
	Router := mux.NewRouter().StrictSlash(true)

	Router.HandleFunc("/", homePage)
	Router.HandleFunc("/user", getAllUser)
	Router.HandleFunc("/user/{id}", getUserById)
	log.Fatal(http.ListenAndServe(":8081", Router))
}

func main() {
	// Rest API v2.0 - Mux Routers
	handleRequestsV2()
}

```

---