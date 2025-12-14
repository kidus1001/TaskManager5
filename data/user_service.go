package data

import (
	"context"
	"errors"
	"strings"
	"taskmanager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (*models.User, error) {
	username = strings.TrimSpace(username)
	if username == "" || password == "" {
		return nil, errors.New("username and password required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := coll.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("username already exists")
	}

	// Determine role: first user becomes admin
	total, err := coll.CountDocuments(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	role := "user"
	if total == 0 {
		role = "admin"
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := models.User{
		ID:           primitive.NewObjectID(),
		Username:     username,
		PasswordHash: string(hashed),
		Role:         role,
		CreatedAt:    time.Now().UTC(),
	}

	_, err = coll.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	u.PasswordHash = ""
	return &u, nil
}

func AuthenticateUser(username, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var u models.User
	err := coll.FindOne(ctx, bson.M{"username": username}).Decode(&u)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	u.PasswordHash = ""
	return &u, nil
}

func PromoteUser(username string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := coll.UpdateOne(ctx,
		bson.M{"username": username},
		bson.M{"$set": bson.M{"role": "admin"}})
	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
