package main

import (
	"html/template"
	"net/http"
)

var pleaseLoginPage *template.Template

func init() {
	pleaseLoginPage = template.Must(template.ParseFiles("./web/template/user/index.html", "./web/template/user/no_login.html"))
}

func authenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := r.Cookie("session_id")
		if err != nil {
			_ = pleaseLoginPage.ExecuteTemplate(w, "base", nil)
			return
		}

		println(session)

		next.ServeHTTP(w, r)
	})
}

func adminAccessMiddleware(next http.Handler) http.Handler {
	return next
}
