package app

import (
	"fmt"
	"github.com/shiyan123/marvel.sy/common/db"
)

type Database struct {
	MongoDB *db.MongoDB
}

func newDatabase(config *Config) (datebase *Database, err error) {

	mongoDB, err := db.NewMongoDB(config.MongoDB.URL, config.MongoDB.DbName, 30)
	if err != nil {
		fmt.Println("mongoDB err:", err)
		return
	}

	if err = db.EnsureIndex(mongoDB.DB()); err != nil {
		return
	}

	datebase = &Database{MongoDB: mongoDB}

	return
}

func (d *Database) Close() (err error) {
	if d.MongoDB != nil {
		d.MongoDB.Close()
	}
	return
}
