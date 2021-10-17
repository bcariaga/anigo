package animes

import (
	"fmt"
)

//region ValueTypes

type AnimeType string

const (
	TV           AnimeType = "TV"
	MOVIE        AnimeType = "MOVIE"
	OVA          AnimeType = "OVA"
	ONA          AnimeType = "ONA"
	SPECIAL      AnimeType = "SPECIAL"
	ANIMEUNKNOWN AnimeType = "UNKNOWN"
)

func NewAnimeType(animeType string) (AnimeType, error) {
	var r AnimeType
	switch an := AnimeType(animeType); an {
	case TV, MOVIE, OVA, ONA, ANIMEUNKNOWN, SPECIAL:
		r = an
	default:
		return r, fmt.Errorf("can't parse \"%s\" as anime type", animeType)
	}
	return r, nil
}

type Status string

const (
	FINISHED      Status = "FINISHED"
	ONGOING       Status = "ONGOING"
	UPCOMING      Status = "UPCOMING"
	STATUSUNKNOWN Status = "UNKNOWN"
)

func NewAnimeStatus(status string) (Status, error) {
	var r Status
	switch s := Status(status); s {
	case FINISHED:
	case ONGOING:
	case UPCOMING:
	case STATUSUNKNOWN:
		r = s
	default:
		return r, fmt.Errorf("can't parse \"%s\" as anime status", status)
	}

	return r, nil
}

type Season string
type AnimeSeason struct {
	Season Season
	Year   int
}

const (
	SPRING    Season = "SPRING"
	SUMMER    Season = "SUMMER"
	FALL      Season = "FALL"
	WINTER    Season = "WINTER"
	UNDEFINED Season = "UNDEFINED"
)

func NewSeason(season string, year int) (AnimeSeason, error) {
	var r AnimeSeason
	var ar Season
	switch s := Season(season); s {
	case SPRING:
	case SUMMER:
	case FALL:
	case WINTER:
	case UNDEFINED:
		ar = s
	default:
		return r, fmt.Errorf("can't parse \"%s\" as anime season", season)
	}
	return AnimeSeason{Season: ar, Year: year}, nil
}

//endregion ValueTypes

type Anime struct {
	Title       string
	AnimeType   AnimeType
	Status      Status
	Episodes    int
	AnineSeason AnimeSeason
	Picture     string
	Thumbnail   string
	Synonyms    []string
	Relations   []string
	Tags        []string
	Source      []string
}

func NewAnime(
	title string,
	animeType AnimeType,
	status Status,
	episodes int,
	season AnimeSeason,
	picture string,
	thumbail string,
	synonyms []string,
	relations []string,
	tags []string,
	source []string) *Anime {
	fmt.Println(animeType)
	return &Anime{
		Title:       title,
		AnimeType:   animeType,
		Status:      status,
		Episodes:    episodes,
		AnineSeason: season,
		Picture:     picture,
		Thumbnail:   thumbail,
		Synonyms:    synonyms,
		Relations:   relations,
		Tags:        tags,
		Source:      source,
	}
}

//region repository

type Respository interface {
	Find(term string) ([]*Anime, error)
}

//endregion repository
