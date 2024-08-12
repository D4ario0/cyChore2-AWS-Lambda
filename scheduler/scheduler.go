package scheduler

import (
	"fmt"
	"lambda-cychore/types"
	"time"
)

// Constants for task descriptions
const (
	TaskTakeOutTrash    = "Sacar Basura / Correspondencia"
	TaskDeepCleaning    = "Limpieza profunda (%s)"
	TaskCleanCommonArea = "Barrer área común / Limpiar Baño"
	TaskSweepKitchen    = "Barrer Cocina"
)

// Constants for start date of the chore cycle
const (
	StartYear  = 2024
	StartMonth = 7
	StartDay   = 22
)

func AssignTasks(user *types.User) error {
	weeksSinceStart := calculateWeeksSinceStart(time.Now())
	cycleOffset := (weeksSinceStart % 8) / 4 // This returns 0 or 1 depending on the current chore cycle.

	specialTask := determineSpecialTask(weeksSinceStart, cycleOffset)
	tasks := []string{TaskTakeOutTrash, specialTask, TaskCleanCommonArea, TaskSweepKitchen}

	// Determine the task based on the user's buffer index and the current week.
	taskIndex := (user.BufferIndex + weeksSinceStart) % len(tasks)
	user.Task = tasks[taskIndex]
	return nil
}

func calculateWeeksSinceStart(currentDate time.Time) int {
	startDate := time.Date(StartYear, StartMonth, StartDay, 0, 0, 0, 0, time.UTC)
	weeks := int(currentDate.Sub(startDate).Hours() / (24 * 7))
	return weeks
}

func determineSpecialTask(weeksSinceStart int, cycleOffset int) string {
	if weeksSinceStart%2 == cycleOffset {
		return fmt.Sprintf(TaskDeepCleaning, "Area Común")
	}
	return fmt.Sprintf(TaskDeepCleaning, "Cocina")
}
