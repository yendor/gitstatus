package main

import "log"

func main() {
	r, err := git.PlainOpen(".")

	ref, err := r.Head()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("HEAD: %v\n", ref)
}
