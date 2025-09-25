package backend

import (
	"fmt"

	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// GinGenerator handles Gin backend generation
type GinGenerator struct{}

// NewGinGenerator creates a new Gin backend generator
func NewGinGenerator() *GinGenerator {
	return &GinGenerator{}
}

// GetFramework returns the framework name
func (g *GinGenerator) GetFramework() types.BackendFramework {
	return types.Gin
}

// GetDependencies returns the list of dependencies for Gin
func (g *GinGenerator) GetDependencies() []string {
	return []string{
		"github.com/gin-gonic/gin",
		"github.com/joho/godotenv",
		"go.uber.org/zap",
		"github.com/gin-contrib/cors",
	}
}

// Generate creates a new Gin backend project
func (g *GinGenerator) Generate(config *types.ProjectConfig) error {
	fmt.Println("🚀 Creating Gin backend...")

	// Create server directory
	serverDir := "server"
	if err := utils.CreateDirectory(serverDir, 0o755); err != nil {
		return fmt.Errorf("error creating server directory: %v", err)
	}

	// Change to server directory
	if err := utils.ChangeDir(serverDir); err != nil {
		return fmt.Errorf("error changing to server directory: %v", err)
	}

	// Initialize Go module
	if err := utils.RunCommand("go", "mod", "init", "server"); err != nil {
		fmt.Printf("Error initializing Go module: %v\n", err)
	}

	// Install dependencies
	if err := g.installDependencies(); err != nil {
		return fmt.Errorf("error installing dependencies: %v", err)
	}

	// Create basic structure and files
	if err := g.createBasicStructure(); err != nil {
		return fmt.Errorf("error creating basic structure: %v", err)
	}

	// Go back to project root
	if err := utils.ChangeDir(".."); err != nil {
		return fmt.Errorf("error returning to project root: %v", err)
	}

	return nil
}

// installDependencies installs Go dependencies
func (g *GinGenerator) installDependencies() error {
	deps := g.GetDependencies()

	for _, dep := range deps {
		if err := utils.RunCommand("go", "get", dep); err != nil {
			fmt.Printf("Error installing dependency %s: %v\n", dep, err)
		}
	}

	return nil
}

// createBasicStructure creates basic structure for Gin
func (g *GinGenerator) createBasicStructure() error {
	// Create basic main.go for Gin
	mainGoContent := `package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Create Gin router
	r := gin.Default()

	// CORS middleware
	r.Use(cors.Default())

	// Health check endpoint
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(r.Run(":" + port))
}
`

	if err := utils.CreateFile("main.go", mainGoContent); err != nil {
		return fmt.Errorf("error creating main.go: %v", err)
	}

	// Create .env files
	envContent := "PORT=8080\nENV=development\n"
	if err := utils.CreateFile(".env", envContent); err != nil {
		return fmt.Errorf("error creating .env: %v", err)
	}
	if err := utils.CreateFile(".env.example", envContent); err != nil {
		return fmt.Errorf("error creating .env.example: %v", err)
	}

	return nil
}