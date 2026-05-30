from sqlalchemy.orm import DeclarativeBase
from sqlalchemy import  Column, Integer, String, Float, DateTime, ForeignKey
from datetime import datetime

class Base(DeclarativeBase): pass


'''
type Task struct {
	ID             int
	Title          string
	Deadline       time.Time
	Importance     int
	Difficulty     int
	EstimatedHours int
	Priority       float64
}
'''
class Task(Base):

    __tablename__ = "tasks"

    id = Column(Integer, primary_key=True, index=True)
    title = Column(String, nullable=False, index=True)
    deadline = Column(DateTime, nullable=False)
    importance = Column(Integer, nullable=False)
    difficulty = Column(Integer, nullable=False)
    estimated_hours = Column(Integer, nullable=False)
    priority = Column(Float, nullable=False)
    user_id = Column(Integer, ForeignKey("users.id"))


'''
type User struct {
	ID       int
	Name     string
	Password string
	Tasks    []Task
}
'''

class User(Base):

    __tablename__ = "users"

    id = Column(Integer, primary_key=True, index=True)
    name = Column(String, nullable=False, index=True)
    password = Column(String, nullable=False)
