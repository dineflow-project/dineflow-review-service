package inits

import (
	"context"
	"dineflow-review-service/configs"
	"dineflow-review-service/models"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeDatabase() {
	db := configs.Db.Database(configs.EnvMongoDBName())
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collectionNames, err := db.ListCollectionNames(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	userCollectionExists := false
	for _, collectionName := range collectionNames {
		if collectionName == "reviews" {
			userCollectionExists = true
		}
	}
	if !userCollectionExists {
		fmt.Println("Reviews collections does not exist, creating `reviews` collection with `root` review")
		CreateRootReview()
	}
}

// This function will be run to populate the databases with an initial review
func CreateRootReview() (user models.Review, err error) {
	rootReviewInfo := models.Review{
		Score:       0,
		Description: "",
		Timestamp:   time.Now().Format("2006-01-02"),
		Vendor_id:   "0000000000000",
		User_id:     "0000000000000",
	}
	var reviewCollection *mongo.Collection = configs.GetCollection(configs.Db, "reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = reviewCollection.InsertOne(ctx, rootReviewInfo)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return rootReviewInfo, err
}
