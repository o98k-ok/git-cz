package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/o98k-ok/git-cz/config"
	"github.com/o98k-ok/git-cz/git"
	"github.com/o98k-ok/git-cz/ui"
)

func main() {
	cfg := config.NewConfigWith("/Users/shadow/Documents/code/git-cz/config/config.json")
	cmt := ui.NewCommitModel(cfg)
	tea.NewProgram(cmt).Run()

	flow := git.NewFlow()
	types := cmt.GetResult(ui.SelectMode).(string)
	typeInfo := strings.Split(types, " ")
	msg := cmt.GetResult(ui.InputMode).([]string)

	// branchMsg := git.BranchMsg{Type: typeInfo[0], Branch: msg[1]}
	gitMsg := git.CommitMsg{
		Type:    typeInfo[0],
		Icon:    typeInfo[1],
		Scope:   msg[0],
		Summary: msg[1],
	}
	flow.Commit(gitMsg)
}
