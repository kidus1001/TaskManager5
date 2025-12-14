package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"taskmanager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var ctx = context.Background()
var coll *mongo.Collection

func ConnectWithDB() error {
	uri := "mongodb+srv://kidus1001:kidus123456yosef@cluster0.bee66s9.mongodb.net/?appName=Cluster0"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll = client.Database("TaskManagementDB").Collection("Tasks")
	if coll != nil {
		fmt.Println("Database connected successfully!")
	}
	return nil
}
func GetData() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var tasks []models.Task
	err = cursor.All(ctx, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func GetSpecificData(id string) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}
	var task models.Task

	err = coll.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateSpecificData(id string, updatedTask models.Task) error {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	res, err := coll.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func Delete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("format")
		return errors.New("invalid ID format")
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := coll.DeleteOne(ctx, bson.M{"_id": objID})

	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return err
	}
	log.Printf("Delete result: %+v", res)

	return nil
}

func Post(task models.Task) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	_, err := coll.InsertOne(ctx, task)
	return err
}
