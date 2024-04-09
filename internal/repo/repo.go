package repo

import (
	"github.com/go-pg/pg/v10"
	"github.com/mbatimel/WB_Tech-L0/internal/config"
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