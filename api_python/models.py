from sqlalchemy import Column, Integer, Float
from database import Base  # Imports the declarative base class for SQLAlchemy models.

class MetricTable(Base):
    __tablename__ = "metrics"  # Defines the table name in the database.

    # Defines the columns for the metrics table.
    id = Column(Integer, primary_key=True, index=True)  # Unique identifier for each row.
    cpu_percent = Column(Float, nullable=False)
    memory_percent = Column(Float, nullable=False)
    disk_percent = Column(Float, nullable=False)
    timestamp = Column(Float, nullable=False)

    # More table models can be defined in this file.