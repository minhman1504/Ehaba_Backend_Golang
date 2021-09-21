package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"ehaba_backend_golang/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	MongoClient  *mongo.Client
	DBURL        string
	DatabaseName string
}

func (m *MongoDB) Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.DBURL))
	m.MongoClient = client

	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect DB Success!")
}

func (m *MongoDB) Disconnect() {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	m.MongoClient.Disconnect(ctx)
}

// hàm này nếu trả về nil thì nghĩ là tìm thấy email trong collection User	-> Trả về Response: email đã đăng ký
//					  mongo.ErrNilCusor -> không tìm thấy email tương ứng   		-> Tiến hành tạo một User
func (m *MongoDB) Email_IsOnDatabase(ctx context.Context, email string) bool {
	result := m.MongoClient.Database(m.DatabaseName).Collection("User").FindOne(ctx, bson.M{"email": email})
	return result.Err() != mongo.ErrNoDocuments
}

//trả về true nếu Insert thành công ngược lại lỗi thì trả ra false
func (m *MongoDB) InsertOneUser(ctx context.Context, user model.User) bool {
	result, err := m.MongoClient.Database(m.DatabaseName).Collection("User").InsertOne(ctx, user)
	if err != nil {
		fmt.Println("InsertOneUser to DB error, full error:", err.Error())
	}
	fmt.Println("User Inserted: ", result.InsertedID)
	return err == nil
}
