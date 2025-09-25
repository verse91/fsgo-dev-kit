package backend

import (
	"fmt"

	"github.com/verse91/fsgo-dev-kit/internal/templates"
	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// FiberGenerator handles Go Fiber backend generation
type FiberGenerator struct{}

// NewFiberGenerator creates a new Fiber backend generator
func NewFiberGenerator() *FiberGenerator {
	return &FiberGenerator{}
}

// Generate creates a new Go Fiber backend project
func (g *FiberGenerator) Generate(config *types.ProjectConfig) error {
	fmt.Println("🚀 Creating Go Fiber backend...")

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

	// Create backend structure
	if err := g.createDirectoryStructure(); err != nil {
		return fmt.Errorf("error creating directory structure: %v", err)
	}

	// Create Go files
	if err := g.createGoFiles(); err != nil {
		return fmt.Errorf("error creating Go files: %v", err)
	}

	// Create configuration files
	if err := g.createConfigFiles(); err != nil {
		return fmt.Errorf("error creating config files: %v", err)
	}

	// Create environment files
	if err := g.createEnvFiles(); err != nil {
		return fmt.Errorf("error creating environment files: %v", err)
	}

	// Go back to project root
	if err := utils.ChangeDir(".."); err != nil {
		return fmt.Errorf("error returning to project root: %v", err)
	}

	return nil
}

// GetFramework returns the framework name
func (g *FiberGenerator) GetFramework() types.BackendFramework {
	return types.Fiber
}

// GetDependencies returns the list of dependencies
func (g *FiberGenerator) GetDependencies() []string {
	return []string{
		"github.com/gofiber/fiber/v3",
		"github.com/joho/godotenv",
		"go.uber.org/zap",
		"github.com/gofiber/helmet/v2",
		"github.com/gofiber/cors",
	}
}

// installDependencies installs Go dependencies
func (g *FiberGenerator) installDependencies() error {
	deps := g.GetDependencies()

	for _, dep := range deps {
		if err := utils.RunCommand("go", "get", dep); err != nil {
			fmt.Printf("Error installing dependency %s: %v\n", dep, err)
		}
	}

	return nil
}

// createDirectoryStructure creates the backend directory structure
func (g *FiberGenerator) createDirectoryStructure() error {
	dirs := []string{
		"api",
		"cmd/server/tmp",
		"cmd/test",
		"db/migrations",
		"internal/config",
		"internal/controller",
		"internal/middleware",
		"internal/model",
		"internal/repo",
		"internal/routes",
		"internal/service",
		"pkg/logger",
		"pkg/response",
		"pkg/utils",
		"tmp",
	}

	for _, dir := range dirs {
		if err := utils.CreateDirectory(dir, 0o755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// createGoFiles creates all the Go source files
func (g *FiberGenerator) createGoFiles() error {
	files := map[string]func() string{
		"cmd/server/main.go":                  templates.MainGoFile,
		"api/api.go":                          templates.ApiGoFile,
		"internal/config/config.go":           templates.ConfigGoFile,
		"internal/routes/routes.go":           templates.RoutesGoFile,
		"pkg/logger/zap.go":                   templates.LoggerGoFile,
		"pkg/utils/env.go":                    templates.EnvUtilsGoFile,
		"pkg/response/response.go":            templates.ResponseGoFile,
		"pkg/response/httpStatusCode.go":      templates.HttpStatusCodeGoFile,
		"internal/middleware/rate-limit.go":   templates.RateLimitMiddleware,
		"internal/middleware/user-api-key.go": templates.ApiKeyMiddleware,
		"cmd/test/db.go":                      templates.TestDbGoFile,
		"db/connect.go":                       templates.DbConnectGoFile,
		"db/migrations/migrate.go":            templates.MigrationsGoFile,
		"db/migrations/schema.sql":            templates.SchemaSQLFile,
		"Dockerfile":                          templates.DockerfileTemplate,
	}

	for path, templateFunc := range files {
		if err := utils.CreateFile(path, templateFunc()); err != nil {
			return fmt.Errorf("error creating file %s: %v", path, err)
		}
	}

	return nil
}

// createConfigFiles creates configuration files
func (g *FiberGenerator) createConfigFiles() error {
	// Create .air.toml
	if err := utils.CreateFile(".air.toml", templates.AirConfigTemplate()); err != nil {
		return fmt.Errorf("error creating .air.toml: %v", err)
	}

	return nil
}

// createEnvFiles creates the environment files
func (g *FiberGenerator) createEnvFiles() error {
	files := map[string]func() string{
		".env":         templates.BackendEnvFile,
		".env.example": templates.BackendEnvExampleFile,
	}

	for path, templateFunc := range files {
		if err := utils.CreateFile(path, templateFunc()); err != nil {
			return fmt.Errorf("error creating file %s: %v", path, err)
		}
	}

	return nil
}

