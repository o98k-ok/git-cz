package main

import "github.com/o98k-ok/git-cz/git"

func main() {
	// model := ui.NewSelect([]list.Item{
	// 	ui.Item{TitleInfo: "xx", Desc: "xx is xx"},
	// 	ui.Item{TitleInfo: "yy", Desc: "yy is yy"},
	// 	ui.Item{TitleInfo: "zz", Desc: "zz is zz"},
	// })
	// tea.NewProgram(model, tea.WithAltScreen()).Run()

	// fmt.Println(model.GetResult())

	// items := []ui.InputItem{
	// 	{Name: "Scope", Limit: 20},
	// 	{Name: "Summary", Limit: 100},
	// 	{Name: "Body", Limit: 200},
	// }

	// m := ui.NewInputModel(items)
	// tea.NewProgram(m).Run()
	// fmt.Println(m.GetResult())

	flow := git.NewFlow()
	flow.CreateBranchWithCommit(git.BranchMsg{Type: "feature", Branch: "test_it"}, git.CommitMsg{
		Type:    "feature",
		Icon:    "okk",
		Scope:   "lucci",
		Summary: "test it",
		Body:    "test it ok",
	})
}
