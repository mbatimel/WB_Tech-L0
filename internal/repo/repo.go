package repo

import (
	"github.com/go-pg/pg"
	"github.com/mbatimel/WB_Tech-L0/internal/config"
	"github.com/mbatimel/WB_Tech-L0/internal/model"
)
type DataBase struct {
	DB *pg.DB
	config *config.Repo
}
func SetConfigs(path string) (*DataBase, error) {
 config, err := config.NewConfigDB(path)
 if err != nil {
	return nil, err
 }
 return &DataBase{nil, config}, nil
}

func (db *DataBase) AddOrder(data model.Order) error {
 if _, err := db.DB.Model(&data).Insert(); err != nil {
	return err
 }
 return nil
}
func (db *DataBase) Close(){
	db.DB.Close()
}

func(db *DataBase)ConnectToDatabase(){
	db.DB =pg.Connect(&pg.Options{
		User: db.config.User,
		Password: db.config.Password,
		Database: db.config.Database,
	})
}