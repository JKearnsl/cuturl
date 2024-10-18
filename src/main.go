package main

import (
	"cuturl/src/adapters"
	"cuturl/src/presentation"
	"github.com/valyala/fasthttp"
	"log"
)

const listenAddr = "localhost:8080"
const databaseDir = "data"

func main() {

	db, err := adapters.Connect(databaseDir)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	ioc := NewIoc(db)
	httpRouter := presentation.MakeUrlHandler(&ioc)
	log.Printf("Starting server on %s", listenAddr)
	if err := fasthttp.ListenAndServe(listenAddr, httpRouter.Handler); err != nil {
		log.Fatalf("error in ListenAndServe: %v", err)
	}
}
