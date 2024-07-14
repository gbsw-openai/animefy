package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	model "backend/model"
)

type AnimeRepository interface {
	GetAllAnimes() ([]model.Anime, error)
	GetAnimeByID(id int) (model.Anime, error)
	AddAnime(anime model.Anime) error
	LoadAnimesFromJSON(filepath string) error
}

type InMemoryAnimeRepository struct {
	animes []model.Anime
}

func NewInMemoryAnimeRepository() *InMemoryAnimeRepository {
	return &InMemoryAnimeRepository{
		animes: []model.Anime{},
	}
}

func (r *InMemoryAnimeRepository) GetAllAnimes() ([]model.Anime, error) {
	return r.animes, nil
}

func (r *InMemoryAnimeRepository) GetAnimeByID(id int) (model.Anime, error) {
	for _, anime := range r.animes {
		if anime.Anime_id == strconv.Itoa(id) {
			return anime, nil
		}
	}
	return model.Anime{}, fmt.Errorf("anime with ID %d not found", id)
}

func (r *InMemoryAnimeRepository) AddAnime(anime model.Anime) error {
	r.animes = append(r.animes, anime)
	return nil
}

func (r *InMemoryAnimeRepository) LoadAnimesFromJSON(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read JSON data: %v", err)
	}

	var animes []model.Anime
	err = json.Unmarshal(data, &animes)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	r.animes = animes
	return nil
}
