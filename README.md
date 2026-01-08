# DevOps Bootcamp - Day 1: Go Basics & Git

## What I Learned
1. **Go Basics**: Created a simple HTTP web server using Go.
2. **Endpoints**:
   - `/`: Welcome message.
   - `/health`: Checks server status and user identity.
3. **Environment Variables**: Used `.env` file to manage configuration (Port) securely using `godotenv`.
4. **Git & GitHub**: Initialized repository and pushed code to GitHub.

## How to Run
1. Install dependencies:
   ```bash
   go mod tidy
   ```
2. Create a `.env` file:
   ```bash
   echo "PORT=8080" > .env
   ```
3. Run the application:
   ```bash
   go run main.go
   ```
4. Test:
   ```bash
   curl http://localhost:8080/health
   ```
