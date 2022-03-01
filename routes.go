package main

import (
	"fmt"
	"net/http"

	"app/data"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	//Middleware Always first

	//Add routes
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)
	a.App.Routes.Get("/create-u", func(w http.ResponseWriter, r *http.Request) {
		u := data.User{
			FirstName: "Andre",
			LastName:  "Leblanc",
			Active:    1,
			Email:     "aleblanc@live.com",
			Password:  "password",
		}
		id, err := a.Models.Users.Insert(u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "%d %s", id, u.FirstName)
	})

	//a.App.Routes.Get("/test-database", func(w http.ResponseWriter, r *http.Request) {
	//	query := "select * from users where id = 1"
	//	row := a.App.DB.Pool.QueryRowContext(r.Context(), query)
	//	var id int
	//	var name string
	//	err := row.Scan(&id, &name)
	//	if err != nil {
	//		a.App.ErrorLog.Println(err)
	//		return
	//	}
	//	fmt.Fprintf(w, "%d %s", id, name)
	//})

	//static routes
	fileServer := http.FileServer(http.Dir("./public"))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return a.App.Routes
}
