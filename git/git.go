package git

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/o98k-ok/git-cz/pkg"
)

type Git struct {
	WithSign     bool
	RandomBranch bool
}

type CommitMsg struct {
	Type    string
	Icon    string
	Scope   string
	Summary string
	Body    string
}

type BranchMsg struct {
	Type   string
	Branch string
}

func (g *Git) Add() *Git {
	return g
}

func (g *Git) Commit(msg CommitMsg) *Git {
	if g == nil {
		return nil
	}

	var err error
	var file *os.File
	if file, err = os.CreateTemp("", "COMMIT_MESSAGE_"); err != nil {
		return nil
	}
	defer os.Remove(file.Name())
	if _, err := file.Write([]byte(g.format(msg))); err != nil {
		return nil
	}

	cmd := fmt.Sprintf("git commit -F %s", file.Name())
	stdout, stderr, err := pkg.ExecCommand(cmd)
	fmt.Println("** create commit >>:\n", stdout, stderr)
	if err != nil {
		return nil
	}
	return g
}

func (g *Git) format(msg CommitMsg) string {
	builder := strings.Builder{}
	builder.WriteString("[" + msg.Type + "]")
	if len(msg.Scope) != 0 {
		builder.WriteString("<" + msg.Scope + ">")
	}
	builder.WriteString(": " + msg.Summary + " " + msg.Icon)
	if len(msg.Body) != 0 {
		builder.WriteString("\n\n" + msg.Body)
	}
	builder.WriteString("[" + msg.Type + "]")
	return builder.String()
}

func (g *Git) CoBranch(branch BranchMsg) *Git {
	if g == nil {
		return nil
	}
	cmd := fmt.Sprintf("git co -b %s/%s_%d", branch.Type, branch.Branch, time.Now().Unix())
	stdout, stderr, err := pkg.ExecCommand(cmd)
	fmt.Println("** switch branch >>\n", stdout, stderr)
	if err != nil {
		return nil
	}
	return g
}
