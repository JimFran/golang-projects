package handlers

import (
	"database/sql"
	"net/http"

	"backend-swagger/models"
	"backend-swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type UserHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUsers(params operations.GetUsersParams) middleware.Responder {
	rows, err := h.db.Query("SELECT * FROM users")
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error obtaning users")
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return middleware.Error(http.StatusInternalServerError, "Error processing users")
		}
		users = append(users, &user)
	}

	return operations.NewGetUsersOK().WithPayload(users)
}

func (h *UserHandler) AddUser(params operations.PostUsersParams) middleware.Responder {
	values := params.Body

	_, err := h.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", values.Name, values.Email)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error adding user")
	}

	//users := []*models.User{values}
	//return operations.NewGetUsersOK().WithPayload(users)
	return operations.NewPostUsersCreated()
}

func (h *UserHandler) UpdateUser(params operations.PutUsersIDParams) middleware.Responder {
	valueID := params.ID
	valueBody := params.Body

	_, err := h.db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", valueBody.Name, valueBody.Email, valueID)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error adding user")
	}

	//users := []*models.User{valueBody}
	//return operations.NewGetUsersOK().WithPayload(users)
	return operations.NewPutUsersIDOK()
}

func (h *UserHandler) DeleteUser(params operations.DeleteUsersIDParams) middleware.Responder {
	valueID := params.ID

	_, err := h.db.Exec("DELETE FROM users WHERE id=$1", valueID)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error deleting user")
	}

	return operations.NewDeleteUsersIDOK()
}
