package models

import (
	"context"
	"dineflow-review-service/configs"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Review struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Score       float64            `json:"score"`
	Description string             `json:"description"`
	Timestamp   string             `json:"timestamp"`
	Vendor_id   string             `json:"vendor_id"`
	User_id     string             `json:"user_id"`
}

var reviewsCollection *mongo.Collection = configs.GetCollection(configs.Db, "reviews")

func GetAllReviews() ([]Review, error) {
	var reviews []Review
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := reviewsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var review Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func GetReviewByID(reviewID string) (Review, error) {
	var review Review
	filter := bson.M{"_id": reviewID}
	err := reviewsCollection.FindOne(context.TODO(), filter).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Review{}, fmt.Errorf("the review ID could not be found")
		}
		return review, err
	}
	return review, nil
}

func CreateReview(review Review) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	reviewData := bson.M{
		"Score":       review.Score,
		"Description": review.Description,
		"Timestamp":   review.Timestamp,
		"Vendor_id":   review.Vendor_id,
		"User_id":     review.User_id,
	}
	res, err := reviewsCollection.InsertOne(ctx, reviewData)
	if err != nil {
		return err
	}
	fmt.Println("New review created with mongodb _id: " + res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

func DeleteReviewByID(reviewID string) error {
	result, err := reviewsCollection.DeleteOne(context.TODO(), bson.M{"_id": reviewID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("the review id could not be found")
	}
	return nil
}

func UpdateReviewByID(reviewID string, updatedReview Review) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": reviewID}
	update := bson.M{"$set": updatedReview}
	result, err := reviewsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("the review id could not be found")
	}

	return nil
}
