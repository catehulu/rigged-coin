package models

import (
	"math/rand"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Rigged        bool               `bson:"rigged" json:"rigged"`
	State         [][]int            `bson:"state" json:"state"`
	Size          int                `bson:"size" json:"size"`
	Prize         []Prize            `bson:"prize" json:"prize"`
	Password      string             `bson:"password"`
	ObtainedPrize int                `bson:"obtained_prize" json:"obtained_prize"`
}

func (b *Board) GetPiece(col, row int) Prize {
	selected := b.State[col][row]
	if selected != -1 {
		return b.Prize[selected]
	}

	if b.ObtainedPrize != -1 {
		return b.Prize[b.ObtainedPrize]
	}

	randArr := []int{}
	for i := 0; i < b.Size; i++ {
		randArr = append(randArr, i)
	}

	if b.Rigged {
		for i := 0; i < 100; i++ {
			randArr = append(randArr, 0)
		}
	}

	randomIndex := rand.Intn(len(randArr))
	pick := randArr[randomIndex]
	b.State[col][row] = pick
	return b.Prize[pick]
}

func (b *Board) PrizeCount() []int {

	num := b.Size
	prizeCount := make([]int, num)

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			selected := b.State[i][j]
			if selected != -1 {
				prizeCount[selected] += 1
			}
		}
	}

	return prizeCount
}

func (b *Board) CheckPrize() {

	num := b.Size
	prizeCount := make([]int, num)

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			selected := b.State[i][j]
			if selected != -1 {
				prizeCount[selected] += 1
				if prizeCount[selected] == b.Size {
					b.ObtainedPrize = selected
				}
			}
		}
	}
}
