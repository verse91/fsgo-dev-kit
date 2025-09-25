package frontend

import (
	"fmt"
	"os/exec"
	"os"

	"github.com/verse91/fsgo-dev-kit/internal/templates"
	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// NextJSGenerator handles Next.js frontend generation
type NextJSGenerator struct {}

// NewNextJSGenerator creates a new Next.js frontend generator
func NewNextJSGenerator() *NextJSGenerator {
	return &NextJSGenerator{}
}

// GetFramework returns the framework name
func (g *NextJSGenerator) GetFramework() types.FrontendFramework {
	return types.NextJS
}

// GetBuildCommands returns the build commands for Next.js
func (g *NextJSGenerator) GetBuildCommands() []string {
	return []string{"bun run build", "bun run start"}
}

// Generate creates a new Next.js frontend project
func (g *NextJSGenerator) Generate(config *types.ProjectConfig) error {
	fmt.Println("🚀 Creating Next.js frontend...")

	// Build create command based on configuration
	cmdArgs := g.buildCreateCommand(config.Frontend)
	cmd := exec.Command("sh", "-c", cmdArgs)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error creating Next.js project: %v", err)
	}

	// Change to client directory
	if err := utils.ChangeDir("client"); err != nil {
		return fmt.Errorf("error changing to client directory: %v", err)
	}

	// Create additional directory structure
	if err := g.createDirectoryStructure(); err != nil {
		return fmt.Errorf("error creating directory structure: %v", err)
	}

	// Create component files
	if err := g.createComponents(); err != nil {
		return fmt.Errorf("error creating components: %v", err)
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

// buildCreateCommand builds the Next.js create command based on configuration
func (g *NextJSGenerator) buildCreateCommand(frontend *types.FrontendConfig) string {
	baseCmd := "yes \"\" | bun create next-app@latest client"
	
	if frontend.TypeScript {
		baseCmd += " --typescript"
	} else {
		baseCmd += " --js"
	}
	
	if frontend.ESLint {
		baseCmd += " --eslint"
	}
	
	if frontend.TailwindCSS {
		baseCmd += " --tailwind"
	}
	
	baseCmd += " --app" // Always use App Router
	return baseCmd
}

// createDirectoryStructure creates the frontend directory structure
func (g *NextJSGenerator) createDirectoryStructure() error {
	dirs := []string{
		"public/assets/fonts/components-fonts",
		"public/assets/fonts/logo-font", 
		"public/assets/icons",
		"src/app/auth/callback",
		"src/app/auth/signin",
		"src/components/homepage",
		"src/components/ui/navbar",
		"src/components/ui/texts",
		"src/lib",
		"src/styles",
	}

	for _, dir := range dirs {
		if err := utils.CreateDirectory(dir, 0o755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// createComponents creates the component files
func (g *NextJSGenerator) createComponents() error {
	files := map[string]func() string{
		"src/components/homepage/Hero.tsx":         templates.HeroComponent,
		"src/components/ui/navbar/Navbar.tsx":     templates.NavbarComponent,
		"src/components/ui/texts/Typography.tsx":  templates.TypographyComponent,
		"src/app/auth/signin/page.tsx":            templates.SignInPage,
		"src/app/auth/callback/page.tsx":          templates.AuthCallbackPage,
	}

	for path, templateFunc := range files {
		if err := utils.CreateFile(path, templateFunc()); err != nil {
			return fmt.Errorf("error creating file %s: %v", path, err)
		}
	}

	return nil
}

// createEnvFiles creates the environment files
func (g *NextJSGenerator) createEnvFiles() error {
	files := map[string]func() string{
		".env":         templates.FrontendEnvFile,
		".env.example": templates.FrontendEnvExampleFile,
	}

	for path, templateFunc := range files {
		if err := utils.CreateFile(path, templateFunc()); err != nil {
			return fmt.Errorf("error creating file %s: %v", path, err)
		}
	}

	return nil
}