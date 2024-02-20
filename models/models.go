package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CategoryID  int                `json:"Category_Id,omitempty" bson:"Category_Id,omitempty"`
	BookName    string             `json:"Book_Name,omitempty" bson:"Book_Name,omitempty"`
	AuthorNames []string           `json:"Author_Name,omitempty" bson:"Author_Name,omitempty"`
	Description string             `json:"Discription,omitempty" bson:"Discription,omitempty"`
	BookCount   int                `json:"Book_Count,omitempty" bson:"Book_Count,omitempty"`
	CatNames    []string           `json:"Cat_Name,omitempty" bson:"Cat_Name,omitempty"`
}

type Booking struct {
	Borrowed_At time.Time
	User_ID     *string
	Book_ID     *string
	Date_taken  *string
	Return_date *string
	Return      *bool
	Fine        *uint64
}

type User struct {
	User_ID   *string
	User_Name *string
	Phone     *string
	Borrowed  []Booking
	Limit     *int
}
