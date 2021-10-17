package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	httpserver "github.com/bcariaga/anigo/src/animes/cmd/animes-api/internal/server/http"
	"github.com/bcariaga/anigo/src/animes/internal/animes/fetching"
	"github.com/bcariaga/anigo/src/animes/internal/animes/storage/mongo"
)

const (
	ApiHostDefault = "localhost"
	ApiPortDefault = 3000

	ReadTimeoutDefault  = 10 // seconds
	WriteTimeoutDefault = 10 // seconds
)

/*
MONGO_INITDB_DATABASE=anigo
MONGO_INITDB_COLLECTION=animes
*/
func main() {

	repository, closeRepo := mongo.NewMongoDbRepository(
		mongo.MongoDbOpt{
			Url:            "mongodb://shinji:eva01@localhost:27017",
			TimeOutSeconds: 30,
			DatabaseName:   "anigo",
			CollectionName: "animes",
		})
	defer closeRepo()
	fetchService := fetching.NewService(repository)
	apiAddress := fmt.Sprintf("%s:%d", ApiHostDefault, ApiPortDefault)

	handler, err := httpserver.MainHandler(fetchService)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:         apiAddress,
		Handler:      handler,
		ReadTimeout:  ReadTimeoutDefault * time.Second,
		WriteTimeout: ReadTimeoutDefault * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
