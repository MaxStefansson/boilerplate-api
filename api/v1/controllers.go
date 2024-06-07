package v1

import (
	"net/http"
)

// GetUsers godoc
// @Summary      Show an user
// @Description  get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Router       /users/{id} [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("UserID, " + userID))
}

// AddUser godoc
// @Summary      Add an user
// @Description  Add an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Router       /users/{id} [post]
func AddUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("AddUser, " + userID))
}

// DeleteUser godoc
// @Summary      Delete an user
// @Description  Delete an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("DeleteUser, " + userID))
}

// DeleteUser godoc
// @Summary      Update an user
// @Description  Update an user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("UpdateUser, " + userID))
}
