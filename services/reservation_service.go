package services

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/yeboahd24/movie-reservation-system/models"
	"gorm.io/gorm"
)

type ReservationService struct {
	DB *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{
		DB: db,
	}
}

func (rs *ReservationService) GetAvailableSeats(showtimeID string) ([]int, error) {
	var showtime models.Showtime
	if err := rs.DB.Where("id = ?", showtimeID).First(&showtime).Error; err != nil {
		return nil, err
	}
	reservations := make([]models.Reservation, 0)
	if err := rs.DB.Where("showtime_id = ?", showtimeID).Find(&reservations).Error; err != nil {
		return nil, err
	}

	var bookedSeats []int
	for _, reservation := range reservations {
		seatNumbers := strings.Split(reservation.SeatNumbers, ",")
		for _, seatStr := range seatNumbers {
			seat, err := strconv.Atoi(strings.TrimSpace(seatStr))
			if err != nil {
				return nil, err
			}
			bookedSeats = append(bookedSeats, seat)
		}
	}

	availableSeats := make([]int, showtime.AvailableSeats)
	for i := 0; i < showtime.AvailableSeats; i++ {
		availableSeats[i] = i + 1
	}

	for _, bookedSeat := range bookedSeats {
		for i, seat := range availableSeats {
			if seat == bookedSeat {
				availableSeats = append(availableSeats[:i], availableSeats[i+1:]...)
				break
			}
		}
	}

	return availableSeats, nil
}

func (rs *ReservationService) CreateReservation(reservation models.Reservation) (*models.Reservation, error) {
	reservation.ID = uuid.New()
	if err := rs.DB.Create(&reservation).Error; err != nil {
		return nil, err
	}
	return &reservation, nil
}

func (rs *ReservationService) GetUserReservations(userID string) ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	if err := rs.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (rs *ReservationService) CancelReservation(reservationID string) error {
	if err := rs.DB.Where("id = ?", reservationID).Delete(&models.Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func (rs *ReservationService) GetAllReservations() ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	if err := rs.DB.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
