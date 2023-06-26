package routes

import (
	"fmt"
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/sessions"
	"github.com/aleale2121/go-eccomerce-app/utils"
)

func registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "register.html", struct {
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func registerPostHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	_, err := models.NewUser(user)
	checkErrRegister(err, w, r)
}

func checkErrRegister(err error, w http.ResponseWriter, r *http.Request) {
	message := "Cadastrado com sucesso!"
	if err != nil {
		switch err {
		case models.ErrMaxLimit,
			models.ErrRequiredFirstName,
			models.ErrRequiredLastName,
			models.ErrRequiredEmail,
			models.ErrInvalidEmail,
			models.ErrEmailTaken,
			models.ErrRequiredPassword:
			message = fmt.Sprintf("%s", err)
		default:
			utils.InternalServerError(w)
			return
		}
		sessions.Message(message, "danger", r, w)
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}
	sessions.Message(message, "success", r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.InternalServerError(w)
		return
	}
	total := int64(len(users))
	utils.ExecuteTemplate(w, "user.html", struct {
		Users []models.User
		Total int64
	}{
		Users: users,
		Total: total,
	})
}
