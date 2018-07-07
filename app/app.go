package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/achanda/testrest/app/handler"
	"github.com/achanda/testrest/app/model"
	"github.com/achanda/testrest/config"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// App represents the overall app
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Username,
		config.DB.Password,
		config.DB.Name)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	db.LogMode(true)

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/payments", a.GetAllPayments)
	a.Get("/payments/{id}", a.GetPayment)
	a.Post("/payments", a.CreatePayment)
	a.Put("/payments/{id}", a.UpdatePayment)
	a.Delete("/payments/{id}", a.DeletePayment)

	a.Get("/version", a.GetVersion)
}

func (a *App) GetVersion(w http.ResponseWriter, r *http.Request) {
	handler.GetVersion(a.DB, w, r)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

func (a *App) GetAllPayments(w http.ResponseWriter, r *http.Request) {
	handler.GetAllPayments(a.DB, w, r)
}

func (a *App) CreatePayment(w http.ResponseWriter, r *http.Request) {
	handler.CreatePayment(a.DB, w, r)
}

func (a *App) GetPayment(w http.ResponseWriter, r *http.Request) {
	handler.GetPayment(a.DB, w, r)
}

func (a *App) DeletePayment(w http.ResponseWriter, r *http.Request) {
	handler.DeletePayment(a.DB, w, r)
}

func (a *App) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	handler.UpdatePayment(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
