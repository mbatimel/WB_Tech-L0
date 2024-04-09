package main

import (
	"github.com/mbatimel/WB_Tech-L0/internal/repo"
	"github.com/mbatimel/WB_Tech-L0/internal/migrate"
)
func main() {
	db, err := repo.SetConfigs("config/config.yaml")
	if err != nil {
		panic(err)
	}
	db.ConnectToDatabase()
	defer db.Close()
	if err = migrate.CreateSchema(db); err != nil {
		panic(err)
	}
}