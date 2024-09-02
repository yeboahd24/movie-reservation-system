# Movie Reservation System

PROJECT_URL: https://roadmap.sh/projects/movie-reservation-system

## Overview

The Movie Reservation System is a web application that allows users to manage movie reservations, including signing up, logging in, and making reservations for showtimes. Admins have additional functionalities for managing movies, showtimes, and user roles.

## Key Functionalities

### 1. User Authentication
- **Signup**: New users can create an account.
- **Login**: Users can log in to access their accounts.
- **Role Management**: Users can be promoted to admin status, allowing for additional privileges.
- **JWT-based Authentication**: Secure access to the API using JSON Web Tokens.

### 2. Movie Management
- **CRUD Operations**: Admins can create, read, update, and delete movie entries.
- **Fetch Movie Details**: Both admins and users can view movie information.

### 3. Showtime Management
- **CRUD Operations**: Admins can manage showtimes for movies.
- **Fetch Showtimes**: Users can view available showtimes for specific movies on given dates.

### 4. Reservation Management
- **Fetch Available Seats**: Users can check available seats for a specific showtime.
- **Reserve Seats**: Users can reserve seats for a showtime.
- **View and Cancel Reservations**: Users can view their upcoming reservations and cancel them if needed.
- **Admin View**: Admins can view all reservations, including capacity and revenue statistics.

### 5. Reporting
- **Generate Reports**: Admins can generate reports on reservations, including total revenue and statistics per movie.

## API Endpoints

The application provides several API endpoints for the functionalities mentioned above. Below are some key endpoints:

- **Signup**: `POST /api/signup`
- **Login**: `POST /api/login`
- **Create Movie**: `POST /api/admin/movies`
- **Add Showtime**: `POST /api/admin/showtimes`
- **Add Reservation**: `POST /api/reservations`
- **Cancel Reservation**: `DELETE /api/reservations/{ID}`
- **Get Available Seats**: `GET /api/showtimes/{SHOWTIME_ID}`
- **Get User Reservations**: `GET /api/user/reservations`
- **Promote User to Admin**: `POST /api/admin/users/{USER_ID}/promote`

## Getting Started

1. **Clone the Repository**: 
   ```bash
   git clone <repository-url>
   ```

2. **Install Dependencies**: 
   ```bash
   cd <project-directory>
   go mod tidy
   ```

3. **Run the Application**: 
   ```bash
   go run cmd/main.go
   ```

4. **Access the API**: Use tools like Postman or cURL to interact with the API endpoints.

## Conclusion

This Movie Reservation System provides a comprehensive solution for managing movie reservations, catering to both users and admins. With its robust authentication and management features, it aims to enhance the movie-going experience.
