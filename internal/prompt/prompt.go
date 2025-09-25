package prompt

import (
	"fmt"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/verse91/fsgo-dev-kit/internal/types"
	"github.com/verse91/fsgo-dev-kit/pkg/utils"
)

// ProjectPrompt handles interactive project configuration
type ProjectPrompt struct{}

// NewProjectPrompt creates a new project prompt
func NewProjectPrompt() *ProjectPrompt {
	return &ProjectPrompt{}
}

// GetProjectConfig prompts user for project configuration
func (p *ProjectPrompt) GetProjectConfig() (*types.ProjectConfig, error) {
	fmt.Println("┌  Creating a new project")
	fmt.Println("│")

	config := &types.ProjectConfig{}
	
	// Get project name/path
	if err := p.promptProjectName(config); err != nil {
		return nil, err
	}

	// Get project type
	if err := p.promptProjectType(config); err != nil {
		return nil, err
	}

	// Get backend framework
	if err := p.promptBackendFramework(config); err != nil {
		return nil, err
	}

	// Get frontend configuration (only for web projects)
	if config.Type == types.WebProject {
		if err := p.promptFrontendConfig(config); err != nil {
			return nil, err
		}
	}

	fmt.Println("└  Configuration complete!")
	fmt.Println()

	return config, nil
}

// promptProjectName prompts for project name/path
func (p *ProjectPrompt) promptProjectName(config *types.ProjectConfig) error {
	currentDir, err := utils.GetCurrentDir()
	if err != nil {
		return err
	}

	projectName := filepath.Base(currentDir)
	
	prompt := &survey.Input{
		Message: "Enter your project name or path (relative to current directory)",
		Default: ".",
		Help:    "Use '.' for current directory or specify a new directory name",
	}

	var result string
	if err := survey.AskOne(prompt, &result); err != nil {
		return err
	}

	if result == "." {
		config.Name = projectName
		config.Path = "."
	} else {
		config.Name = result
		config.Path = result
	}

	fmt.Printf("◇  Project: %s\n", config.Name)
	fmt.Println("│")
	
	return nil
}

// promptProjectType prompts for project type
func (p *ProjectPrompt) promptProjectType(config *types.ProjectConfig) error {
	projectTypes := types.GetProjectTypes()
	options := make([]string, len(projectTypes))
	for i, pt := range projectTypes {
		options[i] = string(pt)
	}

	prompt := &survey.Select{
		Message: "Select project type",
		Options: options,
		Help:    "Web: Full-stack with frontend + backend, API: Backend only",
	}

	var result string
	if err := survey.AskOne(prompt, &result); err != nil {
		return err
	}

	config.Type = types.ProjectType(result)
	fmt.Printf("◇  Type: %s\n", result)
	fmt.Println("│")

	return nil
}

// promptBackendFramework prompts for backend framework
func (p *ProjectPrompt) promptBackendFramework(config *types.ProjectConfig) error {
	frameworks := types.GetBackendFrameworks()
	options := make([]string, len(frameworks))
	for i, fw := range frameworks {
		options[i] = string(fw)
	}

	prompt := &survey.Select{
		Message: "Choose backend framework",
		Options: options,
		Help:    "Select the Go web framework for your backend",
	}

	var result string
	if err := survey.AskOne(prompt, &result); err != nil {
		return err
	}

	config.BackendFramework = types.BackendFramework(result)
	fmt.Printf("◇  Backend: %s\n", result)
	fmt.Println("│")

	return nil
}

// promptFrontendConfig prompts for frontend configuration
func (p *ProjectPrompt) promptFrontendConfig(config *types.ProjectConfig) error {
	config.Frontend = &types.FrontendConfig{}

	// Get framework
	if err := p.promptFrontendFramework(config.Frontend); err != nil {
		return err
	}

	// Get additional options
	if err := p.promptFrontendOptions(config.Frontend); err != nil {
		return err
	}

	return nil
}

// promptFrontendFramework prompts for frontend framework
func (p *ProjectPrompt) promptFrontendFramework(frontend *types.FrontendConfig) error {
	frameworks := types.GetFrontendFrameworks()
	options := make([]string, len(frameworks))
	for i, fw := range frameworks {
		options[i] = string(fw)
	}

	prompt := &survey.Select{
		Message: "Choose frontend framework",
		Options: options,
		Help:    "Select the frontend framework/library",
	}

	var result string
	if err := survey.AskOne(prompt, &result); err != nil {
		return err
	}

	frontend.Framework = types.FrontendFramework(result)
	fmt.Printf("◇  Frontend: %s\n", result)
	fmt.Println("│")

	return nil
}

// promptFrontendOptions prompts for frontend configuration options
func (p *ProjectPrompt) promptFrontendOptions(frontend *types.FrontendConfig) error {
	// TypeScript
	if err := p.promptBoolOption("Use TypeScript?", &frontend.TypeScript, true); err != nil {
		return err
	}

	// TailwindCSS
	if err := p.promptBoolOption("Use Tailwind CSS?", &frontend.TailwindCSS, true); err != nil {
		return err
	}

	// ESLint
	if err := p.promptBoolOption("Use ESLint?", &frontend.ESLint, true); err != nil {
		return err
	}

	return nil
}

// promptBoolOption prompts for a boolean configuration option
func (p *ProjectPrompt) promptBoolOption(message string, target *bool, defaultValue bool) error {
	prompt := &survey.Confirm{
		Message: message,
		Default: defaultValue,
	}

	if err := survey.AskOne(prompt, target); err != nil {
		return err
	}

	status := "No"
	if *target {
		status = "Yes"
	}
	
	fmt.Printf("◇  %s %s\n", message, status)
	fmt.Println("│")

	return nil
}

// ShowSummary displays the project configuration summary
func (p *ProjectPrompt) ShowSummary(config *types.ProjectConfig) {
	fmt.Println("┌  Project Summary")
	fmt.Println("│")
	fmt.Printf("│  Name: %s\n", config.Name)
	fmt.Printf("│  Path: %s\n", config.Path)
	fmt.Printf("│  Type: %s\n", config.Type)
	fmt.Printf("│  Backend: %s\n", config.BackendFramework)
	
	if config.Frontend != nil {
		fmt.Printf("│  Frontend: %s\n", config.Frontend.Framework)
		fmt.Printf("│  TypeScript: %t\n", config.Frontend.TypeScript)
		fmt.Printf("│  Tailwind CSS: %t\n", config.Frontend.TailwindCSS)
		fmt.Printf("│  ESLint: %t\n", config.Frontend.ESLint)
	}
	
	fmt.Println("│")
	fmt.Println("└  Ready to generate!")
	fmt.Println()
}