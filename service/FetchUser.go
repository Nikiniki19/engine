package service

import (
	"context"
	"engine/database"
	"engine/models"
	"engine/proto"
	"errors"
	"log"

	"gorm.io/gorm"
)

type FetchUser struct {
	proto.EngineRequestServer
}

func (f *FetchUser) FetchUser(ctx context.Context, userID *proto.EngineClientID) (*proto.EngineResponse, error) {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	useradata := models.EngineUsers{
		Model: gorm.Model{
			ID: uint(userID.Id),
		},
	}
	result := db.Where("id = ?", useradata.ID).First(&useradata)
	if result.Error != nil && result.RowsAffected == 0 {
		return &proto.EngineResponse{}, errors.New("could not fetch the user")
	}
	return &proto.EngineResponse{
		Username: useradata.Username,
		Email:    useradata.Email,
	}, nil
}
