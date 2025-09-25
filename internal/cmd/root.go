package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/verse91/fsgo-dev-kit/internal/generator"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fsgo",
	Short: "Interactive generator for full-stack projects with multiple framework options",
	Long: `fsgo is an interactive CLI tool that generates full-stack projects with your choice of:

Backend Frameworks:
  • Fiber (Go)
  • Gin (Go) 
  • Echo (Go)
  • Chi (Go)

Frontend Frameworks:
  • Next.js
  • React
  • Svelte
  • SvelteKit
  • Vue
  • Solid

Project Types:
  • Web (Full-stack with frontend + backend)
  • API (Backend only)

Simply run 'fsgo' and follow the interactive prompts to configure your project.`,
	Run: func(cmd *cobra.Command, args []string) {
		runGenerator()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// runGenerator executes the project generation logic
func runGenerator() {
	projectGen := generator.NewProjectGenerator()
	if err := projectGen.Generate(); err != nil {
		fmt.Printf("Error generating project: %v\n", err)
		os.Exit(1)
	}
}