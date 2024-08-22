package services

import (
	"github.com/google/uuid"
	"github.com/yeboahd24/movie-reservation-system/models"
	"gorm.io/gorm"
)

type MovieService struct {
	DB *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		DB: db,
	}
}

func (ms *MovieService) CreateMovie(movie models.Movie) (*models.Movie, error) {
	movie.ID = uuid.New()
	if err := ms.DB.Create(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (ms *MovieService) GetMovies() ([]*models.Movie, error) {
	var movies []*models.Movie
	if err := ms.DB.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (ms *MovieService) UpdateMovie(movie models.Movie) (*models.Movie, error) {
	if err := ms.DB.Save(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (ms *MovieService) DeleteMovie(movieID string) error {
	if err := ms.DB.Where("id = ?", movieID).Delete(&models.Movie{}).Error; err != nil {
		return err
	}
	return nil
}
