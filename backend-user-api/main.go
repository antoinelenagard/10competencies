package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

func main() {
	http.HandleFunc("/api/register", RegisterUser)
	log.Println("Démarrage du serveur sur le port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Erreur de décodage de la requête:", err)
		return
	}

	// Create a MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur lors de la création du client MongoDB:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur lors de la connexion à la base de données:", err)
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("Erreur lors de la déconnexion du client MongoDB:", err)
		}
	}()

	// Get the users collection
	collection := client.Database("users").Collection("collection")

	// Check if the email already exists
	existingUser := User{}
	err = collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		log.Println("L'utilisateur avec l'email", user.Email, "existe déjà")
		return
	} else if err != mongo.ErrNoDocuments {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur lors de la recherche de l'utilisateur dans la base de données:", err)
		return
	}

	// Insert the user into the database
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Erreur lors de l'insertion de l'utilisateur dans la base de données:", err)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
	log.Println("L'utilisateur", user.Email, "a été inscrit avec succès")
}
