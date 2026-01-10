# DevOps Bootcamp Journey: From Code to Container to Cloud 🚀

This repository documents a practical journey through core DevOps concepts using a simple Golang application as the base. It serves as a complete guide from writing the first line of code to automating its deployment.

---

## 📚 Table of Contents
1. [Day 1: Golang Web Server Foundation](#-day-1-golang-web-server-foundation)
2. [Day 2: Dockerization (Containerization)](#-day-2-dockerization-containerization)
3. [Day 3: CI/CD Automation (GitHub Actions)](#-day-3-cicd-automation-github-actions)
4. [How to Run Locally](#%EF%B8%8F-how-to-run-locally)

---

## 🗓️ Day 1: Golang Web Server Foundation
**Goal:** Build a basic HTTP server to understand web service fundamentals.

### Key Learnings:
- **Go's `net/http`:** Used standard library to create a lightweight server.
- **Handlers & Routing:**
    - `/`: Home page welcome message.
    - `/health`: An endpoint to check server status and identity.
- **Configuration Management:** Used `godotenv` to safely read configuration (like `PORT`) from a `.env` file, mimicking real-world app configuration.

---

## 🗓️ Day 2: Dockerization (Containerization)
**Goal:** Solve the "It works on my machine" problem by packaging the app.

### 1. Concept: The Container Revolution
Docker packages the application along with its environment (OS settings, dependencies) into a portable container. This ensures consistency across dev, test, and production environments.

![Docker Container Architecture](https://www.docker.com/wp-content/uploads/2021/11/docker-containerized-app-uct-diagram.png)
*Concept: Unlike VMs, containers share the host OS kernel, making them lightweight and fast.*

### 2. The Recipe: `Dockerfile`
We created a `Dockerfile` to define how to build our application image.

```dockerfile
# Base Image: Small Linux + Go
FROM golang:1.21-alpine
# Set working directory
WORKDIR /app
# Copy source code
COPY . .
# Build the binary
RUN go build -o main .
# Expose port for documentation
EXPOSE 8080
# Command to run on container start
CMD ["./main"]
```

### 3. Essential Docker Commands Learned:
- **Build Image:** `docker build -t go-app .`
- **Run Container (with port mapping):** `docker run -d -p 9090:8080 --name my-go-container go-app`
- **Check Running Containers:** `docker ps`
- **View Logs (Debugging):** `docker logs my-go-container`
- **Enter Container (Debugging):** `docker exec -it my-go-container sh`

---

## 🗓️ Day 3: CI/CD Automation (GitHub Actions)
**Goal:** Automate the build and test process so we don't have to do it manually.

### 1. Concept: The Pipeline
We define a workflow that triggers automatically whenever code is pushed to GitHub. This ensures continuous integration.

![CI/CD Pipeline Concept](https://resources.github.com/assets/images/devops/ci-cd/ci-cd-flowchart.png)
*Concept: Code Push -> Automated Build & Test -> Feedback (Green/Red tick).*

### 2. The Automation Robot: `.github/workflows/ci.yml`
We wrote a YAML file to instruct GitHub Actions. We used the latest stable action versions (`@v4`, `@v5`) for security and features.

```yaml
name: Go CI Pipeline

on:
  push:
    branches: [ "main" ]

jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - name: Build App
        run: go build -v ./...
      - name: Test Docker Build
        run: docker build . -t my-go-app:test
```

---

## 🛠️ How to Run Locally

### Option 1: Using Go (Directly on Host)
```bash
# Install dependencies
go mod tidy

# Setup configuration
echo "PORT=8080" > .env

# Run the server
go run main.go
```
Access at: `http://localhost:8080/health`

### Option 2: Using Docker (In a Container)
```bash
# Build the image
docker build -t go-app-final .

# Run container mapping port 9090 on host to 8080 inside
docker run -d -p 9090:8080 --name final-container go-app-final
```
Access at: `http://localhost:9090/health`

---
**Next Up:** Day 4 - Orchestration with Kubernetes (K8s) 🚢
