package main

import (
	"TestRestAPI/pkg/crypto"
	"TestRestAPI/pkg/mongo"
	"TestRestAPI/pkg/server"
	"log"
)

func main() {
	ms, err := mongo.NewSession("127.0.0.1:27017")
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}
	defer ms.Close()

	h := crypto.Hash{}
	u := mongo.NewUserService(ms.Copy(), "TestRestAPI", "user", &h)
	s := server.NewServer(u)

	s.Start()
}
