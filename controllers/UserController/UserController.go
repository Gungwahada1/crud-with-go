package UserController

import (
	"crud-with-go/entities"
	ReligionModel "crud-with-go/models/Religion"
	UserModel "crud-with-go/models/User"
	"strconv"

	// ReligionModel "crud-with-go/models/Religion"
	// RoleModel "crud-with-go/models/Role"
	"crypto/sha256"
	"encoding/hex"
	"html/template"
	"net/http"
	"time"
)

func HashPassword(password string) string {
	// Membuat hash dari password menggunakan SHA-256
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	return hashedPassword
}

func Index(w http.ResponseWriter, r *http.Request) {
	user := UserModel.GetAll()

	// Tambahkan nomor urut ke setiap user
	for i := 0; i < len(user); i++ {
		user[i].No = i + 1
	}

	data := map[string]any{
		"user": user,
	}

	tmpl := template.Must(template.ParseFiles("./views/user/index.html"))

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./views/user/create.html"))

		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
	if r.Method == "POST" {
		var user entities.User

		user.RoleId = r.FormValue("role_id")
		user.ReligionId = r.FormValue("religion_id")
		user.Email = r.FormValue("email")
		user.Username = r.FormValue("username")
		password := r.FormValue("username")
		hashedPassword := HashPassword(password)
		user.Password = hashedPassword
		user.CreatedAt = time.Now()
		user.UpdatedAt = time.Now()

		if ok := UserModel.Create(user); !ok {
			tmpl := template.Must(template.ParseFiles("./views/user/create.html"))

			err := tmpl.Execute(w, nil)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./views/user/edit.html"))

		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		user := UserModel.GetById(id)
		religion := ReligionModel.GetAll()
		data := map[string]any{
			"user":     user,
			"religion": religion,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			panic(err)
		}
	}

	if r.Method == "POST" {
		var user entities.User

		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		user.Email = r.FormValue("email")
		user.Username = r.FormValue("username")
		user.ReligionId = r.FormValue("religion_id")
		user.RoleId = r.FormValue("role_id")
		password := r.FormValue("password")
		hashedPassword := HashPassword(password)
		user.Password = hashedPassword
		user.UpdatedAt = time.Now()

		if ok := UserModel.Update(id, user); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	if err := UserModel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
