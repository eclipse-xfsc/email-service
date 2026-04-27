# 📧 email-service

A lightweight service responsible for sending emails based on incoming events from **NATS**.

It is designed to handle system-driven email workflows such as:
- User registration confirmation
- Backup and restore notifications
- Account management notifications
- Other event-driven email communications

---

## 🚀 Architecture Overview

- 📩 **Input:** NATS events
- ⚙️ **Processing:** Email message composition and validation
- 📤 **Output:** SMTP/email provider delivery

---

## 🧪 Test and Deploy

CI/CD is handled via **GitLab pipelines**.

---

## 🏃 Local Run

### 1. Using Makefile (recommended)

```bash
make docker-compose-run
```

2. Using Docker Compose directly
```bash
docker compose -f deployment/docker/docker-compose.yml rm

docker compose -f deployment/docker/docker-compose.yml \
  --env-file=.env \
  --env-file=.env.credentials \
  up --build --detach
```

## 🧪 Testing via NATS (Manual)

You can manually trigger the service by publishing a test event to NATS.

Windows PowerShell
```bash
Get-Content .\email-event.json -Raw | nats pub email.receive
```

Linux / macOS (alternative)
```bash
cat email-event.json | nats pub email.receive
```

This allows you to simulate incoming email events without needing a full system flow (e.g., registration or backup process).

## 🧰 Prerequisites

Before running locally, ensure you have:

Docker & Docker Compose installed
Go (if running outside containers)
.env and .env.credentials properly configured

## ⚙️ Configuration

Environment variables are required for proper service operation.

Required files:

.env
.env.credentials

⚠️ Ensure sensitive credentials are not committed to the repository.

## 📦 Installation

TBD

## 📖 Usage

The service listens to NATS subjects and processes incoming email events automatically.

Example event types:

user.registered
user.password.reset
backup.completed
🛠️ Development

TBD

## 🧪 Testing

Run unit tests:

go test ./...

TBD

🤝 Contributing

TBD

## 📄 License

TBD

## 📌 Project Status

Active development 🚧