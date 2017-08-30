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

	branch := branch(r)
	log.Printf("Branch: %v\n", branch)
	os.Setenv("GIT_BRANCH", branch)

	/*
		GIT_BRANCH=feature/NETSIP-1088-sip-usernames
		+update_current_git_vars:7> GIT_AHEAD=1
		+update_current_git_vars:8> GIT_BEHIND=0
		+update_current_git_vars:9> GIT_STAGED=0
		+update_current_git_vars:10> GIT_CONFLICTS=0
		+update_current_git_vars:11> GIT_CHANGED=0
		+update_current_git_vars:12> GIT_UNTRACKED=0
	*/

	clean := isClean(r)
	log.Printf("Clean: %v\n", clean)
}

func branch(r *git.Repository) string {
	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	nameRef := ref.Name()
	name := strings.Replace(nameRef.String(), "refs/heads/", "", 1)

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
