package migrate

import (
	"github.com/go-pg/pg/orm"
	"github.com/mbatimel/WB_Tech-L0/internal/model"
	"github.com/mbatimel/WB_Tech-L0/internal/repo"
)


func CreateSchema(db *repo.DataBase) error {
	models := []interface{}{
		(*model.Order)(nil),
		(*model.Delivery)(nil),
		(*model.Items)(nil),
		(*model.Payment)(nil),
	}
	for _, model := range models {
		op := orm.CreateTableOptions{}
		err := db.DB.Model(model).CreateTable(&op)
		if err != nil {
			return err
		}
	}
	return nil
}