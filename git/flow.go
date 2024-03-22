package git

type Flow struct {
	g *Git
}

func NewFlow() *Flow {
	return &Flow{
		g: &Git{},
	}
}

func (f *Flow) CreateBranchWithCommit(branch BranchMsg, msg CommitMsg) {
	f.g.CoBranch(branch).Commit(msg)
}

func (f *Flow) Commit(msg CommitMsg) {
	f.g.Commit(msg)
}
