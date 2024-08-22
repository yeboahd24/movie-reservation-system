package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yeboahd24/movie-reservation-system/controllers"
	"github.com/yeboahd24/movie-reservation-system/services"
	"github.com/yeboahd24/movie-reservation-system/utils"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Initialize services
	authService := services.NewAuthService(db)
	movieService := services.NewMovieService(db)
	reservationService := services.NewReservationService(db)
	showtimeService := services.NewShowtimeService(db)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)
	reservationController := controllers.NewReservationController(reservationService, showtimeService)
	showtimeController := controllers.NewShowtimeController(showtimeService)

	// Public routes
	public := router.Group("/api")
	{
		public.POST("/signup", authController.SignUp)
		public.POST("/login", authController.Login)
		public.GET("/movies", movieController.GetMovies)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/user/reservations", reservationController.GetUserReservations)
		protected.POST("/reservations", reservationController.CreateReservation)
		protected.DELETE("/reservations/:reservationId", reservationController.CancelReservation)
		protected.GET("/showtimes/:showtimeId/seats", reservationController.GetAvailableSeats)
		protected.GET("/movies/:movieID/showtimes", showtimeController.GetShowtimes)
	}

	// Admin routes
	admin := router.Group("/api/admin")
	admin.Use(utils.AuthMiddleware(), utils.AdminMiddleware())
	// admin.Use(utils.AuthMiddleware())

	{
		admin.POST("/movies", movieController.CreateMovie)
		admin.PUT("/movies/:movieId", movieController.UpdateMovie)
		admin.DELETE("/movies/:movieId", movieController.DeleteMovie)
		admin.GET("/reservations", reservationController.GetAllReservations)
		admin.POST("/users/:userId/promote", authController.PromoteToAdmin)
		admin.POST("/showtimes", showtimeController.CreateShowtime)
		admin.PUT("/showtimes/:showtimeId", showtimeController.UpdateShowtime)
		admin.DELETE("/showtimes/:showtimeId", showtimeController.DeleteShowtime)
	}

	return router
}
