# Sys-monitor


It is a containerized system monitoring solution designed to track CPU, RAM and Disk usage metrics.


## Architecture

The system consists of three main services working together:

1. **Agent:** A Python-based script that periodically collects system metrics (CPU/RAM/DISK).
2. **API:** A FastAPI-based backend that receives, validates, and stores incoming metrics.
3. **Database:** A PostgreSQL instance for persistent storage of historical telemetry data.


## Technologies

- **Backend:** Python (FastAPI)
- **Monitoring:** Python (psutil)
- **Database:** PostgreSQL
- **Orchestration:** Docker
- **Versioning:** Git


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


## How to run

### Prerequisites

Before running the project, make sure you have installed:

- Docker
- Docker Compose
- Python 3
- Git
- A Linux distribution with systemd

> The API and database run inside Docker containers, while the monitoring agent runs directly on the host machine as a systemd service.

### 1. Clone the repository

```bash
git clone https://github.com/RebecaMoon/sys-monitor.git
cd sys-monitor
```

### 2. Start the API and database

```bash
docker compose up -d --build
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