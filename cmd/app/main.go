package main

import "latihan-golang-lala-cantik/internal/database"

func main() {
	database.Connect()
	err := database.Migrate()
	if err != nil {
		panic(err)
	}

	// TODO: BUAT NewApp struct di routes.go
	app := NewApp(&config{
		db: database.DB,
	})

	log.Fatal(app.Listen(":3000"))
}
