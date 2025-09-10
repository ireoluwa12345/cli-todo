package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	noStyle      = lipgloss.NewStyle()
)

type CreateModel struct {
	width       int
	height      int
	inputs      []textinput.Model
	description textarea.Model
	focusIndex  int
}

func (t *Tui) NewCreateModel() CreateModel {
	model := CreateModel{
		width:      0,
		height:     0,
		focusIndex: 0,
	}

	// 4 is the number of fields requiring a textinput in the create form
	model.inputs = make([]textinput.Model, 4)
	for i := range model.inputs {
		ti := textinput.New()
		ti.CharLimit = 156
		ti.Width = 20

		switch i {
		case 0:
			ti.Placeholder = "Task Name"
			ti.Focus()
			ti.PromptStyle = focusedStyle
			ti.TextStyle = focusedStyle
		case 1:
			ti.Placeholder = "Status"
		case 2:
			ti.Placeholder = "Priority (Low, Medium, High)"
		case 3:
			ti.Placeholder = "Due Date"
		default:
			ti.Placeholder = ""
		}

		model.inputs[i] = ti
	}

	model.description = textarea.New()
	model.description.Placeholder = "Description"
	model.description.SetWidth(20)
	model.description.SetHeight(5)

	return model
}

func (m CreateModel) Init() tea.Cmd {
	return nil
}

func (m CreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var (
		cmd    tea.Cmd
		desCmd tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			if m.focusIndex == len(m.inputs) {
				for i := 0; i < len(m.inputs); i++ {
					m.inputs[i].Blur()
					m.inputs[i].PromptStyle = noStyle
					m.inputs[i].TextStyle = noStyle
				}
				m.description.Focus()
			} else {
				for i := 0; i < len(m.inputs); i++ {
					if i == m.focusIndex {
						cmds[i] = m.inputs[i].Focus()
						m.inputs[i].PromptStyle = focusedStyle
						m.inputs[i].TextStyle = focusedStyle
					} else {
						m.inputs[i].Blur()
						m.inputs[i].PromptStyle = noStyle
						m.inputs[i].TextStyle = noStyle
					}
				}
				m.description.Blur()
			}

			return m, tea.Batch(cmds...)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.description.SetWidth(msg.Width - padding)
	}

	cmd = m.updateInputs(msg)
	m.description, desCmd = m.description.Update(msg)

	return m, (tea.Batch(cmd, desCmd))
}

func (m CreateModel) View() string {
	labelStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FF69B4"))

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4"))

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Margin(1, 2).
		Width(m.width - 4)

	output := titleStyle.Render("Create New Task") + "\n\n"

	// Task Name
	output += labelStyle.Render("Task Name") + ":"
	if len(m.inputs) > 0 {
		output += m.inputs[0].View() + "\n\n"
	} else {
		output += "(input not initialized)\n\n"
	}

	// Status
	output += labelStyle.Render("Status") + ":"
	if len(m.inputs) > 1 {
		output += m.inputs[1].View() + "\n\n"
	} else {
		output += "(input not initialized)\n\n"
	}

	// Priority
	output += labelStyle.Render("Priority") + ":"
	if len(m.inputs) > 2 {
		output += m.inputs[2].View() + "\n\n"
	} else {
		output += "(input not initialized)\n\n"
	}

	// Due Date
	output += labelStyle.Render("Due Date") + ":"
	if len(m.inputs) > 3 {
		output += m.inputs[3].View() + "\n\n"
	} else {
		output += "(input not initialized)\n\n"
	}

	// Description
	output += labelStyle.Render("Description") + " \n"
	output += m.description.View() + "\n\n"

	return boxStyle.Render(output)
}

func (m *CreateModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
