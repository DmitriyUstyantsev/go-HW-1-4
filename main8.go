package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type Friendship struct {
	UserID1 int `json:"user_id_1"`
	UserID2 int `json:"user_id_2"`
}

var users map[int]*User

func main() {
	users = make(map[int]*User)
	r := chi.NewRouter()

	r.Post("/create", CreateUser)
	r.Post("/make_friends", MakeFriends)
	r.Delete("/user/{userID}", DeleteUser)
	r.Get("/friends/{userID}", GetFriends)
	r.Put("/user/{userID}", UpdateUserAge)

	http.ListenAndServe(":8080", r)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newUserID := len(users) + 1
	newUser.ID = newUserID
	users[newUserID] = &newUser

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"userID": newUser.ID,
		"status": 201,
	})
}

func MakeFriends(w http.ResponseWriter, r *http.Request) {
	var friends Friendship

	err := json.NewDecoder(r.Body).Decode(&friends)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users[friends.UserID1].Friends = append(users[friends.UserID1].Friends, friends.UserID2)
	users[friends.UserID2].Friends = append(users[friends.UserID2].Friends, friends.UserID1)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User 1 and User 2 are now friends",
		"status":  200,
	})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	delete(users, userID)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User deleted",
		"status":  200,
	})
}

func GetFriends(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	if user, ok := users[userID]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"friends": user.Friends,
		})
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}

func UpdateUserAge(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))
	var update struct {
		Age int `json:"age"`
	}
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user, ok := users[userID]; ok {
		user.Age = update.Age
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "The user's age has been successfully updated",
			"status":  200,
		})
	} else {
		http.Error(w, "User not found", http.StatusNotFound)
	}
}
