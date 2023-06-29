package routes

import (
	"fmt"
	"net/http"

	"github.com/aleale2121/go-eccomerce-app/models"
	"github.com/aleale2121/go-eccomerce-app/sessions"
	"github.com/aleale2121/go-eccomerce-app/utils"
)

func (rs Routes) registerGetHandler(w http.ResponseWriter, r *http.Request) {
	message, alert := sessions.Flash(r, w)
	utils.ExecuteTemplate(w, "register.html", struct {
		Alert utils.Alert
	}{
		Alert: utils.NewAlert(message, alert),
	})
}

func (rs Routes) registerPostHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	var user models.User
	user.FirstName = r.PostForm.Get("firstname")
	user.LastName = r.PostForm.Get("lastname")
	user.Email = r.PostForm.Get("email")
	user.Password = r.PostForm.Get("password")
	_, err := rs.store.NewUser(user)
	rs.checkErrRegister(err, w, r)
}

func (rs Routes) checkErrRegister(err error, w http.ResponseWriter, r *http.Request) {
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

func (rs Routes) userGetHandler(w http.ResponseWriter, r *http.Request) {
	users, err := rs.store.GetUsers()
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
