package driver

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DB holds the database connection pool
type DB struct {
	MongoDB     *mongo.Database
	MongoClient *mongo.Client
}

var dbConn = &DB{}
var clientConn []interface{}

const maxDbLifetime = 5 * time.Minute

func ConnectMongoDB(dsn, database string) (*DB, error) {
	d, err := NewClient(dsn)
	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), maxDbLifetime)
	err = d.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	dtb := d.Database(database)
	dbConn.MongoDB = dtb
	dbConn.MongoClient = d
	clientConn = append(clientConn, d)
	return dbConn, nil
}

func testDB(d *mongo.Client) error {
	ctx, _ := context.WithTimeout(context.Background(), maxDbLifetime)
	err := d.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

func NewClient(dsn string) (*mongo.Client, error) {
	clientOpts := options.Client().ApplyURI(dsn)

	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

// CloseAllConnection closes all database connection
// func CloseAllConnection() error {
// 	for _, v := range clientConn {
// 		v.Disconnect(context.TODO())
// 	}
// 	return nil
// }
