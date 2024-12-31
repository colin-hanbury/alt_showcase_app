package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	addUser(user UserInput, ctx context.Context) (User, error)
	getUser(id string, ctx context.Context) (User, error)
}

type userRepo struct {
	client *mongo.Client
}

func NewUserRepo(client *mongo.Client) UserRepo {
	return &userRepo{client: client}
}

func (repo *userRepo) addUser(userInput UserInput, ctx context.Context) (User, error) {
	collection := repo.client.Database("dbName").Collection("collectionName")
	insertOneResult, err := collection.InsertOne(ctx, userInput)

	if err != nil {
		return User{}, err
	}
	id := insertOneResult.InsertedID.(primitive.ObjectID).String()
	user := User{_id: id, name: userInput.Name, nationality: userInput.Nationality}
	return user, err
}

func (repo *userRepo) getUser(id string, ctx context.Context) (User, error) {
	collection := repo.client.Database("dbName").Collection("collectionName")
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	var user User
	var err error
	collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}
