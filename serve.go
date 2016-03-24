package clashofclients

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/redisstore"
	"net/http"
	"strconv"
)

func (c *ClashOfClients) Serve() {

	n := negroni.Classic()

	c.CreateRedisPool()
	sessionStore, _ := redisstore.New(10, "tcp", ":"+strconv.Itoa(c.Cfg.RedisPort), c.Cfg.RedisPassword, []byte(c.Cfg.SessionSecret))
	n.Use(sessions.Sessions(c.Cfg.Name, sessionStore))

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page %s!", c.Cfg.Name)
	})
	mux.HandleFunc("/register", func(w http.ResponseWriter, req *http.Request) {
		nickname := req.URL.Query().Get("nickname")
		email := req.URL.Query().Get("email")

		session, err := sessionStore.Get(req, c.Cfg.Name)
		if err != nil {
			panic(err)
		}
		if c.CheckNickName(nickname) {
			fmt.Println("Bu nickname kullanılıyor yeni bir tane seç!")
		} else {
			session.Values["nickname"] = nickname
			session.Values["email"] = email
			var status string
			session.Values["lastestGame"], status = c.CreateGameStore(nickname, email)
			fmt.Println(session.Values["lastestGame"])
			fmt.Println(status)
			session.Options.MaxAge = c.Cfg.SessionMaxAge
			if err := session.Save(req, w); err != nil {
				panic(err)
			}
		}

	})
	n.UseHandler(mux)
	n.Run(":3000")
}
