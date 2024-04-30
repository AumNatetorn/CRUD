package main

import (
	"CRUD/app"
	"CRUD/app/crud"
	"CRUD/configs"
	"CRUD/pkg/database"
	"log"
)

func init() {
	configs.Init("")
	conf := configs.GetConfig()
	log.Println(conf.App.Name)
}

func main() {
	conf := configs.GetConfig()
	db, err := database.NewMySqlDB(conf.Mysql, conf.Secrets.SecretsMysql)
	if err != nil {
		log.Panicf("error connect database: %v", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Panicf("close db error: %v", err)
		}
	}()

	var (
		customerStore = crud.NewCustomerStore(db.Client)
		handler       = crud.NewHandler(customerStore)
		server        = app.NewServer(conf)
	)
	api := server.App
	api.Post("/customers", handler.CreateCustomerHandler)
	api.Put("/customers", handler.UpdateCustomerHandler)
	api.Delete("/customers/:id", handler.DeleteCustomerHandler)
	api.Get("/customers/:id", handler.GetCustomerHandler)

	err = server.Start(conf.App.Port)
	if err != nil {
		log.Panic("error start server")
	}

	log.Print("Listening on port:", conf.App.Port)

	if err != nil {
		log.Panic("Server err")
	}

}
