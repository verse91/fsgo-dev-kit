package generator

import (
	"fmt"

	"github.com/verse91/fsgo-dev-kit/internal/prompt"
	"github.com/verse91/fsgo-dev-kit/internal/templates"
	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// ProjectGenerator handles the creation of fullstack projects
type ProjectGenerator struct {
	registry *GeneratorRegistry
	prompter *prompt.ProjectPrompt
}

// NewProjectGenerator creates a new project generator
func NewProjectGenerator() *ProjectGenerator {
	return &ProjectGenerator{
		registry: NewGeneratorRegistry(),
		prompter: prompt.NewProjectPrompt(),
	}
}

// Generate creates a complete fullstack project
func (pg *ProjectGenerator) Generate() error {
	// Get project configuration through interactive prompts
	config, err := pg.prompter.GetProjectConfig()
	if err != nil {
		return fmt.Errorf("error getting project configuration: %v", err)
	}

	// Show configuration summary
	pg.prompter.ShowSummary(config)

	// Change to project path if not current directory
	if config.Path != "." {
		if err := utils.CreateDirectory(config.Path, 0o755); err != nil {
			return fmt.Errorf("error creating project directory: %v", err)
		}
		if err := utils.ChangeDir(config.Path); err != nil {
			return fmt.Errorf("error changing to project directory: %v", err)
		}
		defer func() {
			_ = utils.ChangeDir("..") // Go back to original directory
		}()
	}

	// Generate backend
	if err := pg.generateBackend(config); err != nil {
		return fmt.Errorf("error generating backend: %v", err)
	}

	// Generate frontend (only for web projects)
	if config.Type == types.WebProject {
		if err := pg.generateFrontend(config); err != nil {
			return fmt.Errorf("error generating frontend: %v", err)
		}
	}

	// Create root files
	if err := pg.createRootFiles(config); err != nil {
		return fmt.Errorf("error creating root files: %v", err)
	}

	fmt.Printf("✅ Project %s created successfully!\n", config.Name)
	fmt.Println("\nNext steps:")
	if config.Type == types.WebProject {
		fmt.Println("  make run  # Start both frontend and backend")
	} else {
		fmt.Println("  make b    # Start backend server")
	}

	return nil
}

// generateBackend generates the backend using the appropriate generator
func (pg *ProjectGenerator) generateBackend(config *types.ProjectConfig) error {
	generator, exists := pg.registry.GetBackendGenerator(config.BackendFramework)
	if !exists {
		return fmt.Errorf("backend generator not found for framework: %s", config.BackendFramework)
	}
	return generator.Generate(config)
}

// generateFrontend generates the frontend using the appropriate generator
func (pg *ProjectGenerator) generateFrontend(config *types.ProjectConfig) error {
	if config.Frontend == nil {
		return fmt.Errorf("frontend configuration is nil")
	}
	generator, exists := pg.registry.GetFrontendGenerator(config.Frontend.Framework)
	if !exists {
		return fmt.Errorf("frontend generator not found for framework: %s", config.Frontend.Framework)
	}
	return generator.Generate(config)
}

// createRootFiles creates the project root files
func (pg *ProjectGenerator) createRootFiles(config *types.ProjectConfig) error {
	files := map[string]func() string{
		".gitignore": templates.GitignoreTemplate,
		"Makefile":   templates.MakefileTemplate,
		"README.md":  templates.ReadmeTemplate,
	}

	for path, templateFunc := range files {
		if err := utils.CreateFile(path, templateFunc()); err != nil {
			return fmt.Errorf("error creating file %s: %v", path, err)
		}
	}

	return nil
}