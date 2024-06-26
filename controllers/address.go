package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/erenyusufduran/gocommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Query("id")
		if userId == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid code"})
			c.Abort()
			return
		}

		address, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var addresses models.Address
		addresses.Address_ID = primitive.NewObjectID()
		if err = c.BindJSON(&addresses); err != nil {
			c.IndentedJSON(http.StatusNotAcceptable, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		matchFilter := bson.D{{Key: "$match", Value: bson.D{primitive.E{Key: "_id", Value: address}}}}
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value: "$address"}}}}
		group := bson.D{
			{Key: "$group", Value: bson.D{
				primitive.E{Key: "_id", Value: "$address_id"},
				{Key: "count", Value: bson.D{primitive.E{Key: "$sum", Value: 1}}}}}}

		pointCursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{matchFilter, unwind, group})
		if err != nil {
			c.IndentedJSON(500, "internal server error")
		}

		var addressInfo []bson.M
		if err = pointCursor.All(ctx, &addressInfo); err != nil {
			panic(err)
		}

		var size int32
		for _, addressNo := range addressInfo {
			count := addressNo["count"]
			size = count.(int32)
		}

		if size < 2 {
			filter := bson.D{primitive.E{Key: "_id", Value: address}}
			update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "address", Value: addresses}}}}
			_, err := UserCollection.UpdateOne(ctx, filter, update)
			if err != nil {
				fmt.Println(err)
			}

		} else {
			c.IndentedJSON(400, "not allowed")
		}

		ctx.Done()
	}
}

func EditHomeAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdAsHex := c.Query("id")
		if userIdAsHex == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}

		userId, err := primitive.ObjectIDFromHex(userIdAsHex)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
			return
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: userId}}
		update := bson.D{
			{Key: "$set", Value: bson.D{primitive.E{Key: "address.0.house_name", Value: editAddress.House},
				{Key: "address.0.street_name", Value: editAddress.Street},
				{Key: "address.0.city_name", Value: editAddress.City},
				{Key: "address.0.pin_code", Value: editAddress.Pincode}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}

		ctx.Done()
		c.IndentedJSON(200, "successfully updated the home address")
	}
}

func EditWorkAddress() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIdAsHex := c.Query("id")
		if userIdAsHex == "" {
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid"})
			c.Abort()
			return
		}

		userId, err := primitive.ObjectIDFromHex(userIdAsHex)
		if err != nil {
			c.IndentedJSON(500, "internal server error")
			return
		}

		var editAddress models.Address
		if err := c.BindJSON(&editAddress); err != nil {
			c.IndentedJSON(http.StatusBadRequest, err.Error())
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		filter := bson.D{primitive.E{Key: "_id", Value: userId}}
		update := bson.D{
			{Key: "$set", Value: bson.D{primitive.E{Key: "address.1.house_name", Value: editAddress.House},
				{Key: "address.1.street_name", Value: editAddress.Street},
				{Key: "address.1.city_name", Value: editAddress.City},
				{Key: "address.1.pin_code", Value: editAddress.Pincode}}}}

		_, err = UserCollection.UpdateOne(ctx, filter, update)
		if err != nil {
			c.IndentedJSON(500, "something went wrong")
			return
		}

		ctx.Done()
		c.IndentedJSON(200, "successfully updated the work address")
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
