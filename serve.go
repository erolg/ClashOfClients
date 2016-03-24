package clashofclients

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/redisstore"
	"net/http"
	"strconv"
)

func (c *ClashOfClients) Serve() *http.ServeMux {

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

		if nickname == "" || email == "" {
			fmt.Fprintf(w, "Lütfen mail ve nickname gönder")
		} else if c.CheckNickName(nickname) {
			fmt.Fprintf(w, "Bu nickname kullanılıyor yeni bir tane seç!")
		} else {
			session, err := sessionStore.Get(req, c.Cfg.Name)
			if err != nil {
				panic(err)
			} else {
				session.Values["nickname"] = nickname
				session.Values["email"] = email
				c.game.session, _ = c.CreateGameStore(nickname, email)
				session.Values["lastestGame"] = c.game.session
				session.Options.MaxAge = c.Cfg.SessionMaxAge
				if err := session.Save(req, w); err != nil {
					panic(err)
				}
			}

		}
	})
	n.UseHandler(mux)
	n.Run(":3000")

	return mux
}
