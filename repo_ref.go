// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package git

import "strings"

// RefType Ref type
type RefType string

// Ref type const
const (
	REF_COMMIT RefType = "commit"
	REF_TREE   RefType = "tree"
	REF_BLOB   RefType = "blob"
	REF_TAG    RefType = "tag"
)

// Ref a git ref.
type Ref struct {
	ID     string // The ID of this ref object
	Type   string
	Branch Branch
}

// GetRecentRef Get most recent ref
func (repo *Repository) GetRecentRef() (*Ref, error) {
	stdout, err := NewCommand("for-each-ref", "--sort=-committerdate", "--count=1").RunInDir(repo.Path)
	if err != nil {
		return nil, err
	}

	infos := strings.Fields(stdout)

	if len(infos) != 3 {
		// TODO
	}

	b := Branch{Name: infos[2][len(BRANCH_PREFIX):], Path: infos[2]}

	return &Ref{
		ID:     infos[0],
		Type:   infos[1],
		Branch: b,
	}, nil
}
