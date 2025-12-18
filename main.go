package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Go Task Tracker")
	myWindow.Resize(fyne.NewSize(400, 350)
                  
	var tasks []string

	taskEntry := widget.NewEntry()
	taskEntry.SetPlaceHolder("Enter a new task")

	taskList := widget.NewLabel("No tasks yet")
	taskCount := widget.NewLabel("Total tasks: 0")

	addButton := widget.NewButton("Add Task", func() {
		if taskEntry.Text == "" {
			taskList.SetText(" Please enter a task")
			return
		}

		tasks = append(tasks, taskEntry.Text)
		taskEntry.SetText("")

		updateTaskDisplay(taskList, taskCount, tasks)
	})

	clearButton := widget.NewButton("Clear All Tasks", func() {
		tasks = []string{}
		updateTaskDisplay(taskList, taskCount, tasks)
	})

	content := container.NewVBox(
		widget.NewLabelWithStyle(
			"Simple Task Tracker",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		taskEntry,
		addButton,
		taskList,
		taskCount,
		clearButton,
	)

	myWindow.SetContent(container.NewCenter(content))
	myWindow.ShowAndRun()
}

func updateTaskDisplay(taskList *widget.Label, taskCount *widget.Label, tasks []string) {
	if len(tasks) == 0 {
		taskList.SetText("No tasks yet")
		taskCount.SetText("Total tasks: 0")
		return
	}

	displayText := ""
	for i, task := range tasks {
		displayText += fmt.Sprintf("%d. %s\n", i+1, task)
	}

	taskList.SetText(displayText)
	taskCount.SetText(fmt.Sprintf("Total tasks: %d", len(tasks)))
}
