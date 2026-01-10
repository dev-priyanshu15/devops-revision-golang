# DevOps Bootcamp Journey: From Code to Container to Cloud 🚀

This repository documents a practical journey through core DevOps concepts using a simple Golang application as the base. It serves as a complete guide from writing the first line of code to automating its deployment.

---

## 📚 Table of Contents
1. [Day 1: Golang Web Server Foundation](#-day-1-golang-web-server-foundation)
2. [Day 2: Dockerization (Containerization)](#-day-2-dockerization-containerization)
3. [Day 3: CI/CD Automation (GitHub Actions)](#-day-3-cicd-automation-github-actions)
4. [Day 4: Kubernetes Orchestration (K8s)](#-day-4-kubernetes-orchestration-k8s)
5. [Day 5: Networking & Configuration (Pro Setup)](#-day-5-networking--configuration-pro-setup)
6. [How to Run Locally](#%EF%B8%8F-how-to-run-locally)

---

## 🗓️ Day 1: Golang Web Server Foundation
**Goal:** Build a basic HTTP server to understand web service fundamentals.

### Key Learnings:
- **Go's `net/http`:** Used standard library to create a lightweight server.
- **Handlers & Routing:**
    - `/`: Home page welcome message.
    - `/health` endpoint: Checks server status and identity.
- **Configuration Management:** Used `godotenv` to safely read configuration (like `PORT`) from a `.env` file.

---

## 🗓️ Day 2: Dockerization (Containerization)
**Goal:** Solve the "It works on my machine" problem by packaging the app.

### 1. Concept: The Container Revolution
Docker packages the application along with its environment (OS settings, dependencies) into a portable container. Unlike VMs, containers share the host OS kernel, making them lightweight.

### 2. The Recipe: `Dockerfile`
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
```

### 3. Essential Commands
- **Build:** `docker build -t go-app .`
- **Run:** `docker run -d -p 9090:8080 --name my-container go-app`

---

## 🗓️ Day 3: CI/CD Automation (GitHub Actions)
**Goal:** Automate the build and test process.

### 1. Concept: The Pipeline
Defined a workflow in `.github/workflows/ci.yml` that triggers on code push.
**Flow:** Code Push -> Automated Build & Test -> Feedback (Green/Red tick).

### 2. The YAML Configuration
Used latest actions (`@v4`, `@v5`) for stability.
```yaml
name: Go CI Pipeline
on:
  push:
    branches: [ "main" ]
jobs:
  build-test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      - run: go build -v ./...
      - run: docker build . -t my-go-app:test
```

---

## 🗓️ Day 4: Kubernetes Orchestration (K8s)
**Goal:** Manage clusters, move from single container to many, and achieve High Availability.

### 1. The Concepts (Theory)
- **Pod:** The smallest unit (Freelancer). It wraps the container. If it dies, it stays dead.
- **Deployment:** The Manager (Company). It ensures a specific number of pods are always running.
  - **Auto-Healing:** If a pod crashes/deleted, Deployment creates a new one immediately.
  - **Scaling:** Can increase replicas (copies) to handle traffic spikes.

### 2. 🚀 Step-by-Step Execution Log

| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **1. Setup** | `kubectl config use-context docker-desktop` | Switched context to Local K8s (Docker Desktop). |
| **2. Verify** | `kubectl get nodes` | Confirmed the Cluster is Ready. |
| **3. Pod** | `kubectl run my-first-pod --image=nginx` | **Hello World:** Created a standalone pod to test K8s. |
| **4. Access** | `kubectl port-forward pod/my-first-pod 8080:80` | **Tunnel:** Accessed isolated pod via localhost. |
| **5. Deploy** | `kubectl create deployment my-web --image=nginx` | **Manager:** Created Deployment for Auto-Healing. |
| **6. Heal** | `kubectl delete pod <pod-name>` | **Test:** Killed pod manually. **Result:** New one created instantly. |
| **7. Scale** | `kubectl scale deployment my-web --replicas=3` | **Viral Mode:** Increased capacity from 1 to 3 servers. |
| **8. Clean** | `kubectl delete deployment my-web` | Removed resources to save memory. |

---

## 🗓️ Day 5: Networking & Configuration (Pro Setup)
**Goal:** Deploy the **Real Golang App**, expose it to the world, and manage settings securely.

### 1. The Concepts
- **ConfigMap:** A separate "Diary" for settings (like `PORT`). Decouples config from code.
- **Service (LoadBalancer):** The "Receptionist". Gives a stable IP/Port to access the App, even if Pods die/restart.

### 2. The Infrastructure Code (k8s-final.yaml)
We used a single YAML file to define our entire infrastructure.
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: golang-deployment
spec:
  replicas: 2
  selector:
    matchLabels:
      app: my-go-app
  template:
    metadata:
      labels:
        app: my-go-app
    spec:
      containers:
      - name: go-container
        image: my-go-app:v1      # Uses our local Go image
        imagePullPolicy: Never
        ports:
        - containerPort: 8080
        envFrom:                 # Links to ConfigMap
        - configMapRef:
            name: my-go-config
---
apiVersion: v1
kind: Service
metadata:
  name: golang-service
spec:
  type: LoadBalancer             # Exposes App to localhost
  selector:
    app: my-go-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
```

### 3. 🚀 Step-by-Step Execution Log

| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **1. Build Image** | `docker build -t my-go-app:v1 .` | **Packaging:** K8s needs an image. We packaged our *actual* Go code. |
| **2. ConfigMap** | `kubectl create configmap my-go-config --from-literal=PORT=8080` | **Settings:** Saved `PORT=8080` in K8s so the app knows where to run. |
| **3. Apply** | `kubectl apply -f k8s-final.yaml` | **Execution:** Deployed both App and Network using the YAML file above. |
| **4. Verify** | `kubectl get all` | **Check:** Confirmed Pods are `Running` and Service has an IP. |
| **5. Access** | Browser -> `http://localhost/health` | **Final Test:** Accessed the Go App directly via browser (No port-forward needed!). |

---

## 🛠️ How to Run Locally

### Option 1: Using Go
```bash
go run main.go
# Access at http://localhost:8080/health
```

### Option 2: Using Docker
```bash
docker run -d -p 9090:8080 go-app-final
# Access at http://localhost:9090/health
```

### Option 3: Using Kubernetes (The Professional Way) 🏆
```bash
# 1. Build Image
docker build -t my-go-app:v1 .

# 2. Apply Infrastructure
kubectl apply -f k8s-final.yaml

# 3. Access App
# Open Browser: http://localhost/health
```

---
**Current Status:** Day 5 Completed ✅
**Next Up:** Day 6 - Database Integration (MongoDB)
