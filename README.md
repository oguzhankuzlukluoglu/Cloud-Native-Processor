# Cloud-Native Asynchronous Image Processor

This project is a **high-performance, microservices-based** image processing system built with **Go (Golang)** and orchestrated via **Kubernetes**. It demonstrates a robust asynchronous architecture using a message queue to handle resource-intensive tasks.



---

## System Architecture

The system is designed with a decoupled architecture to ensure scalability and fault tolerance:

1.  **API Service (Go/Fiber):** A high-speed RESTful API that handles file uploads and pushes metadata to the task queue.
2.  **Message Broker (Redis):** Serves as the communication bridge, buffering tasks between the API and workers.
3.  **Worker Service (Go):** An isolated background processor that consumes tasks from Redis and performs image operations.
4.  **Storage (K8s PVC):** A persistent volume layer that ensures uploaded and processed files survive pod restarts.



---

## Tech Stack

* **Language:** Go (Golang) 1.25+
* **API Framework:** [Gofiber/Fiber](https://gofiber.io/)
* **Queue/Cache:** Redis
* **Orchestration:** Kubernetes (K8s)
* **Containerization:** Docker
* **Infrastructure:** Kubernetes Persistent Volume Claims (PVC)

---

## Getting Started

### Prerequisites
* [Docker Desktop](https://www.docker.com/products/docker-desktop/) (with Kubernetes enabled)
* `kubectl` CLI

### Installation & Deployment

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/cloud-native-processor.git](https://github.com/your-username/cloud-native-processor.git)
    cd cloud-native-processor
    ```

2.  **Build Docker Images locally:**
    ```bash
    docker build -t cloud-native-api:v1 ./api
    docker build -t cloud-native-worker:v1 ./worker
    ```

3.  **Deploy to Kubernetes:**
    ```bash
    kubectl apply -f k8s/
    ```

4.  **Verify the deployment:**
    ```bash
    kubectl get pods
    ```

---

## Testing the Pipeline

Once the pods are `Running`, the API is exposed via **NodePort 30080**.

### Upload a Task
Use **Postman** or **cURL** to send a file:
```bash
curl -X POST http://localhost:30080/upload \
  -F "document=@sample_image.jpg"
