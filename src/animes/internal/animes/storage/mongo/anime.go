package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bcariaga/anigo/src/animes/internal/animes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AnimeSeasonDto struct {
	Season string
	Year   int
}
type AnimeDto struct {
	Title       string
	Type        string
	Status      string
	Episodes    int
	AnimeSeason AnimeSeasonDto
	Picture     string
	Thumbnail   string
	Synonyms    []string
	Relations   []string
	Tags        []string
	Source      []string
}

func (dto *AnimeDto) ConvertToDomain() *animes.Anime {
	//TODO: manage errors
	// fmt.Print("dto")
	// fmt.Println(dto.Type)
	aType, err := animes.NewAnimeType(dto.Type)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(aType)
	aStatus, err := animes.NewAnimeStatus(dto.Status)
	if err != nil {
		log.Fatal(err)
	}
	aSeason, err := animes.NewSeason(dto.AnimeSeason.Season, dto.AnimeSeason.Year)
	if err != nil {
		log.Fatal(err)
	}
	return animes.NewAnime(dto.Title, aType, aStatus, dto.Episodes, aSeason, dto.Picture, dto.Thumbnail, dto.Synonyms, dto.Relations, dto.Tags, dto.Source)
}

type MongoDbOpt struct {
	Url            string
	TimeOutSeconds int
	DatabaseName   string
	CollectionName string
}

func NewClient(ctx context.Context, serverURL string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(serverURL))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func OpenCollection(
	client *mongo.Client,
	database string,
	collectionName string) *mongo.Collection {

	return client.Database(database).Collection(collectionName)
}

type animeMongoRepository struct {
	colletion *mongo.Collection
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewMongoDbRepository(opt MongoDbOpt) (repository animes.Respository, close func()) {
	fmt.Print("new repo!")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(opt.TimeOutSeconds)*time.Second)
	close = cancel
	client := NewClient(ctx, opt.Url)
	collection := OpenCollection(client, opt.DatabaseName, opt.CollectionName)
	repository = &animeMongoRepository{colletion: collection, ctx: ctx, cancel: cancel}
	return
}

func (r *animeMongoRepository) Find(term string) ([]*animes.Anime, error) {
	var results []*animes.Anime
	filter := bson.M{
		"title": bson.M{
			"$regex":   fmt.Sprintf(".*%s.*", term),
			"$options": "i",
		}}
	cur, err := r.colletion.Find(r.ctx, filter)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("fail on find!")
	}
	for cur.Next(r.ctx) {
		var s AnimeDto
		err := cur.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, s.ConvertToDomain())
	}
	return results, nil
}
