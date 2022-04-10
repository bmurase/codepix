package main

import (
	"github.com/bmurase/codepix/application/grpc"
	"github.com/bmurase/codepix/infrastructure/db"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB()
	grpc.StartGrpcServer(database, 50051)
}
