package handlers

import (
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"profile-app/database"
	"profile-app/utils"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/auth.html"))
	tmpl.Execute(w, nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	confirm := r.FormValue("confirm")

	if confirm != password {
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		data := struct{ Error string }{"Passwords do not match"}
		tmpl.Execute(w, data)
		return
	}

	hash, hashErr := HashPassword(password)
	if hashErr != nil {
		log.Println("Hashing Error: ", hashErr)
		http.Error(w, "Server Error", 500)
		return
	}

	insertErr := database.InsertUser(username, hash)
	if insertErr != nil {
		log.Println("Insert error: ", insertErr)
		data := struct{ 
			Error string 
			Tab string			
		}{
			Error: "Username Already Exist",
			Tab: "register",
		}
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		tmpl.Execute(w, data)
		return
	}
	data := struct {
		Success string
		Tab string
	} {
		Success: "Registration Successful. Please Login",
		Tab: "login",
	}
	tmpl := template.Must(template.ParseFiles("templates/auth.html"))
	tmpl.Execute(w, data)
	return
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	hash, err := database.GetHashedPassword(username)
	if err != nil {
		log.Println("Login fetch error:", err)
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		data := struct{
			Error string
			Tab string
		}{
			Error: "User not found"
			Tab: "register"
		}
		tmpl.Execute(w, data)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		tmpl := template.Must(template.ParseFiles("templates/auth.html"))
		data := struct{ Error string }{"Incorrect Password"}
		tmpl.Execute(w, data)
		return
	}

	session, _ := utils.Store.Get(r, "session-name")
	session.Values["username"] = username
	session.Save(r, w)

	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session-name")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
