package types

// ProjectType represents the type of project to create
type ProjectType string

const (
	WebProject ProjectType = "Web"
	APIProject ProjectType = "API"
)

// BackendFramework represents available backend frameworks
type BackendFramework string

const (
	Fiber BackendFramework = "Fiber"
	Gin   BackendFramework = "Gin"
	Echo  BackendFramework = "Echo"
	Chi   BackendFramework = "Chi"
)

// FrontendFramework represents available frontend frameworks
type FrontendFramework string

const (
	NextJS     FrontendFramework = "Next.js"
	React      FrontendFramework = "React"
	Vue        FrontendFramework = "Vue"
	Svelte     FrontendFramework = "Svelte"
	SvelteKit  FrontendFramework = "SvelteKit"
	Solid      FrontendFramework = "Solid"
)

// ProjectConfig holds the configuration for generating a project
type ProjectConfig struct {
	Name             string
	Path             string
	Type             ProjectType
	BackendFramework BackendFramework
	Frontend         *FrontendConfig // nil for API-only projects
}

// FrontendConfig holds frontend-specific configuration
type FrontendConfig struct {
	Framework FrontendFramework
	TypeScript bool
	TailwindCSS bool
	ESLint     bool
}

// GetBackendFrameworks returns available backend frameworks
func GetBackendFrameworks() []BackendFramework {
	return []BackendFramework{Fiber, Gin, Echo, Chi}
}

// GetFrontendFrameworks returns available frontend frameworks
func GetFrontendFrameworks() []FrontendFramework {
	return []FrontendFramework{NextJS, React, Vue, Svelte, SvelteKit, Solid}
}

// GetProjectTypes returns available project types
func GetProjectTypes() []ProjectType {
	return []ProjectType{WebProject, APIProject}
}