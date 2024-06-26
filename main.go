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

	var outPath string

	if len(args) == 3 {
		outPath = filepath.Join(".", args[2])
	} else {
		outPath = filepath.Join(".", "out")
	}

	splitTitle := strings.Split(fileName, "-")

	for i, word := range splitTitle {
		splitTitle[i] = strings.ToUpper(string(word[0])) + word[1:]
	}

	title := strings.Join(splitTitle, " ")

	html := strings.Replace(string(template), "{{Body}}", string(content), 1)
	html = strings.Replace(html, "{{Title}}", title, 1)

	err = os.MkdirAll(outPath, 0755)

	if err != nil {
		log.Fatal(err)
	}

	filePath = filepath.Join(outPath, fileName+".html")

	err = os.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Build page complete")
}
