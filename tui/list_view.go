package tui

import (
	"cli/todo/models"

	tbl "github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type ListModel struct {
	tasks  []models.Tasks
	cursor int
	width  int
	height int
}

func (t *Tui) NewListModel(tasks []models.Tasks) ListModel {

	return ListModel{
		tasks:  tasks,
		cursor: 0,
		width:  0,
		height: 0,
	}
}

func (m ListModel) Init() tea.Cmd {
	return nil
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.tasks) > 0 {
				t := Tui{}
				detailsModel := t.NewDetailsModel(m.tasks[m.cursor], m.width, m.height)
				return detailsModel, nil
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m ListModel) View() string {
	if len(m.tasks) == 0 {
		return "No tasks found.\n"
	}

	singleColumnSize := (m.width - padding) / 6

	table := tbl.New(
		tbl.WithColumns([]tbl.Column{
			{Title: "ID", Width: singleColumnSize},
			{Title: "Name", Width: singleColumnSize * 2},
			{Title: "Status", Width: singleColumnSize},
			{Title: "Due Date", Width: singleColumnSize},
			{Title: "Completed Date", Width: singleColumnSize},
		}),
	)

	rows := []tbl.Row{}
	for _, task := range m.tasks {
		rows = append(rows, tbl.Row{
			task.Id,
			task.Name,
			task.Status,
			task.DueDate,
			task.CompletedDate,
		})
	}

	s := tbl.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("#513B56")).
		Bold(false)

	table.SetStyles(s)
	table.SetRows(rows)
	table.SetCursor(m.cursor)

	return baseStyle.Render(table.View())
}

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "esc":
// 			if m.table.Focused() {
// 				m.table.Blur()
// 			} else {
// 				m.table.Focus()
// 			}
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		case "enter":
// 			return m, tea.Batch(
// 				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
// 			)
// 		}
// 	}
// 	m.table, cmd = m.table.Update(msg)
// 	return m, cmd
// }

// func (m model) View() string {
// 	return baseStyle.Render(m.table.View()) + "\n"
// }
