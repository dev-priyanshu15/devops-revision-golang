# DevOps Bootcamp Journey: From Code to Container to Cloud 🚀

This repository documents a practical journey through core DevOps concepts using a simple Golang application as the base. It serves as a complete guide from writing the first line of code to automating its deployment.

---

## 📚 Table of Contents
1. [Day 1: Golang Web Server Foundation](#-day-1-golang-web-server-foundation)
2. [Day 2: Dockerization (Containerization)](#-day-2-dockerization-containerization)
3. [Day 3: CI/CD Automation (GitHub Actions)](#-day-3-cicd-automation-github-actions)
4. [Day 4: Kubernetes Orchestration (K8s)](#-day-4-kubernetes-orchestration-k8s)
5. [How to Run Locally](#%EF%B8%8F-how-to-run-locally)

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

Here is the breakdown of the commands we ran and the logic behind them:

#### Phase 1: Setup (Setting up the Playground)
| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **1. Enable K8s** | Docker Desktop > Settings > Enable | Minikube had issues, so we used Docker Desktop's built-in K8s for a smoother experience. |
| **2. Switch Context** | `kubectl config use-context docker-desktop` | **Correction:** `kubectl` was targeting AWS Cloud. We forced it to talk to the **Local Cluster**. |
| **3. Verify** | `kubectl get nodes` | To confirm the "Master Node" is ready to receive orders. |

#### Phase 2: Pod Phase (The Freelancer - Manual Work)
| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **4. Create Pod** | `kubectl run my-first-pod --image=nginx` | **Hello World:** Created a standalone pod just to test if a container runs in K8s. |
| **5. Access App** | `kubectl port-forward pod/my-first-pod 8080:80` | **The Tunnel:** K8s is isolated (Locked Room). We opened a tunnel to view the site on localhost. |
| **6. Delete Pod** | `kubectl delete pod my-first-pod` | **Reliability Test:** We deleted the pod to see if it comes back. **Result:** It did not (No Auto-healing). |

#### Phase 3: Deployment Phase (The Manager - Production Work)
| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **7. Deploy** | `kubectl create deployment my-web --image=nginx` | **Hiring a Manager:** We stopped creating Pods manually. Now the Deployment manages them. |
| **8. Check List** | `kubectl get all` | To verify that the Manager (Deployment), Assistant (ReplicaSet), and Worker (Pod) are all active. |
| **9. Auto-Heal** | `kubectl delete pod <pod-name>` | **Murder Mystery:** Intentionally killed a pod. **Result:** Manager immediately created a new one (System stayed UP). |
| **10. Scaling** | `kubectl scale deployment my-web --replicas=3` | **Viral Mode:** Increased capacity from 1 to 3 servers with a single command to distribute load. |

#### Phase 4: Cleanup
| Step | Command | Logic (Why?) |
| :--- | :--- | :--- |
| **11. Delete All** | `kubectl delete deployment my-web` | Removed all resources to free up system RAM/CPU. |

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

### Option 3: Using Kubernetes (K8s)
```bash
# Create Deployment
kubectl create deployment go-app --image=nginx

# Access using Port Forward
kubectl port-forward deployment/go-app 8080:80
```

---
**Current Status:** Day 4 Completed ✅
**Next Up:** Day 5 - Networking (Services & LoadBalancers)
