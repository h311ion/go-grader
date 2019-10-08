package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/h311ion/go-grader/cmd/front/common"
	"github.com/h311ion/go-grader/internal/user"
)

type route string

var templates map[route]*template.Template

const profileRoute route = "profile"

func init() {
	templates = map[route]*template.Template{
		profileRoute: template.Must(template.ParseFiles("./web/template/user/index.html", "./web/template/user/profile.html")),
	}
}

func ListTasksAction(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("List user's tasks"))
	if err != nil {
		log.Printf("failed to write response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetTaskAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := strconv.Atoi(idVar)
	if err != nil || id < 1 {
		common.SendBadRequestResponse(w, "task id must be int and > 0")
		return
	}

	fmt.Printf("%+v", reflect.TypeOf(id))

	_, err = w.Write([]byte("User's task details"))
	if err != nil {
		log.Printf("failed to write response: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func AuthAction(writer http.ResponseWriter, r *http.Request) {

}

func RegisterAction(writer http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		err := user.RegisterUser(email, password)
		if err != nil {

		}
	}
}

func ProfileAction(w http.ResponseWriter, r *http.Request) {
	profile := user.GetProfile(user.NewUser("test@example.com"))

	err := templates[profileRoute].ExecuteTemplate(w, "base", profile)
	if err != nil {
		println(err.Error())
	}
}
