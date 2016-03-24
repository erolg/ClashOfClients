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

	mux.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page %s!", c.Cfg.Name)
	})
}
```
