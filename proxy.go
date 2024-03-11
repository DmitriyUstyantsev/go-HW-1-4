package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRequest)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Проксируем запрос в локальный порт, где работает наше приложение
	resp, err := http.Post("http://localhost:8080"+r.URL.String(), r.Header.Get("Content-Type"), r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Читаем ответ от нашего приложения
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ обратно клиенту
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
