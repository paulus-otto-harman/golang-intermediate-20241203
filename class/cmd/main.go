package main

import (
	"flag"
	"log"
	"project/class/infra"
	"project/class/routes"

	_ "project/class/docs"
)

// @title Voucher API
// @version 1.0
// @description This is a sample server for a Swagger API.
// @termsOfService http://example.com/terms/
// @contact.name Paulus Otto Harman
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey ID-KEY
// @in header
// @name ID-KEY
// @securityDefinitions.apikey token
// @in header
// @name token

func main() {
	migrateDb := flag.Bool("m", false, "use this flag to migrate database")
	seedDb := flag.Bool("s", false, "use this flag to seed database")
	flag.Parse()

	ctx, err := infra.NewServiceContext(*migrateDb, *seedDb)
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
