package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Item struct {
	TitleInfo, Desc string
}

func (i Item) Title() string       { return i.TitleInfo }
func (i Item) Description() string { return i.Desc }
func (i Item) FilterValue() string { return i.TitleInfo }

type model struct {
	list list.Model
}

func NewSelect(items []list.Item) *model {
	l := list.New(items, list.NewDefaultDelegate(), 50, 20)
	l.Title = "Commit Type:"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	return &model{
		list: l,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) GetResult() string {
	return m.list.SelectedItem().FilterValue()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if msg.String() == "enter" {
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return lipgloss.NewStyle().Render(m.list.View())
}
