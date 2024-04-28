package database

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrCantFindProduct    = errors.New("cannot find the product")
	ErrCantDecodeProducts = errors.New("cannot find the product")
	ErrUserIdIsNotValid   = errors.New("this user is not valid")
	ErrCantUpdateUser     = errors.New("cannot update the user")
	ErrCantRemoveItemCart = errors.New("cannot remove this item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart(ctx context.Context, prodColletion, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {
	return nil
}

func RemoveCartItem(ctx context.Context, prodColletion, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {
	return nil
}

func BuyItemFromCart(ctx context.Context, userCollection *mongo.Collection, userId string) error {
	return nil
}

func InstantBuyer(ctx context.Context, prodColletion, userCollection *mongo.Collection, productId primitive.ObjectID, userId string) error {
	return nil
}
