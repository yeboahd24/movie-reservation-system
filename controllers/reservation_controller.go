package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yeboahd24/movie-reservation-system/models"
	"github.com/yeboahd24/movie-reservation-system/services"
)

type ReservationController struct {
	ReservationService *services.ReservationService
	ShowtimeService    *services.ShowtimeService
}

func NewReservationController(reservationService *services.ReservationService, showtimeService *services.ShowtimeService) *ReservationController {
	return &ReservationController{
		ReservationService: reservationService,
		ShowtimeService:    showtimeService,
	}
}

func (rc *ReservationController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeId")
	fmt.Printf("Fetching seats for showtime ID: %s\n", showtimeID)

	seats, err := rc.ShowtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		fmt.Printf("Error fetching seats: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Available seats: %+v\n", seats)

	c.JSON(http.StatusOK, seats)
}

func (rc *ReservationController) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract userID from the token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to uuid.UUID
	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Set the UserID from the token
	reservation.UserID = userUUID

	newReservation, err := rc.ReservationService.CreateReservation(reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newReservation)
}

func (rc *ReservationController) GetUserReservations(c *gin.Context) {
	// Extract user ID from the token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to string (assuming it's stored as a string in the token)
	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservations, err := rc.ReservationService.GetUserReservations(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func (rc *ReservationController) CancelReservation(c *gin.Context) {
	reservationID := c.Param("reservationId")
	if err := rc.ReservationService.CancelReservation(reservationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}

func (rc *ReservationController) GetAllReservations(c *gin.Context) {
	reservations, err := rc.ReservationService.GetAllReservations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}
