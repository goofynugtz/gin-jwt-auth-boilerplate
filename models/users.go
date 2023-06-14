package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" bson:"first_name" validate:"required,min=1,max=100"`
	LastName     *string            `json:"last_name" bson:"last_name" validate:"required,min=1,max=100"`
	Password     *string            `json:"password" bson:"password" validate:"required,min=8"`
	Email        *string            `json:"email" bson:"email" validate:"email,required"`
	Phone        *string            `json:"phone" bson:"phone"`
	UserType     *string            `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	AccessToken  *string            `json:"access" bson:"access"`
	RefreshToken *string            `json:"refresh" bson:"refresh"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	UserId       string             `json:"user_id" bson:"user_id"`
}
