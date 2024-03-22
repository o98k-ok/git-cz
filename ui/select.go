package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type selectItem struct {
	TitleInfo, Desc, Icon string
}

func (i selectItem) Title() string       { return i.TitleInfo }
func (i selectItem) Description() string { return i.Desc }
func (i selectItem) FilterValue() string { return i.TitleInfo + " " + i.Icon }

type selectModel struct {
	list list.Model
}

func NewSelect(items []list.Item) selectModel {
	l := list.New(items, list.NewDefaultDelegate(), 50, 20)
	l.Title = "Commit Type:"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	return selectModel{
		list: l,
	}
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) GetResult() string {
	return m.list.SelectedItem().FilterValue()
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m selectModel) View() string {
	return lipgloss.NewStyle().Render(m.list.View())
}
