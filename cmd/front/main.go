package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/h311ion/go-grader/cmd/front/admin"
	"github.com/h311ion/go-grader/cmd/front/user"
)

func main() {
	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./web/static")),
	)
	http.Handle("/data/", staticHandler)

	router := mux.NewRouter()

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Use(authenticationMiddleware)

	userRouter.HandleFunc("/task", user.ListTasksAction).
		Methods(http.MethodGet)
	userRouter.HandleFunc("/task/{id:[0-9]+}", user.GetTaskAction).
		Methods(http.MethodGet)
	userRouter.HandleFunc("/profile", user.ProfileAction)

	newUserRouter := router.PathPrefix("/user").Subrouter()
	newUserRouter.HandleFunc("/register", user.RegisterAction)
	newUserRouter.HandleFunc("/auth", user.AuthAction)

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(adminAccessMiddleware)

	adminRouter.HandleFunc("/task", admin.ListTasksAction).
		Methods(http.MethodGet)
	adminRouter.HandleFunc("/task", admin.CreateTaskAction)
	adminRouter.HandleFunc("/task/{id:[0-9]+}", admin.EditTaskAction)
	adminRouter.HandleFunc("/task/{id:[0-9]+}/info", admin.CompletionTaskAction)

	fmt.Println("starting server at :8080")
	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
