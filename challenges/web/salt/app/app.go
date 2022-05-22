package main

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
)

var (
	key   []byte
	store *sessions.CookieStore
)

type Macaca struct {
	id       int
	name     string
	password string
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func auth(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		name, ok := session.Values["macaca"].(string)
		if !ok || name == "" {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		f(w, r)
	}
}

func getMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func getMacaca(name string, password string) Macaca {
	db, err := sql.Open("sqlite3", "./files/macaca.db")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM MacacaHub WHERE name=$1 AND password=$2;", name, getMD5Hash(password))
	checkErr(err)

	macaca := Macaca{}
	if rows.Next() {
		rows.Scan(&macaca.id, &macaca.name, &macaca.password)
	}
	return macaca
}

func index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	name, _ := session.Values["macaca"].(string)

	templ := template.Must(template.ParseFiles("./templates/index.html"))
	FL4G := os.Getenv("FL4G")
	templ.Execute(w, struct {
		Name string
		FLAG string
	}{Name: name, FLAG: FL4G})
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	tmpl := template.Must(template.ParseFiles("./templates/login.html"))

	switch r.Method {
	case http.MethodGet:
		tmpl.Execute(w, struct{ Failed bool }{Failed: false})

	case http.MethodPost:
		name := r.FormValue("name")
		password := r.FormValue("password")

		macaca := getMacaca(name, password)
		if macaca == (Macaca{}) {
			tmpl.Execute(w, struct{ Failed bool }{Failed: true})
			return
		}
		session.Values["macaca"] = macaca.name
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func init() {
	key = make([]byte, 128)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}

func main() {
	fs := http.FileServer(http.Dir("./files"))
	http.Handle("/files/", http.StripPrefix("/files/", fs))
	http.HandleFunc("/", auth(index))
	http.HandleFunc("/login", login)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
