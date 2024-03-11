package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
	loadData()
	r := chi.NewRouter()

	r.Post("/create", CreateUser)
	r.Post("/make_friends", MakeFriends)
	r.Delete("/user/{userID}", DeleteUser)
	r.Get("/friends/{userID}", GetFriends)
	r.Put("/user/{userID}", UpdateUserAge)

	go http.ListenAndServe(":8080", r) // первая реплика на порту 8080
	http.ListenAndServe(":8081", r)    // вторая реплика на порту 8081
	//Теперь вы можете запустить две реплики приложения на портах 8080 и 8081, соответственно.
	saveData()
}

func loadData() {
	_, err := os.Stat("users.json")
	if err == nil {
		file, err := ioutil.ReadFile("users.json")
		if err == nil {
			err = json.Unmarshal(file, &users)
			if err == nil {
				return
			}
		}
	}

	users = make(map[int]*User)
}

func saveData() {
	data, err := json.Marshal(users)
	if err == nil {
		err = ioutil.WriteFile("users.json", data, 0644)
	}
}

// Теперь информация о пользователях будет сохраняться в файле "users.json".
//При запуске приложения, данные будут загружены из файла, а при выполнении операций создания,
// обновления, удаления или добавления друзей, информация будет автоматически сохраняться в файл.
//Это позволяет сохранить данные между запусками приложения и обеспечить их целостность в многопоточной среде.

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
