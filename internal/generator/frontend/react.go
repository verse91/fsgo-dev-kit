package frontend

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// ReactGenerator handles React frontend generation
type ReactGenerator struct{}

// NewReactGenerator creates a new React frontend generator
func NewReactGenerator() *ReactGenerator {
	return &ReactGenerator{}
}

// GetFramework returns the framework name
func (g *ReactGenerator) GetFramework() types.FrontendFramework {
	return types.React
}

// GetBuildCommands returns the build commands for React
func (g *ReactGenerator) GetBuildCommands() []string {
	return []string{"npm run build", "npm start"}
}

// Generate creates a new React frontend project
func (g *ReactGenerator) Generate(config *types.ProjectConfig) error {
	fmt.Println("🚀 Creating React frontend...")

	// Build create command based on configuration
	cmdArgs := g.buildCreateCommand(config.Frontend)
	cmd := exec.Command("sh", "-c", cmdArgs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error creating React project: %v", err)
	}

	// Change to client directory
	if err := utils.ChangeDir("client"); err != nil {
		return fmt.Errorf("error changing to client directory: %v", err)
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

// buildCreateCommand builds the React create command based on configuration
func (g *ReactGenerator) buildCreateCommand(frontend *types.FrontendConfig) string {
	baseCmd := "npx create-react-app client"
	
	if frontend.TypeScript {
		baseCmd += " --template typescript"
	}
	
	return baseCmd
}

// createEnvFiles creates the environment files
func (g *ReactGenerator) createEnvFiles() error {
	envContent := "REACT_APP_API_URL=http://localhost:8080\n"
	envExampleContent := "REACT_APP_API_URL=http://localhost:8080\n"

	if err := utils.CreateFile(".env", envContent); err != nil {
		return fmt.Errorf("error creating .env: %v", err)
	}
	if err := utils.CreateFile(".env.example", envExampleContent); err != nil {
		return fmt.Errorf("error creating .env.example: %v", err)
	}

	return nil
}