package handlers

import (
	"context"
	"net/http"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"

	db "task-prioritizer/internal/database"
	"task-prioritizer/internal/models"
)

func CreateTask(c *gin.Context) {

	title := c.PostForm("title")
	deadlineStr := c.PostForm("deadline")

	deadline, _ := time.Parse("2006-01-02", deadlineStr)

	importanceStr := c.PostForm("importance")
	difficultyStr := c.PostForm("difficulty")
	hoursStr := c.PostForm("hours")
	importance, _ := strconv.Atoi(importanceStr)
	difficulty, _ := strconv.Atoi(difficultyStr)
	hours, _ := strconv.Atoi(hoursStr)

	task := models.Task{
		Title: title,
		Deadline: deadline,
		Importance: importance,
		Difficulty: difficulty,
		EstimatedHours: hours,
	}
	if err := task.Validate(); err != nil {
		c.String(400, err.Error())
		return
	}
	task.CalculatePriority()

	_, err := db.DB.Exec(context.Background(),
	`INSERT INTO tasks
	(title, deadline, importance, difficulty, estimated_hours, priority)
	VALUES ($1,$2,$3,$4,$5,$6)`,
	task.Title,
	task.Deadline,
	task.Importance,
	task.Difficulty,
	task.EstimatedHours,
	task.Priority,
	)

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func ListTasks(c *gin.Context) {

	rows, _ := db.DB.Query(context.Background(),
		`SELECT id,title,priority FROM tasks
		 ORDER BY priority DESC`)

	type Task struct {
		ID int
		Title string
		Priority float64
	}

	var tasks []Task

	for rows.Next() {
		var t Task
		rows.Scan(&t.ID, &t.Title, &t.Priority)
		tasks = append(tasks, t)
	}

	c.HTML(200, "index.html", gin.H{
		"tasks": tasks,
	})
}

