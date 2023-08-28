package UserController

import (
	"crud-with-go/entities"
	UserModel "crud-with-go/models/User"
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
	users := UserModel.GetAll()

	// Tambahkan nomor urut ke setiap user
	for i := 0; i < len(users); i++ {
		users[i].No = i + 1
	}

	data := map[string]any{
		"users": users,
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

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
