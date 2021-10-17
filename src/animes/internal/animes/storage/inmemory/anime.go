package inmemory

import (
	"strings"
	"sync"

	"github.com/bcariaga/anigo/src/animes/internal/animes"
)

type animeInMemoryRepository struct {
	animes []animes.Anime
}

var (
	animesOnce     sync.Once
	animesInstance *animeInMemoryRepository
)

func NewAnimeInMemoryRepository() animes.Respository {
	animesOnce.Do(func() {
		animesInstance = &animeInMemoryRepository{
			[]animes.Anime{
				{
					Title:       "!NVADE SHOW!",
					AnimeType:   "SPECIAL",
					Status:      "FINISHED",
					Episodes:    1,
					AnineSeason: animes.AnimeSeason{},
					Picture:     "https://cdn.myanimelist.net/images/anime/1718/110969.jpg",
					Thumbnail:   "https://cdn.myanimelist.net/images/anime/1718/110969t.jpg",
					Synonyms:    []string{},
					Relations:   []string{},
					Tags:        []string{},
					Source:      []string{},
				},
			},
		}
	})
	return animesInstance
}

func (r *animeInMemoryRepository) Find(term string) ([]*animes.Anime, error) {
	var matchs []*animes.Anime

	//if term is empty, returns all
	for _, v := range r.animes {
		if strings.Contains(v.Title, term) || len(term) == 0 {
			matchs = append(matchs, &v)
		}
	}

	return matchs, nil
}
