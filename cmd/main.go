package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	commands "github.com/mateusmacedo/users-service/internal/application"
	"github.com/mateusmacedo/users-service/internal/infrastructure/database"
	http_transport "github.com/mateusmacedo/users-service/internal/infrastructure/http"
	"github.com/mateusmacedo/users-service/internal/infrastructure/services"
)

func main() {
	identity := services.NewUUIDIdentityService()
	repository, err := database.NewSQLiteUserRepository("users.db", "users")
	if err != nil {
		log.Fatal(err)
	}

	handler := commands.NewSignUpCommandHandler(repository, identity)

	ctx := context.Background()

	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req http_transport.SingUpUserRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		cmd := commands.SignUpCommand{
			Name: req.Name,
		}

		user, err := handler.Handle(ctx, &cmd)
		if err != nil {
			http.Error(w, "Error processing request", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	})

	port := ":8080"
	log.Printf("Listening on %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
