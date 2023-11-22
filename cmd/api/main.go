package main

import (
	"fmt"
	"log"
	"net"
	"order/internal/config"
	"order/internal/delivery/order"
	"order/internal/repo/postgres"
	"order/pkg/order/pkg/prots"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func main() {
	if err := run(":8010"); err != nil {
		log.Fatal(err)
	}
}

func run(serv string) error {
	configDB, err := config.NewConfig("./config/database.yaml")
	if err != nil {
		return err
	}
	//INIT DB CONN
	s := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		configDB.Username, configDB.Password, configDB.Database, configDB.Host, configDB.Port,
	)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return err
	}
	myRepo := postgres.NewRepo(db)

	//GRPC USERS
	list, err := net.Listen("tcp", serv)
	if err != nil {
		return err
	}
	serverRegistration := grpc.NewServer()
	newOrder := order.NewOrder(myRepo)
	prots.RegisterOrderServiceServer(serverRegistration, newOrder)

	return serverRegistration.Serve(list)
}
