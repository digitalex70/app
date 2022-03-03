package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {
	//Middleware Always first

	//Add routes
	a.App.Routes.Get("/", a.Handlers.Home)
	a.App.Routes.Get("/go-page", a.Handlers.GoPage)
	a.App.Routes.Get("/jet-page", a.Handlers.JetPage)
	a.App.Routes.Get("/sessions", a.Handlers.SessionTest)
	a.App.Routes.Get("/users/login", a.Handlers.UserLogin)
	a.App.Routes.Post("/users/login", a.Handlers.PostUserLogin)
	a.App.Routes.Get("/users/logout", a.Handlers.Logout)

	a.App.Routes.Get("/get-all-u", func(w http.ResponseWriter, r *http.Request) {
		users, err := a.Models.Users.GetAll()
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		for _, x := range users {
			fmt.Fprintf(w, x.LastName)
		}
	})
	a.App.Routes.Get("/u-get/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		x, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "%s %s %d", x.LastName, x.Token.Expires, x.Id)

		//time.Sleep(10 * time.Second)

	})

	a.App.Routes.Get("/u-user/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		u, err := a.Models.Users.Get(id)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		u.LastName = a.App.RandomString(12)
		err = u.Update(*u)
		if err != nil {
			a.App.ErrorLog.Println(err)
			return
		}
		fmt.Fprintf(w, "Update LastName to %s", u.LastName)
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
