package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"profile-app/database"
	"profile-app/utils"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := utils.Store.Get(r, "session-name")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		bio := r.FormValue("bio")
		hobbies := r.FormValue("hobbies")
		friends := r.FormValue("friends")

		err := database.UpdateProfile(username, name, email, bio, hobbies, friends)
		if err != nil {
			http.Error(w, "Failed to update profile", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/profile?updated=true", http.StatusSeeOther)
		return
	}

	profile, err := database.GetProfile(username)
	if err != nil {
		http.Error(w, "Profile not found", http.StatusInternalServerError)
		return
	}

	if r.URL.Query().Get("updated") == "true" {
		profile.Success = "Profile Updated Successfully"
	}
	fmt.Printf("%+v\n", profile)

	tmpl := template.Must(template.ParseFiles("templates/profile.html"))
	tmpl.Execute(w, profile)
}
