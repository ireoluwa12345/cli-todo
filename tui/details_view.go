package tui

import (
	"cli/todo/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DetailsModel struct {
	task   models.Tasks
	width  int
	height int
}

func (t *Tui) NewDetailsModel(task models.Tasks, width int, height int) DetailsModel {
	return DetailsModel{
		task:   task,
		width:  width,
		height: height,
	}
}

func (m DetailsModel) Init() tea.Cmd {
	return nil
}

func (m DetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m DetailsModel) View() string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	labelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF69B4"))

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5"))

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#A7A7A7"))

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Margin(1, 2).
		Width(m.width - padding)

	output := titleStyle.Render("Task Details") + "\n\n"
	output += labelStyle.Render("ID: ") + valueStyle.Render(m.task.Id) + "\n"
	output += labelStyle.Render("Name: ") + valueStyle.Render(m.task.Name) + "\n"
	output += labelStyle.Render("Description: ") + descStyle.Render(m.task.Description) + "\n"

	var statusStyle lipgloss.Style
	switch m.task.Status {
	case "not_started":
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")) // red
	case "in_progress":
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFF00")) // yellow
	case "completed":
		statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#43BF6D")) // green
	default:
		statusStyle = valueStyle
	}
	output += labelStyle.Render("Status: ") + statusStyle.Render(m.task.Status) + "\n"

	output += labelStyle.Render("Priority: ") + valueStyle.Render(m.task.Priority) + "\n"
	output += labelStyle.Render("Due Date: ") + valueStyle.Render(m.task.DueDate) + "\n"
	output += labelStyle.Render("Completed Date: ") + valueStyle.Render(m.task.CompletedDate) + "\n"

	return boxStyle.Render(output)
}
