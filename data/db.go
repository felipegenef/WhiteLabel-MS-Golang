package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var DB *mongo.Client

func GetConnectionCache() (*mongo.Client,error)  {
	if DB == nil{
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	
		uri:=os.Getenv("MONGOURI")
		
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err==nil{
			return client,err
		}
		//ping the database
		err = client.Ping(ctx, nil)
		if err != nil {
			return client,err
		}
		fmt.Println("Connected to MongoDB")

		DB = client;
	
	}
	return DB,nil;
}

func EnsureConnected() error  {
	if DB == nil{
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	
		uri:=os.Getenv("MONGOURI")
		
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
		if err==nil{
			return err
		}
		//ping the database
		err = client.Ping(ctx, nil)
		if err != nil {
			return err
		}
		fmt.Println("Connected to MongoDB")

		DB = client;
	
	}
	return nil;
}


//getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {


	dbName:=os.Getenv("DB_NAME")
    collection := client.Database(dbName).Collection(collectionName)
    return collection
}