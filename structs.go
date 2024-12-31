package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID     string             `bson:"userId" json:"userId"`
	EntityType string             `bson:"entityType" json:"entityType"`
	EntityID   string             `bson:"entityId" json:"entityId"`
	Rating     int                `bson:"rating" json:"rating"`
	Comment    string             `bson:"comment" json:"comment"`
	Date       time.Time          `bson:"date" json:"date"`
}

// type Review struct {
// 	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	EntityID   string             `json:"entity_id" bson:"entity_id"`
// 	EntityType string             `json:"entity_type" bson:"entity_type"` // "event" or "place"
// 	Reviewer   primitive.ObjectID `json:"reviewer" bson:"reviewer"`
// 	Comment    string             `json:"comment,omitempty" bson:"comment,omitempty"`
// 	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
// 	Content    string             `bson:"content" json:"content"`
// 	ReviewID   string             `json:"reviewid" bson:"reviewid"`
// 	// EventID     string             `json:"eventid" bson:"eventid"` // Reference to Event ID
// 	UserID      string    `json:"userid" bson:"userid"` // Reference to User ID
// 	Rating      int       `json:"rating" bson:"rating"` // Rating out of 5
// 	Date        time.Time `json:"date" bson:"date"`     // Date of the review
// 	Likes       int       `json:"likes,omitempty" bson:"likes,omitempty"`
// 	Dislikes    int       `json:"dislikes,omitempty" bson:"dislikes,omitempty"`
// 	Attachments []Media   `json:"attachments,omitempty" bson:"attachments,omitempty"`
// 	CreatedAt   time.Time `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
// }
