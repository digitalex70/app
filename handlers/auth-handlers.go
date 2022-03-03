package handlers

import "net/http"

func (h *Handlers) UserLogin(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "login", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println(err)
	}
}

func (h *Handlers) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	email := r.Form.Get("email")
	pwd := r.Form.Get("password")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	user, err := h.Models.Users.GetByEmail(email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	matches, err := user.PasswordMatches(pwd)
	if err != nil {
		w.Write([]byte("Error validating Password !!!"))
		return
	}

	if !matches {
		w.Write([]byte("User not found"))
		return
	}
	h.App.Session.Put(r.Context(), "userId", user.Id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	h.App.Session.RenewToken(r.Context())
	h.App.Session.Remove(r.Context(), "userId")
	http.Redirect(w, r, "/users/login", http.StatusSeeOther)
}
