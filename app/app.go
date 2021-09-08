package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-task-intech/config"
	"test-task-intech/handler"
	"test-task-intech/storage"
)

type App struct {
	Router *mux.Router
	DB     storage.DB
}

func (a *App) Initialize(config *config.Config) {
	a.DB = storage.InitService(config)
	a.Router = mux.NewRouter()
	a.SetRouters()
}

// Запуск сервера.
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

// Настройка роутеров.
func (a *App) SetRouters() {
	a.Get("/GetBookByAuthor/{author}", a.GetBookByAuthor)
}

// Обёртка GET метода для роутеров.
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Обработчик запросов.
func (a *App) GetBookByAuthor(w http.ResponseWriter, r *http.Request) {
	handler.GetByAuthorHandler(a.DB, w, r)
}
