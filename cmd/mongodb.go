package cmd

import (
	"log"

	"github.com/anyshare/anyshare-common/mongodb"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongodbIndexes(ctx *cli.Context) error {
	result, err := mongodb.GetCollection("users").Indexes().CreateMany(ctx.Context, []mongo.IndexModel{
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "joinTime", Value: -1}}},
	})
	if err != nil {
		panic(err)
	}
	log.Println(result)
	result, err = mongodb.GetCollection("admins").Indexes().CreateMany(ctx.Context, []mongo.IndexModel{
		{Keys: bson.D{{Key: "userId", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "joinTime", Value: -1}}},
		{Keys: bson.D{{Key: "roles", Value: 1}}},
	})
	if err != nil {
		panic(err)
	}
	log.Println(result)
	return nil
}
