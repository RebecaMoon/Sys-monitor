from fastapi import FastAPI, Depends
from pydantic import BaseModel
from datetime import datetime
from sqlalchemy.orm import Session
from fastapi import HTTPException
# To run this use: uvicorn main:app --reload

import database # Import the database config and model.
import models

app = FastAPI()

models.Base.metadata.create_all(bind=database.engine) # It creates the tables if they don't exist.


def get_db(): # This function manages the database sessions.
    db = database.SessionLocal()
    try:
        yield db  # Open the session.
    finally:
        db.close() # When the endpoint ends, it closes the connection.

class SystemMetrics(BaseModel): # This function uses Pydantic to verify the data received from the agent, it checks that the data types are correct.
    cpu_percent: float
    memory_percent: float
    disk_percent: float
    timestamp: float

class MetricResponse(BaseModel):  #This function will be the response model used in the metrics/latest endpoint.
    id: int
    cpu_percent: float
    memory_percent: float
    disk_percent: float
    timestamp: float

@app.post("/metrics")
def receive_metrics(metrics: SystemMetrics, db: Session = Depends(get_db)): # This is the main function.
    # metrics contains the values that have already been checked with Pydantic.
    # db is the "session variable" we are going to use to talk with Postgres.

   
    # Saving process ----------------------------------------------------------------------
    db_metric = models.MetricTable(  # Convert our data into an SQL database row object.
        cpu_percent=metrics.cpu_percent,
        memory_percent=metrics.memory_percent,
        disk_percent=metrics.disk_percent,
        timestamp=metrics.timestamp
    )
    
    
    db.add(db_metric) # Save the row in the db variable (it isn't saved in the database yet).
   
    db.commit()  # Commit it to SAVE IT to the Postgres database container.
    
    db.refresh(db_metric) # Refresh the object db_metric to get the auto-generated data.



    readable_time = datetime.fromtimestamp(metrics.timestamp).strftime("%Y-%m-%d %H:%M:%S")

    print(
        f"Saved to DB (ID: {db_metric.id}) | "
        f"CPU: {metrics.cpu_percent}% | "
        f"Memory: {metrics.memory_percent}% | "
        f"Disk: {metrics.disk_percent}% | "
        f"Time: {readable_time}"
    )

    
    return {
        "status": "success",
        "database_id": db_metric.id
    }


@app.get("/metrics/latest", response_model=MetricResponse)
def get_latest_metrics(db: Session = Depends(get_db)):
    latest_metric = (
        db.query(models.MetricTable)
        .order_by(models.MetricTable.id.desc())
        .first()
    )

    if latest_metric is None:
        raise HTTPException(
            status_code=404,
            detail="No metrics available"
        )

    return {
        "id": latest_metric.id,
        "cpu_percent": latest_metric.cpu_percent,
        "memory_percent": latest_metric.memory_percent,
        "disk_percent": latest_metric.disk_percent,
        "timestamp": latest_metric.timestamp,
    }


@app.get("/metrics/history", response_model=list[MetricResponse])
def get_metrics_history(db: Session = Depends(get_db)):
    metrics = (
        db.query(models.MetricTable)
        .order_by(models.MetricTable.id.desc())
        .limit(25)
        .all()
    )

    metrics.reverse()

    return metrics
