package db

import (
	"fmt"
	"log"
)

/**
* Author Company: PT. ASLI ISOAE SOLUSINE
* Author Url: www.isoae.id
* Author Name : Mahfuzon Akhiar
* Author Email : mahfuzon0@gmail.com
* Date : 29/12/24
* Time : 19.00
* Project Name : grpc-server
* Licensed to :
**/

func (d *dataBaseAdapter) Begin() *dataBaseAdapter {
	log.Println("begin trx")
	tx := d.db.Begin()

	if tx.Error != nil {
		log.Printf("error when create new trx %v,%v", tx.Error.Error(), "\n")
	}

	return &dataBaseAdapter{db: tx}
}

func (d *dataBaseAdapter) Commit() {
	fmt.Println("commit trx")
	err := d.db.Commit().Error
	if err != nil {
		log.Printf("error when commit trx %v,%v", err.Error(), "\n")
	}
}

func (d *dataBaseAdapter) Rollback() {
	fmt.Println("rollback trx")
	err := d.db.Rollback().Error
	if err != nil {
		log.Printf("error when rollback trx %v,%v", err.Error(), "\n")
	}
}
