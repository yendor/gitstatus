package main

import (
	"fmt"
	"log"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	var r *git.Repository
	var err error

	// Starting at ., search up the directory tree looking for a git repo, stopping at /
	p := "."
	for {
		p, err = filepath.Abs(p)
		if err != nil {
			return
		}
		if p == "/" {
			return
		}
		r, err = git.PlainOpen(p)
		if err != nil {
			p = fmt.Sprintf("%s/../", p)
		} else {
			break
		}
	}

	// Get details of the current git repo
	branch := branch(r)
	ahead := 1
	behind := 1
	staged := 0
	conflicts := 0
	changed := 3
	untracked := 2

	// Display the output formatted for git-prompt
	fmt.Printf(
		"%s %d %d %d %d %d %d",
		branch,
		ahead,
		behind,
		staged,
		conflicts,
		changed,
		untracked,
	)
}

func branch(r *git.Repository) string {
	ref, err := r.Head()
	if err != nil {
		return ""
	}
	name := ref.Name().Short()
	return name
}

func isClean(r *git.Repository) bool {
	wt, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	status, err := wt.Status()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", status)

	return status.IsClean()
}
