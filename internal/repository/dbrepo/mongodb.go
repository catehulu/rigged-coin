package dbrepo

import (
	"context"
	"log"

	"github.com/catehulu/rigged-coin/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *mongoDBRepo) FindBoard(id string) *models.Board {

	boardCollection := m.DB.Collection("board")

	oId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}

	q := bson.D{{"_id", oId}}

	var board models.Board
	err = boardCollection.FindOne(context.TODO(), q).Decode(&board)
	if err != nil {
		return nil
	}

	return &board
}

func (m *mongoDBRepo) InsertBoard(board *models.Board) {

	boardCollection := m.DB.Collection("board")

	_, err := boardCollection.InsertOne(context.TODO(), board)
	if err != nil {
		log.Printf("Error when inserting : %+v", err)
	}
}

func (m *mongoDBRepo) UpdateBoard(board *models.Board) {

	boardCollection := m.DB.Collection("board")
	set := bson.D{{"$set", board}}
	_, err := boardCollection.UpdateByID(context.TODO(), board.Id, set)
	if err != nil {
		log.Printf("Error when updating : %+v", err)
	}
}
