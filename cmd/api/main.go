package main

import (
	"fmt"
	"log"

	"github.com/ErwinSalas/go-bank-api/pkg"
)

var (
	CertFilePath = "/ssl/server-cert.pem"
	KeyFilePath  = "/ssl/server.pem"
)

func main() {
	config, err := pkg.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	ds, err := pkg.NewPostgresDatastore(config)
	if err != nil {
		log.Fatal(err)
	}
	service := pkg.NewAccountService(ds)
	accountHandler := pkg.NewAccountHandler(service)

	port := fmt.Sprintf(":%s", config.Port)
	fmt.Printf("server running on port %s", config.Port)
	server := pkg.NewAPIServer(port, accountHandler)
	server.Run()
}
