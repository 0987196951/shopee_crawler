package database

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shopee.rd/crawler"
	"shopee.rd/utils"
)

func Create_product(collection *mongo.Collection, product crawler.Product) error {
	err := collection.FindOne(context.TODO(), bson.D{{utils.PRODUCT_ID, product.ProductID}})
	if err == nil {
		return errors.New("Product existed")
	}
	_, errr := collection.InsertOne(context.TODO(), product)
	return errr
}
func Read_one_product_by_id(collection *mongo.Collection, product_id string) (crawler.Product, error) {
	var product crawler.Product
	err := collection.FindOne(context.TODO(), bson.D{{utils.PRODUCT_ID, product_id}}).Decode(&product)
	return product, err
}

func Update_product_by_id(collection *mongo.Collection, product crawler.Product) error {
	result, err := collection.UpdateOne(context.TODO(), bson.D{{utils.PRODUCT_ID, product.ProductID}}, bson.D{{"$set", product}})
	fmt.Println(result.MatchedCount)
	return err
}

func Delete_product_by_id(collection *mongo.Collection, product_id string) error {
	result, err := collection.DeleteOne(context.TODO(), bson.D{{utils.PRODUCT_ID, product_id}})
	fmt.Println(result.DeletedCount)
	return err
}
