from fastapi import FastAPI
from pydantic import BaseModel
from datetime import datetime

app = FastAPI()

class SystemMetrics(BaseModel):
    cpu_percent: float
    memory_percent: float
    disk_percent: float
    timestamp: float

@app.post("/metrics")
def receive_metrics(metrics: SystemMetrics):

    readable_time = datetime.fromtimestamp(metrics.timestamp)
    formatted_time = readable_time.strftime("%d-%m-%y %H:%M:%S")
    
    print(f"Cpu: {metrics.cpu_percent}%")
    print(f"Memory: {metrics.memory_percent}%")
    print(f"Disk: {metrics.disk_percent}%")
    print(f"Time: {formatted_time}")
    print("------------------------------")
    return {"status": "success", "received": metrics.timestamp}