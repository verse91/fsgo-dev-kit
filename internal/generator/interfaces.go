package generator

import (
	"github.com/verse91/fsgo-dev-kit/internal/generator/backend"
	"github.com/verse91/fsgo-dev-kit/internal/generator/frontend"
	"github.com/verse91/fsgo-dev-kit/internal/types"
)

// BackendGenerator interface for backend framework generators
type BackendGenerator interface {
	Generate(config *types.ProjectConfig) error
	GetFramework() types.BackendFramework
	GetDependencies() []string
}

// FrontendGenerator interface for frontend framework generators
type FrontendGenerator interface {
	Generate(config *types.ProjectConfig) error
	GetFramework() types.FrontendFramework
	GetBuildCommands() []string
}

// GeneratorRegistry manages available generators
type GeneratorRegistry struct {
	backendGenerators  map[types.BackendFramework]BackendGenerator
	frontendGenerators map[types.FrontendFramework]FrontendGenerator
}

// NewGeneratorRegistry creates a new registry with all available generators
func NewGeneratorRegistry() *GeneratorRegistry {
	registry := &GeneratorRegistry{
		backendGenerators:  make(map[types.BackendFramework]BackendGenerator),
		frontendGenerators: make(map[types.FrontendFramework]FrontendGenerator),
	}
	
	// Register backend generators
	registry.RegisterBackendGenerator(backend.NewFiberGenerator())
	registry.RegisterBackendGenerator(backend.NewGinGenerator())
	registry.RegisterBackendGenerator(backend.NewEchoGenerator())
	
	// Register frontend generators
	registry.RegisterFrontendGenerator(frontend.NewNextJSGenerator())
	registry.RegisterFrontendGenerator(frontend.NewReactGenerator())
	registry.RegisterFrontendGenerator(frontend.NewSvelteGenerator())
	
	return registry
}

// RegisterBackendGenerator registers a backend generator
func (r *GeneratorRegistry) RegisterBackendGenerator(gen BackendGenerator) {
	r.backendGenerators[gen.GetFramework()] = gen
}

// RegisterFrontendGenerator registers a frontend generator
func (r *GeneratorRegistry) RegisterFrontendGenerator(gen FrontendGenerator) {
	r.frontendGenerators[gen.GetFramework()] = gen
}

// GetBackendGenerator returns a backend generator for the specified framework
func (r *GeneratorRegistry) GetBackendGenerator(framework types.BackendFramework) (BackendGenerator, bool) {
	gen, exists := r.backendGenerators[framework]
	return gen, exists
}

// GetFrontendGenerator returns a frontend generator for the specified framework
func (r *GeneratorRegistry) GetFrontendGenerator(framework types.FrontendFramework) (FrontendGenerator, bool) {
	gen, exists := r.frontendGenerators[framework]
	return gen, exists
}

// GetAvailableBackendFrameworks returns list of registered backend frameworks
func (r *GeneratorRegistry) GetAvailableBackendFrameworks() []types.BackendFramework {
	frameworks := make([]types.BackendFramework, 0, len(r.backendGenerators))
	for framework := range r.backendGenerators {
		frameworks = append(frameworks, framework)
	}
	return frameworks
}

// GetAvailableFrontendFrameworks returns list of registered frontend frameworks
func (r *GeneratorRegistry) GetAvailableFrontendFrameworks() []types.FrontendFramework {
	frameworks := make([]types.FrontendFramework, 0, len(r.frontendGenerators))
	for framework := range r.frontendGenerators {
		frameworks = append(frameworks, framework)
	}
	return frameworks
}