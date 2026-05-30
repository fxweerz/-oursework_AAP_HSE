package models

import (
	"time"
)

type Task struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Deadline       time.Time `json:"deadline"`
	Importance     int       `json:"importance"`
	Difficulty     int       `json:"difficulty"`
	EstimatedHours int       `json:"estimated_hours"`
	Priority       float64   `json:"priority"`
	UserID         int       `json:"user_id"`
}
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (t *Task) CalculatePriority() {
	daysLeft := time.Until(t.Deadline).Hours() / 24	
	if daysLeft < 0 {
		daysLeft = 0
	}
	deadlineScore := 10 / (daysLeft + 1)
	importanceScore := float64(t.Importance * 2)
	difficultyScore := float64(t.Difficulty)
	timeScore := float64(t.EstimatedHours) / 2

	t.Priority = deadlineScore + importanceScore + difficultyScore + timeScore
}