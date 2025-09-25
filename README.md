# fsgo - Interactive Full-Stack Project Generator

An interactive CLI tool for generating modern full-stack projects with your choice of popular frameworks.

## ✨ Features

### 🎯 Interactive Configuration
- Beautiful CLI prompts with arrow key navigation
- Real-time validation and help text
- Project type selection (Web vs API)
- Framework-specific options (TypeScript, TailwindCSS, ESLint)

### 🔧 Multiple Backend Frameworks
- **Fiber** - Express-inspired web framework for Go
- **Gin** - HTTP web framework with high performance
- **Echo** - High performance, extensible web framework
- **Chi** - Lightweight, idiomatic router for building HTTP services

### 🎨 Multiple Frontend Frameworks
- **Next.js** - React framework with SSR/SSG capabilities
- **React** - Popular JavaScript library for building user interfaces
- **Svelte** - Cybernetically enhanced web apps
- **SvelteKit** - Full-stack framework powered by Svelte
- **Vue** - Progressive JavaScript framework
- **Solid** - Simple and performant reactivity for building user interfaces

### 📋 Project Types
- **Web Projects** - Full-stack with both frontend and backend
- **API Projects** - Backend-only for microservices or APIs

## 🚀 Quick Start

### Installation

```bash
# Build from source
git clone https://github.com/verse91/fsgo-dev-kit
cd fsgo-dev-kit
make cli
```

### Usage

Simply run the command and follow the interactive prompts:

```bash
fsgo
```

The tool will guide you through:

1. **Project Name/Path** - Specify where to create your project
2. **Project Type** - Choose between Web or API
3. **Backend Framework** - Select your preferred Go framework
4. **Frontend Framework** - Choose your frontend stack (Web projects only)
5. **Configuration Options** - TypeScript, TailwindCSS, ESLint, etc.

### Example Interactive Flow

```
┌  Creating a new project
│
◇  Enter your project name or path (relative to current directory)
│  my-awesome-app
│
◇  Select project type
│  Web
│
◇  Choose backend framework
│  Gin
│
◇  Choose frontend framework
│  React
│
◇  Use TypeScript? Yes
│
◇  Use Tailwind CSS? Yes  
│
◇  Use ESLint? Yes
│
└  Configuration complete!

┌  Project Summary
│
│  Name: my-awesome-app
│  Path: my-awesome-app
│  Type: Web
│  Backend: Gin
│  Frontend: React
│  TypeScript: true
│  Tailwind CSS: true
│  ESLint: true
│
└  Ready to generate!

🚀 Creating Gin backend...
🚀 Creating React frontend...
✅ Project my-awesome-app created successfully!

Next steps:
  make run  # Start both frontend and backend
```

## 📁 Generated Project Structure

### Web Projects
```
my-project/
├── client/          # Frontend application
│   ├── src/         # Source code
│   ├── public/      # Static assets
│   └── package.json # Dependencies
├── server/          # Backend application
│   ├── main.go      # Entry point
│   ├── go.mod       # Go dependencies
│   └── .env         # Environment variables
├── Makefile         # Build and run commands
├── .gitignore       # Git ignore rules
└── README.md        # Project documentation
```

### API Projects
```
my-api/
├── server/          # Backend application
│   ├── main.go      # Entry point
│   ├── go.mod       # Go dependencies
│   └── .env         # Environment variables
├── Makefile         # Build and run commands
├── .gitignore       # Git ignore rules
└── README.md        # Project documentation
```

## 🔧 Available Commands

After generating a project, you can use these commands:

```bash
# Start both frontend and backend (Web projects)
make run

# Start backend only
make b

# Start frontend only (Web projects) 
make f

# Test database connection
make testdb

# Stop all processes
make stop
```

## 🏗️ Framework Details

### Backend Frameworks

#### Fiber
- Ultra-fast HTTP framework inspired by Express
- Built-in middleware for CORS, logging, rate limiting
- Comprehensive project structure with clean architecture

#### Gin  
- High-performance HTTP web framework
- Minimal memory footprint
- Simple setup with essential middleware

#### Echo
- High-performance, extensible web framework
- Built-in middleware and JSON binding
- Clean and minimal implementation

### Frontend Frameworks

#### Next.js
- Full-featured React framework
- App Router with TypeScript support
- Pre-configured with TailwindCSS and ESLint
- Comprehensive component structure

#### React
- Standard Create React App setup
- TypeScript template option
- Environment configuration for API integration

#### Svelte
- Modern reactive framework
- Fast build times and small bundles
- Vite-powered development server

## 🔧 Development

### Building from Source

```bash
git clone https://github.com/verse91/fsgo-dev-kit
cd fsgo-dev-kit
go mod tidy
go build -o fsgo .
```

### Running Tests

```bash
make test
```

### Project Architecture

The project follows clean architecture principles:

- `cmd/fsgo/` - CLI entry point
- `internal/cmd/` - Cobra CLI configuration
- `internal/generator/` - Core generation logic
  - `backend/` - Backend framework generators
  - `frontend/` - Frontend framework generators
- `internal/prompt/` - Interactive CLI prompts
- `internal/templates/` - File templates
- `internal/types/` - Type definitions
- `pkg/utils/` - Utility functions

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Adding New Frameworks

1. Create a new generator in `internal/generator/backend/` or `internal/generator/frontend/`
2. Implement the `BackendGenerator` or `FrontendGenerator` interface
3. Register the generator in `internal/generator/interfaces.go`
4. Add the framework to `internal/types/framework.go`

## 📄 License

This project is licensed under the MIT License.

## 🙏 Acknowledgments

- [Survey](https://github.com/AlecAivazis/survey) for interactive prompts
- [Cobra](https://github.com/spf13/cobra) for CLI framework
- All the amazing framework maintainers for their incredible work