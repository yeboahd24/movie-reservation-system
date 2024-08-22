package services

import (
	"strings"

	"github.com/google/uuid"
	"github.com/yeboahd24/movie-reservation-system/models"
	"gorm.io/gorm"
)

type ShowtimeService struct {
	DB *gorm.DB
}

func NewShowtimeService(db *gorm.DB) *ShowtimeService {
	return &ShowtimeService{
		DB: db,
	}
}

func (ss *ShowtimeService) CreateShowtime(showtime models.Showtime) (*models.Showtime, error) {
	showtime.ID = uuid.New()
	if err := ss.DB.Create(&showtime).Error; err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (ss *ShowtimeService) GetShowtimes(movieID string) ([]*models.Showtime, error) {
	var showtimes []*models.Showtime
	if err := ss.DB.Where("movie_id = ?", movieID).Find(&showtimes).Error; err != nil {
		return nil, err
	}
	return showtimes, nil
}

func (ss *ShowtimeService) UpdateShowtime(showtime models.Showtime) (*models.Showtime, error) {
	if err := ss.DB.Save(&showtime).Error; err != nil {
		return nil, err
	}
	return &showtime, nil
}

func (ss *ShowtimeService) DeleteShowtime(showtimeID string) error {
	if err := ss.DB.Where("id = ?", showtimeID).Delete(&models.Showtime{}).Error; err != nil {
		return err
	}
	return nil
}

func (ss *ShowtimeService) GetAvailableSeats(showtimeID string) ([]string, error) {
	var seats string
	err := ss.DB.Table("showtimes").Select("available_seats").Where("id = ?", showtimeID).Scan(&seats).Error
	if err != nil {
		return nil, err
	}

	// Split the seats string into an array
	seatArray := strings.Split(seats, ",")
	// Trim spaces from each seat number
	for i, seat := range seatArray {
		seatArray[i] = strings.TrimSpace(seat)
	}

	return seatArray, nil
}
