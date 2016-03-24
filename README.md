# ClashOfClients
Clash Of Clients
 Sample
 ```go
package main

import (
	"github.com/erolg/ClashOfClients"
	"fmt"
	"net/http"
)

func main() {

	c := clashofclients.New()
	mux := c.Serve()

	mux.HandleFunc("/play", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "It works!")
	})
}
```
