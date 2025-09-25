package frontend

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// SvelteGenerator handles Svelte frontend generation
type SvelteGenerator struct{}

// NewSvelteGenerator creates a new Svelte frontend generator
func NewSvelteGenerator() *SvelteGenerator {
	return &SvelteGenerator{}
}

// GetFramework returns the framework name
func (g *SvelteGenerator) GetFramework() types.FrontendFramework {
	return types.Svelte
}

// GetBuildCommands returns the build commands for Svelte
func (g *SvelteGenerator) GetBuildCommands() []string {
	return []string{"npm run build", "npm run dev"}
}

// Generate creates a new Svelte frontend project
func (g *SvelteGenerator) Generate(config *types.ProjectConfig) error {
	fmt.Println("🚀 Creating Svelte frontend...")

	// Build create command based on configuration
	cmdArgs := g.buildCreateCommand(config.Frontend)
	cmd := exec.Command("sh", "-c", cmdArgs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error creating Svelte project: %v", err)
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

// buildCreateCommand builds the Svelte create command based on configuration
func (g *SvelteGenerator) buildCreateCommand(frontend *types.FrontendConfig) string {
	baseCmd := "npm create svelte@latest client"
	
	// Svelte create process is interactive, so we'll use defaults
	// In a real implementation, you might want to handle this differently
	return baseCmd
}

// createEnvFiles creates the environment files
func (g *SvelteGenerator) createEnvFiles() error {
	envContent := "VITE_API_URL=http://localhost:8080\n"
	envExampleContent := "VITE_API_URL=http://localhost:8080\n"

	if err := utils.CreateFile(".env", envContent); err != nil {
		return fmt.Errorf("error creating .env: %v", err)
	}
	if err := utils.CreateFile(".env.example", envExampleContent); err != nil {
		return fmt.Errorf("error creating .env.example: %v", err)
	}

	return nil
}