from datetime import datetime

def calculate_priority(deadline, difficulty, importance):
    days_left = (deadline - datetime.now()).days

    if days_left <= 0:
        return 999

    priority = (importance * 2 + difficulty) / days_left
    return priority