package main

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
	"github.com/invopop/ctxi18n"
	"github.com/ituoga/go-start/locales"
)

func main() {
	if err := ctxi18n.LoadWithDefault(locales.Content, "en"); err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	sessionStore := sessions.NewCookieStore([]byte("secret")) // @todo: move to env
	sessionStore.MaxAge(int(24 * time.Hour / time.Second))

}
