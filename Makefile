.PHONY: help dev-backend dev-frontend build-backend build-frontend test clean

help:
	@echo "MpangoWaCuddles.com - Development Commands"
	@echo ""
	@echo "  make dev-backend    - Run Go backend server (port 8080)"
	@echo "  make dev-frontend   - Run Next.js frontend (port 3000)"
	@echo "  make build-backend  - Build Go binary"
	@echo "  make build-frontend - Build Next.js production bundle"
	@echo "  make test          - Run Go tests"
	@echo "  make clean         - Clean build artifacts"

dev-backend:
	@echo "Starting Go backend server..."
	@cd cmd/server && go run main.go

dev-frontend:
	@echo "Starting Next.js frontend..."
	@cd web/nextjs-app && npm run dev

build-backend:
	@echo "Building Go backend..."
	@go build -o bin/server cmd/server/main.go

build-frontend:
	@echo "Building Next.js frontend..."
	@cd web/nextjs-app && npm run build

test:
	@echo "Running tests..."
	@go test ./...

clean:
	@echo "Cleaning build artifacts..."
	@rm -rf bin/
	@rm -rf web/nextjs-app/.next/
	@rm -rf web/nextjs-app/out/

install-deps:
	@echo "Installing Go dependencies..."
	@go mod tidy
	@echo "Installing Node.js dependencies..."
	@cd web/nextjs-app && npm install

