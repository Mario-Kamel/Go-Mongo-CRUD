package controllers

import (
	"context"
	"encoding/json"
	"github/Mario-Kamel/Go-Mongo-CRUD/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(c *mongo.Client) *UserController {
	return &UserController{
		client: c,
	}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var u models.User

	err = uc.client.Database("mongo-golang").Collection("users").FindOne(context.Background(), bson.M{"_id": oid}).Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = primitive.NewObjectID()
	res, err := uc.client.Database("mongo-golang").Collection("users").InsertOne(context.Background(), u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	params := mux.Vars(r)
	id := params["id"]
	oid, err := primitive.ObjectIDFromHex(id)
	u.Id = oid
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&u)
	_, err = uc.client.Database("mongo-golang").Collection("users").UpdateOne(context.Background(), bson.M{"_id": oid}, bson.M{"$set": u})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	_, err = uc.client.Database("mongo-golang").Collection("users").DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
