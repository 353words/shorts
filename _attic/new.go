package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"time"
)

var (
	goFileRe = regexp.MustCompile(`^[a-z_]+.go`)
)

func main() {
	file, err := os.Open("_attic/TODO.md")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	maxTime := time.Time{}

	const gemsDir = "/home/miki/work/gems/go"
	s := bufio.NewScanner(file)
	for s.Scan() {
		if goFileRe.MatchString(s.Text()) {
			path := fmt.Sprintf("%s/%s", gemsDir, s.Text())
			info, err := os.Stat(path)
			if err != nil {
				log.Printf("warning: %q - %s", path, err)
				continue
			}
			if t := info.ModTime(); t.After(maxTime) {
				maxTime = t
			}
		}
	}
	if err := s.Err(); err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(maxTime)
	matches, err := filepath.Glob(fmt.Sprintf("%s/*.go", gemsDir))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	var newGems []string
	for _, path := range matches {
		info, err := os.Stat(path)
		if err != nil {
			log.Printf("warning: %q - %s", path, err)
			continue
		}
		if t := info.ModTime(); t.After(maxTime) {
			newGems = append(newGems, filepath.Base(path))
		}
	}
	sort.Strings(newGems)
	for _, name := range newGems {
		fmt.Println(name)
	}
}
