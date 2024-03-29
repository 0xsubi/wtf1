package modules

import (
	"context"
	"fmt"
	"github.com/sudhabindu1/wtf1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	client *mongo.Client
	collection *mongo.Collection
	err error
)

func init()  {
	rand.Seed(time.Now().UnixNano())
	dbUri := os.Getenv("DB_URI")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s", dbUsername, dbPassword, dbUri)
	clientOptions := options.Client().ApplyURI(connStr)

	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database("f1").Collection("messages")
}

func InsertMessage(m *models.RadioMessage) error {
	_, err := collection.InsertOne(context.TODO(), *m)
	if err != nil {
		return err
	}
	return nil
}

func FindMessageWithId(uid string) (*models.RadioMessage, error) {
	filter := bson.M { "uid": uid }
	findResult := collection.FindOne(context.TODO(), filter)
	m := models.RadioMessage{}
	if err := findResult.Decode(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

func FindMessage() (*models.RadioMessage, error) {
	cur, err := collection.Find(context.TODO(), bson.M { })
	defer func() {
		_ = cur.Close(context.TODO())
	}()

	messages := make([]models.RadioMessage, 0)

	err = cur.All(context.TODO(), &messages)
	if err != nil {
		return nil, err
	}

	return &messages[rand.Intn(len(messages))], nil
}


