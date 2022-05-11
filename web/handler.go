package web

import (
	"forumDemo"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"text/template"
)

// NewHandler
// j'ai plus bas défini ma structure, je crée maintenant le constructeur
func NewHandler(store forumDemo.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)
	h.Route("/threads", func(r chi.Router) {
		r.Get("/", h.ThreadsList())
	})

	return h
}

type Handler struct {
	*chi.Mux
	store forumDemo.Store // Notez que je demande un truc qui implémente cette interface ici
}

// c'est data.Threads, je peux y accéder avec .Threads direct
const threadsListHTML = `
<h1>Threads</h1>
<dl>
{{range .Threads}}
	<dt>{{.Title}}</dt>
	<dd>{{.Description}}</dd>
{{end}}
</dl>
`

func (h *Handler) ThreadsList() http.HandlerFunc {
	// Je peux ici déclarer des variables qui seront appelées une seule fois
	// plutôt qu'à chaque requête, c'est efficient en termes de ressources.
	type data struct {
		Threads []forumDemo.Thread
	}

	// Je laisse le nom du template vide, je ne vais avoir qu'un seul template dans cette instance

	// template.Must va faire planter le programme si je n'arrive pas à gérer le template, ce qui est
	// logique parce-que sinon, rien n'aurait marché
	tmpl := template.Must(template.New("").Parse(threadsListHTML))

	return func(writer http.ResponseWriter, request *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(writer, data{Threads: tt})
	}
}
