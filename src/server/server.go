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
	http.HandleFunc("/", hendler)

	// Запуск сервера на порту 8080
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
