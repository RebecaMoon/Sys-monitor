from sqlalchemy import create_engine
from sqlalchemy.orm import declarative_base, sessionmaker



DATABASE_URL = "postgresql://reb_admin:Z4mperf1@localhost:5432/monitor_db" # Database connection URL.

 
engine = create_engine(DATABASE_URL) # Creates the SQLAlchemy engine used to connect to the database.


SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine) # Creates a factory for database sessions.


Base = declarative_base() # Creates the base class used to define database tables.