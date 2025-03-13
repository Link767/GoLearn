package server

import (
	"fmt"
	"net/http"
)

// обработчик для маршрута "/"
func hendler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GoLearn")
}

func StartServer() {
	// Регистрация обработчика на маршрут "/"
	http.HandleFunc("/localhost", hendler) //http://localhost:8082/localhost - путь до сервера, если оставить "/" будет просто http://localhost:8082

	// Запуск сервера на порту 8080
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8082", nil) // порт который надо слушать
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
