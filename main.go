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

	fmt.Println(html)
	fmt.Println("Build page")
}
