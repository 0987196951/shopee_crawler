package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"shopee.rd/utils"
)

func Connect_to_database(database_uri string) *mongo.Database {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(database_uri))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return client.Database(utils.DATABASE_NAME)
}
func Disconnect_database(database *mongo.Database) error {
	var client *mongo.Client
	client = database.Client()
	err := client.Disconnect(context.TODO())
	return err
}
