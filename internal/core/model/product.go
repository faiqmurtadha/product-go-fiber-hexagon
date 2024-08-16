package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name,omitempty"`
	Stock int                `json:"stock,omitempty" bson:"stock,omitempty"`
	Price int                `json:"price,omitempty" bson:"price,omitempty"`
}

type PaginatedProduct struct {
	Products   []Product `json:"products"`
	TotalCount int64     `json:"totalCount"`
	Page       int64     `json:"page"`
	Limit      int64     `json:"limit"`
}
