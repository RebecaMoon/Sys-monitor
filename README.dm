# Sys-monitor


It is a containerized system monitoring solution designed to track CPU, RAM and Disk usage metrics.


## Architecture

The system consists of three main services working together:

1. **Agent:** A Python-based script that periodically collects system metrics (CPU/RAM/DISK).
2. **API** A FastAPI-based backend that receives, validates, and stores incoming metrics.
3. **Database:** A PostgreSQL instance for persistent storage of historical telemetry data.


## Technologies

- **Backend:** Python (FastAPI)
- **Monitoring:** Python (psutil)
- **Database:** PostgreSQL
- **Orchestration:** Docker
- **Versioning:** Git


--*Created by Oscar Cuevas

