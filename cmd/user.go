package cmd

import (
	"log"
	"time"

	"github.com/anyshare/anyshare-common/mongodb"
	"github.com/anyshare/anyshare-common/schemas"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateDefaultUser(ctx *cli.Context) error {
	userCollection := mongodb.GetCollection("users")
	user := schemas.User{}
	userCollection.FindOne(ctx.Context, bson.M{"email": "gthienphu@gmail.com"}).Decode(&user)
	if user.ID.IsZero() {
		hash, _ := bcrypt.GenerateFromPassword([]byte("matkhau"), 10)
		result, err := userCollection.InsertOne(ctx.Context, schemas.User{
			Email:    "gthienphu@gmail.com",
			Password: string(hash),
			FullName: "Giang Thiên Phú",
			JoinTime: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		log.Println("Created user: ", result.InsertedID)
		user.ID = result.InsertedID.(primitive.ObjectID)
	} else {
		log.Println("Default user existed!", user)
	}

	adminCollection := mongodb.GetCollection("admins")
	result, err := adminCollection.InsertOne(ctx.Context, schemas.Admin{
		UserID:   user.ID,
		Roles:    []string{"owner"},
		JoinTime: time.Now().Unix(),
	})
	if err != nil {
		log.Println(result)
		return err
	}
	log.Println("Created admin: ", result.InsertedID)
	return nil
}

func FakeUser(ctx *cli.Context) error {
	userCollection := mongodb.GetCollection("users")
	hash, _ := bcrypt.GenerateFromPassword([]byte("matkhau"), 10)
	for i := 0; i < 100; i++ {
		result, err := userCollection.InsertOne(ctx.Context, schemas.User{
			Email:    gofakeit.Email(),
			Password: string(hash),
			FullName: gofakeit.Name(),
			Address:  gofakeit.Address().Address,
			JoinTime: time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		log.Println("Created user: ", result.InsertedID)
	}
	return nil
}
