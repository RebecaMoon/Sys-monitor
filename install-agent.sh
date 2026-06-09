#!/bin/bash

set -e

PROJECT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
AGENT_DIR="$PROJECT_DIR/agent"
SERVICE_FILE="/etc/systemd/system/sysmonitor-agent.service"
CURRENT_USER="$(whoami)"

echo "Installing SysMonitor Agent..."

cd "$AGENT_DIR"

python3 -m venv venv
source venv/bin/activate
pip install -r requirements.txt

sudo tee "$SERVICE_FILE" > /dev/null <<EOF
[Unit]
Description=SysMonitor Agent
After=docker.service network-online.target
Wants=network-online.target

[Service]
Type=simple
User=$CURRENT_USER
WorkingDirectory=$AGENT_DIR
ExecStart=$AGENT_DIR/venv/bin/python agent.py
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

sudo systemctl daemon-reload
sudo systemctl enable sysmonitor-agent
sudo systemctl start sysmonitor-agent

echo "SysMonitor Agent installed and started successfully."