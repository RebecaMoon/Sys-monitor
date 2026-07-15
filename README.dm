# Sys-monitor


It is a containerized system monitoring solution designed to collect, store, and visualize CPU, RAM, and Disk usage metrics through a real-time web dashboard.


## Architecture

The system consists of four main services working together:

1. **Agent:** A Python-based script that periodically collects system metrics (CPU/RAM/DISK).
2. **API:** A FastAPI-based backend that receives, validates, and stores incoming metrics.
3. **Frontend:** A React-based dashboard served by Nginx that displays real-time metrics and historical usage graphs.
4. **Database:** A PostgreSQL instance for persistent storage of historical telemetry data.

> This project also includes an alternative API implementation written in Go. The Go API works with the existing Python agent and stores metrics in the same PostgreSQL database.

## Technologies

- **Backend:** Python (FastAPI)
- **Backend alternative:** Go
- **Monitoring:** Python (psutil)
- **Database:** PostgreSQL
- **Orchestration:** Docker
- **Versioning:** Git
- **Frontend:** React (Vite)
- **Web Server:** Nginx


## Data flow

```txt
[ Incoming JSON ] ──>  1. main.py (Gatekeeper)
                           │
                           ▼ Uses Pydantic to check structure
                       2. models.py (Translator)
                           │
                           ▼ Converts valid data into a database row
                       3. database.py (Construction Worker)
                           │
                           ▼ Opens a session & pushes raw SQL 
                    [ PostgreSQL Container ]
```

## Dashboard flow

```txt
          Browser
              │
              ▼
      Nginx (React Frontend)
              │
      GET /api/metrics/*
              │
              ▼
          FastAPI
              │
              ▼
         PostgreSQL
```

## How to run

### Prerequisites

Before running the project, make sure you have installed:

- Docker
- Docker Compose
- Python 3
- Node.js (only required for frontend development)
- Git
- A Linux distribution with systemd

> The API and database run inside Docker containers, while the monitoring agent runs directly on the host machine as a systemd service.

### 1. Clone the repository

```bash
git clone https://github.com/RebecaMoon/sys-monitor.git
cd sys-monitor
```

### 2. Start the application

```bash
docker compose up -d --build
```
After the containers start:
- Frontend: http://localhost:5173
- API: http://localhost:8000

#### To run the Go API implementation instead, use:

```bash
docker compose -f compose.go.yml up -d --build
```

### 3. Install the monitoring agent

```bash
chmod +x install-agent.sh
./install-agent.sh
```

### 4. Start the agent

```bash
sudo systemctl start sysmonitor-agent
```


## Author

Created by Oscar Cuevas