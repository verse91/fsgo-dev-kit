package templates

// MainGoFile returns the main.go template for the server
func MainGoFile() string {
	return `package main

import (
	"log"
	"server/internal/config"
	"server/internal/routes"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/cors"
	"github.com/gofiber/helmet/v2"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize logger
	logger.Init()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Your API Server",
	})

	// Middleware
	app.Use(helmet.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// Routes
	routes.Setup(app)

	// Start server
	log.Fatal(app.Listen(":" + cfg.Port))
}`
}

// ApiGoFile returns the api.go template
func ApiGoFile() string {
	return `package api

import (
	"github.com/gofiber/fiber/v3"
)

func HealthCheck(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "ok",
		"message": "Server is running",
	})
}`
}

// ConfigGoFile returns the config.go template
func ConfigGoFile() string {
	return `package config

import (
	"os"
	"server/pkg/utils"
)

type Config struct {
	Port  string
	Env   string
	DbURL string
}

func Load() *Config {
	utils.LoadEnv()

	return &Config{
		Port:  getEnv("PORT", "8080"),
		Env:   getEnv("ENV", "development"),
		DbURL: getEnv("DB_URL", ""),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}`
}

// RoutesGoFile returns the routes.go template
func RoutesGoFile() string {
	return `package routes

import (
	"server/api"

	"github.com/gofiber/fiber/v3"
)

func Setup(app *fiber.App) {
	// API routes
	apiRoutes := app.Group("/api/v1")
	
	// Health check
	apiRoutes.Get("/health", api.HealthCheck)
}`
}

// LoggerGoFile returns the logger.go template
func LoggerGoFile() string {
	return `package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}`
}

// EnvUtilsGoFile returns the env utils template
func EnvUtilsGoFile() string {
	return `package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}`
}

// ResponseGoFile returns the response utils template
func ResponseGoFile() string {
	return `package response

import "github.com/gofiber/fiber/v3"

type Response struct {
	Success bool        ` + "`json:\"success\"`" + `
	Message string      ` + "`json:\"message\"`" + `
	Data    interface{} ` + "`json:\"data,omitempty\"`" + `
	Error   interface{} ` + "`json:\"error,omitempty\"`" + `
}

func Success(c fiber.Ctx, data interface{}, message string) error {
	return c.JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c fiber.Ctx, statusCode int, message string, err interface{}) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}`
}

// HttpStatusCodeGoFile returns the HTTP status codes template
func HttpStatusCodeGoFile() string {
	return `package response

const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
)`
}

// RateLimitMiddleware returns the rate limit middleware template
func RateLimitMiddleware() string {
	return `package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
)

func RateLimit() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        100,
		Expiration: time.Minute,
		KeyGenerator: func(c fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c fiber.Ctx) error {
			return c.Status(429).JSON(fiber.Map{
				"error": "Rate limit exceeded",
			})
		},
	})
}`
}

// ApiKeyMiddleware returns the API key middleware template
func ApiKeyMiddleware() string {
	return `package middleware

import (
	"server/pkg/response"

	"github.com/gofiber/fiber/v3"
)

func ValidateAPIKey() fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := c.Get("X-API-Key")
		
		if apiKey == "" {
			return response.Error(c, response.StatusUnauthorized, "API Key is required", nil)
		}

		// Add your API key validation logic here
		
		return c.Next()
	}
}`
}

// TestDbGoFile returns the test database file template
func TestDbGoFile() string {
	return `package main

import (
	"fmt"
	"server/internal/config"
)

func main() {
	cfg := config.Load()
	fmt.Printf("Testing database connection: %s\n", cfg.DbURL)
	
	// Add your database connection test logic here
	fmt.Println("Database test completed")
}`
}

// DbConnectGoFile returns the database connection template
func DbConnectGoFile() string {
	return `package db

import (
	"database/sql"
	"server/internal/config"
)

var DB *sql.DB

func Connect() error {
	cfg := config.Load()
	
	// Add your database connection logic here
	// Example: DB, err := sql.Open("postgres", cfg.DbURL)
	
	return nil
}`
}

// MigrationsGoFile returns the migrations template
func MigrationsGoFile() string {
	return `package migrations

func RunMigrations() error {
	// Add your migration logic here
	return nil
}`
}

// SchemaSQLFile returns the schema SQL template
func SchemaSQLFile() string {
	return `-- Add your database schema here
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
}

// DockerfileTemplate returns the Dockerfile template
func DockerfileTemplate() string {
	return `FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]`
}

// AirConfigTemplate returns the .air.toml template
func AirConfigTemplate() string {
	return `root = "."
tmp_dir = "tmp"

[build]
    cmd = "go build -o ./cmd/server/tmp/main.exe ./cmd/server/main.go"
    bin = "./cmd/server/tmp/main.exe"
    include_ext = ["go", "yaml", "yml", "env"]
    exclude_dir = ["assets", "tmp", "vendor", "client/node_modules"]
    exclude_file = []
    delay = 200
    send_interrupt = true
    stop_on_error = true
    log = "build.log"

[log]
    time = true

[color]
    main = "magenta"
    watcher = "cyan"
    build = "yellow"
    runner = "green"`
}

// BackendEnvFile returns the backend environment file template
func BackendEnvFile() string {
	return "PORT=8080\nENV=development\nDB_URL=postgresql://user:password@localhost:5432/dbname\n"
}

// BackendEnvExampleFile returns the backend environment example file template
func BackendEnvExampleFile() string {
	return "PORT=8080\nENV=development\nDB_URL=postgresql://user:password@localhost:5432/dbname\n"
}
