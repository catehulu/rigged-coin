package models

type Prize struct {
	Name string `bson:"name" json:"name"`
	Path string `bson:"path" json:"path"`
}
