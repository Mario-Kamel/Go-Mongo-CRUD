package main

import (
	"context"
	"fmt"
	"github/Mario-Kamel/Go-Mongo-CRUD/pkg/controllers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load("./.env")
	r := mux.NewRouter()
	uc := controllers.NewUserController(getClient())
	r.HandleFunc("/user/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", uc.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", uc.DeleteUser).Methods("DELETE")

	server := http.Server{
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Addr:         ":8000",
		Handler:      r,
	}
	fmt.Printf("Server is listening on port %v...\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func getClient() *mongo.Client {
	fmt.Println("Connecting to MongoDB...", os.Getenv("MONGO_URI"))
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
