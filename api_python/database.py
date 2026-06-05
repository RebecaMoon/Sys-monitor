import os
from dotenv import load_dotenv

from sqlalchemy import create_engine
from sqlalchemy.orm import declarative_base, sessionmaker


load_dotenv() # Load env variables from .env


DATABASE_URL = os.getenv("DATABASE_URL")# Read the database URL from the .env

if DATABASE_URL is None:
    raise ValueError("DATABASE_URL is not set in the environment variables")

engine = create_engine(DATABASE_URL) # Creates the SQLAlchemy engine used to connect to the database.

SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine) # Creates a factory for database sessions.

Base = declarative_base() # Creates the base class used to define database tables.