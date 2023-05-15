package client

import (
	"context"
	"fmt"
	"log"
	"os"

	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB            *mongo.Client
	IsDBConnected bool
)

func NewMongo(url string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	return client, nil
}

var (
	SUDOERS_DB *mongo.Collection
)

func AddSudoer(id int64) error {
	if IsCachedSudoer(id) {
		return fmt.Errorf("sudoer already exists")
	}
	if !IsDBConnected {
		CACHED_SUDOERS = append(CACHED_SUDOERS, id)
		return nil
	}
	_, err := SUDOERS_DB.InsertOne(context.Background(), map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}
	CACHED_SUDOERS = append(CACHED_SUDOERS, id)
	return nil
}

func RemoveSudoer(id int64) error {
	if !IsCachedSudoer(id) {
		return fmt.Errorf("sudoer not found")
	}
	if !IsDBConnected {
		for i, sudoer := range CACHED_SUDOERS {
			if sudoer == id {
				CACHED_SUDOERS = append(CACHED_SUDOERS[:i], CACHED_SUDOERS[i+1:]...)
				return nil
			}
		}
		return nil
	}
	_, err := SUDOERS_DB.DeleteOne(context.Background(), map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	for i, sudoer := range CACHED_SUDOERS {
		if sudoer == id {
			CACHED_SUDOERS = append(CACHED_SUDOERS[:i], CACHED_SUDOERS[i+1:]...)
			return nil
		}
	}
	return nil
}

func IsSudoer(id int64) bool {
	if !IsDBConnected {
		return IsCachedSudoer(id)
	}
	var result map[string]interface{}
	err := SUDOERS_DB.FindOne(context.Background(), map[string]interface{}{
		"id": id,
	}).Decode(&result)
	if err != nil {
		return false
	}
	return true
}

func GetSudoers() []int64 {
	if !IsDBConnected {
		return []int64{}
	}
	var result []int64
	cursor, err := SUDOERS_DB.Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return []int64{}
	}
	for cursor.Next(context.Background()) {
		var data map[string]interface{}
		err := cursor.Decode(&data)
		if err != nil {
			return []int64{}
		}
		result = append(result, data["id"].(int64))
	}
	return result
}

var (
	CACHED_SUDOERS []int64 = GetSudoers()
)

func IsCachedSudoer(id int64) bool {
	for _, sudoer := range CACHED_SUDOERS {
		if sudoer == id {
			return true
		}
	}
	return false
}

func init() {
	IsDBConnected = false
	var DB_URL string
	if dbUrl, ok := os.LookupEnv("MONGO_URL"); ok {
		DB_URL = dbUrl
	}
	mongoDb, err := NewMongo(DB_URL)
	if err != nil {
		log.Println("Hyperion: Failed to init to MongoDB")
		return
	}
	err = mongoDb.Connect(context.Background())
	if err != nil {
		log.Println("Hyperion: Failed to connect to MongoDB")
	}
	IsDBConnected = true
	DB = mongoDb
	SUDOERS_DB = DB.Database("hyperion").Collection("sudoers")
}
