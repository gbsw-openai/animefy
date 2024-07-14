package services

import (
	model "backend/model"
	"backend/repository"
	"context"
	"strconv"
)

type AnimeService struct {
	repo repository.AnimeRepository
}

func NewAnimeService(repo repository.AnimeRepository) *AnimeService {
	return &AnimeService{
		repo: repo,
	}
}

func (s *AnimeService) RecommendAnime(ctx context.Context, preferences model.UserPreferences) ([]model.Anime, error) {
	allAnimes, err := s.repo.GetAllAnimes()
	if err != nil {
		return nil, err
	}

	var recommendedAnimes []model.Anime

	for _, anime := range allAnimes {
		rating, err := strconv.ParseFloat(anime.Rating, 64)
		if err != nil {
			continue
		}

		if preferences.MinRating > 0 && rating < preferences.MinRating {
			continue
		}

		if len(preferences.Genres) > 0 && !isGenrePreferred(preferences.Genres, anime.Genre) {
			continue
		}

		recommendedAnimes = append(recommendedAnimes, anime)
	}

	return recommendedAnimes, nil
}

func (s *AnimeService) EvaluateAnime(ctx context.Context, animeID int, preferences model.UserPreferences) (float64, error) {
	anime, err := s.repo.GetAnimeByID(animeID)
	if err != nil {
		return 0.0, err
	}

	rating, err := strconv.ParseFloat(anime.Rating, 64)
	if err != nil {
		return 0.0, err
	}

	return rating, nil
}

func isGenrePreferred(preferredGenres []string, animeGenre string) bool {
	for _, genre := range preferredGenres {
		if genre == animeGenre {
			return true
		}
	}
	return false
}
