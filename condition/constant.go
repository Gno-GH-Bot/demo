package condition

import (
	"github.com/GnoCheckBot/demo/utils"

	"github.com/google/go-github/v66/github"
	"github.com/xlab/treeprint"
)

// Always Condition
type always struct{}

var _ Condition = &always{}

func (*always) IsMet(_ *github.PullRequest, details treeprint.Tree) bool {
	return utils.AddStatusNode(true, "On every pull request", details)
}

func Always() Condition {
	return &always{}
}

// Never Condition
type never struct{}

var _ Condition = &never{}

func (*never) IsMet(_ *github.PullRequest, details treeprint.Tree) bool {
	return utils.AddStatusNode(false, "On no pull request", details)
}

func Never() Condition {
	return &never{}
}
