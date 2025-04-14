# 📦 Distributed Message Queue System with Go, RabbitMQ, and Kubernetes

A scalable, fault-tolerant distributed messaging system built using Go, RabbitMQ, Docker, and Kubernetes.  
Designed to support reliable message publishing and consumption under high concurrency and traffic spikes.

---

## 🚀 Features
- Concurrent message publishing and consumption using Go routines.
- Automatic retry and error handling with RabbitMQ.
- Health-check endpoints for publisher and consumer services.
- Kubernetes manifests for cloud-native deployment and auto-scaling.
- Local development supported with Docker Compose.

---

## 🛠 Tech Stack
- **Language:** Go (Golang)
- **Message Broker:** RabbitMQ
- **Containerization:** Docker
- **Orchestration:** Kubernetes
- **Configuration:** Environment Variables (12-Factor App principles)
- **Local Dev:** Docker Compose

---

## 📚 Architecture Overview

```plaintext
[Publisher Service] --publish--> [RabbitMQ Queue] --consume--> [Consumer Service]
```

•	Publisher Service: Publishes messages to a RabbitMQ exchange.
•	RabbitMQ Broker: Queues, routes, and buffers messages.
•	Consumer Service: Subscribes to queues and processes incoming messages concurrently.

📁 Project Structure

```
distributed-message-queue/
├── cmd/
│   ├── publisher/
│   │   └── main.go
│   └── consumer/
│       └── main.go
├── internal/
│   ├── rabbitmq/
│   │   ├── publisher.go
│   │   └── consumer.go
│   └── config/
│       └── config.go
├── manifests/
│   ├── rabbitmq-deployment.yaml
│   ├── publisher-deployment.yaml
│   ├── consumer-deployment.yaml
│   └── namespace.yaml
├── Dockerfile.publisher
├── Dockerfile.consumer
├── docker-compose.yml
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

⚙️ Setup Instructions

1. Clone the Repository

```bash

git clone https://github.com/your-username/distributed-message-queue.git
cd distributed-message-queue

```

2. Environment Variables

Create a .env file (or configure environment manually):

```plaintext

RABBITMQ_HOST=rabbitmq
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_QUEUE=my_queue

```


🐳 Local Development (Docker Compose)

Start RabbitMQ, Publisher, and Consumer services locally:

```bash

docker-compose up --build

```

•	RabbitMQ Dashboard available at: http://localhost:15672
•	Default credentials: guest/guest

☁️ Kubernetes Deployment (Production-Ready)
1.	Ensure you have kubectl and a running Kubernetes cluster (e.g., Minikube, EKS).
2.	Apply the Kubernetes manifests:

```bash

kubectl apply -f manifests/namespace.yaml
kubectl apply -f manifests/rabbitmq-deployment.yaml
kubectl apply -f manifests/publisher-deployment.yaml
kubectl apply -f manifests/consumer-deployment.yaml

```

	3.	Monitor Pods:

```bash

kubectl get pods -n messaging

```

📊 Metrics and Monitoring (Optional)
•	Expose Prometheus metrics for RabbitMQ using plugins.
•	Monitor service health and queue depth over time.


📈 Future Improvements
	•	Implement dead-letter queues (DLQ) for failed messages.
	•	Add exponential backoff retry mechanism in consumers.
	•	Enable secure AMQP connections (SSL/TLS).
	•	Deploy a RabbitMQ cluster instead of a single instance.

👨‍💻 Author
•	Rahul Saxena