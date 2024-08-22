package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yeboahd24/movie-reservation-system/models"
	"github.com/yeboahd24/movie-reservation-system/services"
)

type ShowtimeController struct {
	showtimeService *services.ShowtimeService
}

func NewShowtimeController(showtimeService *services.ShowtimeService) *ShowtimeController {
	return &ShowtimeController{
		showtimeService: showtimeService,
	}
}

func (sc *ShowtimeController) CreateShowtime(c *gin.Context) {
	var showtime models.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdShowtime, err := sc.showtimeService.CreateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdShowtime)
}

func (sc *ShowtimeController) GetShowtimes(c *gin.Context) {
	movieID := c.Param("movieID")
	showtimes, err := sc.showtimeService.GetShowtimes(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, showtimes)
}

func (sc *ShowtimeController) UpdateShowtime(c *gin.Context) {
	var showtime models.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedShowtime, err := sc.showtimeService.UpdateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedShowtime)
}

func (sc *ShowtimeController) DeleteShowtime(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	err := sc.showtimeService.DeleteShowtime(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Showtime deleted successfully"})
}

func (sc *ShowtimeController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	availableSeats, err := sc.showtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available_seats": availableSeats})
}
