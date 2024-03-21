package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/o98k-ok/git-cz/config"
)

const (
	SelectMode string = "select"
	InputMode  string = "input"
)

type CommitMoldel struct {
	input tea.Model
	sel   tea.Model
	mode  string
}

func NewCommitModel(c *config.Config) *CommitMoldel {
	var sels []list.Item = make([]list.Item, 0, len(c.Types))
	for _, i := range c.Types {
		sels = append(sels, selectItem{
			TitleInfo: i.Name,
			Desc:      i.Desc,
			Icon:      i.Icon,
		})
	}

	var inputs []inputItem = make([]inputItem, 0, 1)
	if c.Summary == nil {
		c.Summary = &config.ConfigElem{Size: c.Summary.Size, Name: c.Summary.Name}
	}
	// if c.Branch != nil {
	// 	inputs = append(inputs, inputItem{
	// 		Name:  c.Branch.Name,
	// 		Limit: c.Branch.Size,
	// 	})
	// }

	if c.Scope != nil {
		inputs = append(inputs, inputItem{
			Name:  c.Scope.Name,
			Limit: c.Scope.Size,
		})
	}
	inputs = append(inputs, inputItem{
		Name:  c.Summary.Name,
		Limit: c.Summary.Size,
	})
	if c.Body != nil {
		inputs = append(inputs, inputItem{
			Name:  c.Body.Name,
			Limit: c.Body.Size,
		})
	}

	return &CommitMoldel{
		mode:  SelectMode,
		input: NewInputModel(inputs),
		sel:   NewSelect(sels),
	}
}

func (c *CommitMoldel) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, c.sel.Init())
	cmds = append(cmds, c.input.Init())
	return tea.Batch(cmds...)
}

func (c *CommitMoldel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch c.mode {
	case SelectMode:
		c.sel, cmd = c.sel.Update(msg)
		if v, ok := msg.(tea.KeyMsg); ok {
			if v.String() == "enter" {
				c.mode = InputMode
				return c, nil
			}
		}
		return c, cmd
	case InputMode:
		c.input, cmd = c.input.Update(msg)
	}
	return c, cmd
}

func (c *CommitMoldel) View() string {
	var v1, v2 string
	if c.mode == SelectMode {
		v1 = c.sel.View()
	} else {
		if mm, ok := c.sel.(selectModel); ok {
			v1 = fmt.Sprintf("Commit Type: %s\n", mm.GetResult())
		}
		v2 = c.input.View()
	}
	return lipgloss.NewStyle().Render(v1, "\n", v2)
}

func (c *CommitMoldel) GetResult(mode string) interface{} {
	switch mode {
	case SelectMode:
		return c.sel.(selectModel).GetResult()
	case InputMode:
		return c.input.(inputModel).GetResult()
	}
	return nil
}
