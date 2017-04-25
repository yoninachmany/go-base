package main

import (
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User is a person on the site.
type User struct {
	gorm.Model
	Confirmed    bool   `gorm:"default:'true'"` // reverse when confirmation flow
	FirstName    string `gorm:"size:64;index"`
	LastName     string `gorm:"size:64;index"`
	Email        string `gorm:"size:64;unique;index"`
	PasswordHash string `gorm:"size:128"`
}

var templates = template.Must(template.ParseFiles("templates/index.html", "templates/about.html", "templates/register.html", "templates/login.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	type Data struct {
		LoggedIn bool
	}

	data := Data{false}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index", data)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/about.html", "templates/base.html"))
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/account/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			renderTemplate(w, "register", data)
		} else {
			r.ParseForm()
			FirstName := r.Form["first_name"][0]
			LastName := r.Form["last_name"][0]
			Email := r.Form["email"][0]
			Password := r.Form["password"][0]
			user := User{FirstName: FirstName, LastName: LastName, Email: Email, PasswordHash: Password}
			db.Create(&user)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	})
	http.HandleFunc("/account/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			renderTemplate(w, "login", data)
		} else {
			r.ParseForm()
			Email := r.Form["email"][0]
			FormPassword := r.Form["password"][0]
			var user User
			db.First(&user, "Email = ?", Email)
			if user.PasswordHash == FormPassword {
				data.LoggedIn = true
				http.Redirect(w, r, "/", http.StatusFound)
			} else {
				renderTemplate(w, "login", data)
			}
		}
	})
	http.HandleFunc("/account/logout", func(w http.ResponseWriter, r *http.Request) {
		data.LoggedIn = false
		http.Redirect(w, r, "/", http.StatusFound)
	})
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
