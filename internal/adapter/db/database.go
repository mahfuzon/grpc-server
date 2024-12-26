package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 26/12/24
* Time : 17.56
* Project Name : grpc-server
* Licensed to :
**/
type dataBaseAdapter struct {
	db *gorm.DB
}

func NewDatabaseAdapter(conn *sql.DB) (*dataBaseAdapter, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database gorm: %w", err)
	}

	return &dataBaseAdapter{db: db}, nil
}
