package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("No file provided")
		return
	}

	template, err := os.ReadFile(args[0])
	content, err := os.ReadFile(args[1])

	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}

	filePath := args[1]
	fileNameWithExt := filepath.Base(filePath)
	ext := filepath.Ext(fileNameWithExt)
	fileName := strings.TrimSuffix(fileNameWithExt, ext)

	html := strings.Replace(string(template), "{{.Body}}", string(content), 1)
	html = strings.Replace(html, "{{.Title}}", fileName, 1)

	err = os.MkdirAll("out", 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("./out/"+fileName+".html", []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Build page complete")
}
