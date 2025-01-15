package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"gitgub.com/fuad7161/models"
	"gitgub.com/fuad7161/utils"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front-end/home.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../front-end/home.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func UserInfoPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println("Form is not parsed sussessfully", err)
		return
	}
	handles := r.Form.Get("username")
	user, err := utils.FetchUserInfo(handles)
	fmt.Println(user.Email)
	if err != nil {
		log.Fatalf("Error fetching user info: %v", err)
	}
	tmpl, err := template.ParseFiles("../front-end/user-info.html")
	if err != nil {
		http.Error(w, "failed to load template", http.StatusInternalServerError)
		return
	}

	data := models.PageData{
		User: user,
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}

}
