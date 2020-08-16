package app

import (
	"net/http"
	"search/src/app/handler"
	"search/src/app/middleware"
	"search/src/config"
	"search/src/logger"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Config *config.Config
}

var log = logger.GetLogger("app")

func NewApp(c *config.Config) *App {
	log.Printf("initialzing application with configs [%+v], [%+v]\n", c.AppConfig, c.DeezerConfig)

	r := mux.NewRouter()
	m := middleware.NewMiddleware(c)
	h := handler.NewHandler(c)

	r.HandleFunc("/", h.Base)
	r.HandleFunc("/ecv", h.Ecv)
	r.HandleFunc("/running", h.Running)
	r.HandleFunc("/uptime", m.AuthMiddleware(h.Uptime))
	r.HandleFunc("/api/v1/cache/clear", m.AuthMiddleware(h.ClearCache)).Methods("POST")
	r.HandleFunc("/api/v1/cache/{key}", m.AuthMiddleware(h.GetCacheKey)).Methods("GET")
	r.HandleFunc("/api/v1/cache/{key}/clear", m.AuthMiddleware(h.ClearCacheKey)).Methods("POST")
	r.HandleFunc("/api/v1/search/{query}", m.AuthMiddleware(h.Search)).Methods("GET")
	r.HandleFunc("/api/v2/search/{query}", m.AuthMiddleware(h.SearchV2)).Methods("GET")

	return &App{
		Router: r,
		Config: c,
	}
}

func (a *App) Run() {
	log.Printf("starting application on port %s", a.Config.AppConfig.Port)
	log.Fatal(http.ListenAndServe(":"+a.Config.AppConfig.Port, a.Router))
}
