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
	Score       float64            `json:"score" bson:"score"`
	Description string             `json:"description" bson:"description"`
	Timestamp   time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	Vendor_id   string             `json:"vendor_id" bson:"vendor_id"`
	User_id     string             `json:"user_id" bson:"user_id"`
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
	fmt.Println(cursor)
	for cursor.Next(context.TODO()) {
		var review Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		// fmt.Println(review)
		reviews = append(reviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func GetReviewByID(reviewID string) (Review, error) {
	var review Review
	objectID, iderr := primitive.ObjectIDFromHex(reviewID)
	fmt.Println(objectID)
	if iderr != nil {
		return review, iderr
	}
	filter := bson.M{"_id": objectID}
	err := reviewsCollection.FindOne(context.TODO(), filter).Decode(&review)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return Review{}, fmt.Errorf("the review ID could not be found")
		}
		return review, err
	}
	return review, nil
}

func GetReviewByVendorID(vendorID string) ([]Review, error) {
	var reviews []Review
	filter := bson.M{"Vendor_id": vendorID}
	fmt.Println(vendorID)
	fmt.Println(filter)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	cursor, err := reviewsCollection.Find(context.Background(), filter)
	// fmt.Println(err)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("the vendor ID could not be found")
		}
		return nil, err
	}
	fmt.Println(cursor)
	for cursor.Next(context.TODO()) {
		var review Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		fmt.Println(review)
		reviews = append(reviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}

func CreateReview(review Review) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	review.Timestamp = time.Now()
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
	objectID, iderr := primitive.ObjectIDFromHex(reviewID)
	if iderr != nil {
		return iderr
	}
	filter := bson.M{"_id": objectID}
	result, err := reviewsCollection.DeleteOne(context.TODO(), filter)
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
	objectID, iderr := primitive.ObjectIDFromHex(reviewID)
	if iderr != nil {
		return iderr
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"Score":       updatedReview.Score,
			"Description": updatedReview.Description,
			"Timestamp":   updatedReview.Timestamp,
			"Vendor_id":   updatedReview.Vendor_id,
			"User_id":     updatedReview.User_id,
		},
	}
	result, err := reviewsCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("the review id could not be found")
	}

	return nil
}
