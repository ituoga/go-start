package handlers

import (
	"net/http"

	"github.com/delaneyj/toolbelt/embeddednats"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupPWA(router chi.Router, session sessions.Store, ns *embeddednats.Server) error {
	router.Get("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
		const manifest = `{
		"version": "1.3",
  "short_name": "App Dev",
  "name": "App Dev",
  "share_target": {
    "action": "/phone-share",
    "method": "GET",
    "params": {
      "title": "title",
      "text": "text",
      "url": "url"
    }
  },
  "start_url": "/",
  "background_color": "#ffffff",
  "display": "standalone",
  "theme_color": "#0d085c",
  "icons": [
    {
      "src": "/static/favicon.png",
      "type": "image/png",
      "sizes": "512x512",
      "purpose": "any maskable"
    }
  ]
}`
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(manifest))

	})
	return nil
}
