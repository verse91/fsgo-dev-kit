package templates

// GitignoreTemplate returns the .gitignore template
func GitignoreTemplate() string {
	return `# ================== Dependencies ==================
node_modules/
.yarn/*
.pnp*
.pnpm-debug.log*

# Keep necessary Yarn folders
!.yarn/patches
!.yarn/plugins
!.yarn/releases
!.yarn/versions

# ================== Build & Production ==================
.build/
dist/
out/
.next/
.vercel/
.turbo/
.output/

# ================== Testing & Coverage ==================
coverage/
*.test.js
*.test.ts

# ================== Env files ==================
.env*
**/.env*
!.env.example
!.env.local.example
!client/.env.example
!server/.env.example

# ================== Logs & Cache ==================
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*
.eslintcache
.cache/
*.tsbuildinfo
log/
**/*.log

# ================== Editor / IDE ==================
.vscode/
.idea/
*.sublime-*
*.code-workspace

# ================== System ==================
.DS_Store
Thumbs.db
ehthumbs.db
*.swp
*.swo

# ================== Docker ==================
**/docker-compose.override.yml
**/docker-compose.local.yml

# ================== Frontend (Next.js) ==================
client/node_modules/
client/.next/
client/out/
client/.env*
client/.pnp
client/.vercel
client/yarn.lock
client/pnpm-lock.yaml
client/package-lock.json
client/bun.lockb
client/npm-debug.log*
client/yarn-*.log

# ================== Backend (Go) ==================
server/bin/
server/cmd/server/tmp/
server/tmp/
server/.env*
server/coverage/
server/*.test
server/go.sum
server/go.work*
server/.DS_Store
server/internal/video_pipeline/videos`
}

// MakefileTemplate returns the Makefile template
func MakefileTemplate() string {
	return `.PHONY: docker-b docker-f docker-up docker-down docker-rb docker-rf run

# d = docker
# Windows users must run docker desktop before running these commands
db:
	@docker compose up --build b

df:
	@docker compose up --build f

dup:
	@docker compose up --build

ddown:
	@docker compose down

drb:
	@docker compose build b

drf:
	@docker compose build f

b: #backend
	@cd server && air

f: #frontend
	@cd client && bun run dev

testdb: #test database
	@go run -C server ./cmd/test

run:
	@$(MAKE) -j2 b f

stop:
	@xargs kill < .dev.pids 2>/dev/null || true
	@rm -f .dev.pids`
}

// ReadmeTemplate returns the README.md template
func ReadmeTemplate() string {
	return `# Fullstack Project

A modern fullstack application with Next.js frontend and Go Fiber backend.

## Project Structure

` + "```" + `
client/          # Next.js frontend
server/          # Go Fiber backend
` + "```" + `

## Quick Start

1. Install dependencies:
   ` + "```bash" + `
   # Install Go dependencies
   cd server && go mod tidy && cd ..
   
   # Install Node.js dependencies  
   cd client && bun install && cd ..
   ` + "```" + `

2. Set up environment variables:
   ` + "```bash" + `
   cp client/.env.example client/.env
   cp server/.env.example server/.env
   ` + "```" + `

3. Start development servers:
   ` + "```bash" + `
   make run
   ` + "```" + `

## Available Commands

- ` + "`make run`" + ` - Start both frontend and backend
- ` + "`make b`" + ` - Start backend only
- ` + "`make f`" + ` - Start frontend only
- ` + "`make testdb`" + ` - Test database connection
- ` + "`make stop`" + ` - Stop all running processes

## Tech Stack

### Frontend
- Next.js 14+
- TypeScript
- Tailwind CSS
- React

### Backend
- Go
- Fiber v3
- Zap Logger
- Air (hot reload)

## Development

The backend runs on http://localhost:8080
The frontend runs on http://localhost:3000`
}