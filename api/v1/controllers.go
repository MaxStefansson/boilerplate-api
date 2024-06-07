package v1

import "net/http"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("UserID, " + userID))
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("AddUser, " + userID))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("DeleteUser, " + userID))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.PathValue("userID")
	w.Write([]byte("UpdateUser, " + userID))
}
