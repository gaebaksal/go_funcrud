package main

import (
        "fmt"
        "log"
        "net/http"    // ?

        "github.com/gorilla/mux"
        "github.com/jinzhu/gorm"
)

type App struct {
  Router  *mux.Router
  DB      *gorm.DB
}

// 그 함수가 어떤 struct를 위한 메서드인지 표시
func (app *App) Initialize(config *Config) {
  dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
          config.DB.Username,
          config.DB.Password,
          config.DB.Name,
          config.DB.Charset)

  db, err := gorm.Open(config.DB.Dialect, dbURI)
  if err != nil {
    log.Fatal("Could not connect database")
  }

  app.DB = DBMigrate(db)
  app.Router = mux.NewRouter()
  app.setRouters()
}

func (app *App) setRouters() {
    app.Get("/projects", app.GetAllProjects)
    app.Post("/projects", app.CreateProject)
    app.Get("/projects/{title}", app.GetProject)
    app.Put("/projects/{title}", app.UpdateProject)
    app.Delete("/projects/{title}", app.DeleteProject)
    app.Put("/projects/{title}/archive", app.ArchiveProject)
    app.Delete("/projects/{title}/archive", app.RestoreProject)

    // Routing for handling the tasks
    app.Get("/projects/{title}/tasks", app.GetAllTasks)

    app.Post("/projects/{title}/tasks", app.CreateTask)
    app.Get("/projects/{title}/tasks/{id:[0-9]+}", app.GetTask)
    app.Put("/projects/{title}/tasks/{id:[0-9]+}", app.UpdateTask)
    app.Delete("/projects/{title}/tasks/{id:[0-9]+}", app.DeleteTask)
    app.Put("/projects/{title}/tasks/{id:[0-9]+}/complete", app.CompleteTask)
    app.Delete("/projects/{title}/tasks/{id:[0-9]+}/complete", app.UndoTask)
}

func (app *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
    app.Router.HandleFunc(path, f).Methods("GET")
}

func (app *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
    app.Router.HandleFunc(path, f).Methods("POST")
}

func (app *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
    app.Router.HandleFunc(path, f).Methods("PUT")
}

func (app *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
    app.Router.HandleFunc(path, f).Methods("DELETE")
}

func (app *App) Run(host string) {
    log.Fatal(http.ListenAndServe(host, app.Router))
}
