package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/erenyusufduran/gocommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdAsHex := c.Query("id")
		if userIdAsHex == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid search index"})
			c.Abort()
			return
		}

		addresses := make([]models.Address, 0)
		userId, err := primitive.ObjectIDFromHex(userIdAsHex)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: userId}}
		update := bson.D{{Key: "$set", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(404, "wrong command")
			return
		}

		ctx.Done()
		c.IndentedJSON(200, "Succesfully deleted")
	}
}
