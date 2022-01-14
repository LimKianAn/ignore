package main

import (
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	buildConstraint := "//go:build ignore\n\n"
	for _, e := range entries {
		if !strings.Contains(e.Name(), "_test.go") {
			continue
		}

		bb, err := os.ReadFile(e.Name())
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) == 2 && os.Args[1] == "undo" {
			i := bytes.LastIndex(bb, []byte(buildConstraint))
			if i == -1 {
				continue
			}
			bb = bb[i+len(buildConstraint):]
		}

		if err := os.Truncate(e.Name(), 0); err != nil {
			log.Fatal(err)
		}

		f, err := os.OpenFile(e.Name(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		if len(os.Args) <= 1 {
			if _, err := f.WriteString(buildConstraint); err != nil {
				log.Fatal(err)
			}
		}

		if _, err := f.Write(bb); err != nil {
			log.Fatal(err)
		}

		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}
}
