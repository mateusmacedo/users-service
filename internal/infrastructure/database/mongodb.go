package database

import (
	"context"
	"time"

	"github.com/mateusmacedo/users-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoUserRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

func NewMongoUserRepository(mongoURI string, database string, collection string) (*MongoUserRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoUserRepository{client: client, database: database, collection: collection}, nil
}

func (r *MongoUserRepository) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	coll := r.client.Database(r.database).Collection(r.collection)
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *MongoUserRepository) Get(ctx context.Context, id string) (*domain.User, error) {
	coll := r.client.Database(r.database).Collection(r.collection)
	var user domain.User

	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *MongoUserRepository) Delete(ctx context.Context, id string) error {
	coll := r.client.Database(r.database).Collection(r.collection)
	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoUserRepository) List(ctx context.Context, filter map[string]interface{}, limit int, offset int) ([]*domain.User, error) {
	coll := r.client.Database(r.database).Collection(r.collection)

	opts := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var users []*domain.User
	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
