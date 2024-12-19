// run docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -d postgres:17.1
// Using gorilla/mux for http handlers with middleware
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "mydb"
)

var db *sql.DB

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	postgres_connection()
	defer db.Close()

	router := mux.NewRouter()
	routes(router)

	fmt.Println("Server is listening on port 8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func routes(router *mux.Router) {

	router.HandleFunc("/users", authMiddleware(getUsers)).Methods("GET")
	router.HandleFunc("/users", authMiddleware(addUser)).Methods("POST")
	router.HandleFunc("/users/{id}", authMiddleware(updateUser)).Methods("PUT")
	router.HandleFunc("/users/{id}", authMiddleware(deleteUser)).Methods("DELETE")
}

func postgres_connection() {

	pgConnStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := sql.Open("postgres", pgConnStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	db = conn
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	fmt.Println("Connected to the PostgreSQL database")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", user.Name, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User added successfully")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User updated successfully")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted successfully")
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "Bearer secret" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
