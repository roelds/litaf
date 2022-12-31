package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	const (
		fileperms = 0644
	)
	var (
		file1flag string
		file2flag string
		file3flag string
	)
	lineafter := flag.Bool("n", false, "new line after")
	linebefore := flag.Bool("b", false, "new line before")
	linerpl := flag.Bool("r", false, "replace line")
	flag.StringVar(&file1flag, "f1", "file1.tmp", "file with search line")
	flag.StringVar(&file2flag, "f2", "file2.tmp", "file with action line")
	flag.StringVar(&file3flag, "f3", "data.tmp", "file with lines")
	flag.Parse()

	file1, err := os.ReadFile(file1flag)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert []byte to string & trim trailing newline
	text1 := strings.TrimRight(string(file1), "\n")
	fmt.Println(text1)

	file2, err := os.ReadFile(file2flag)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert []byte to string & trim trailing newline
	text2 := strings.TrimRight(string(file2), "\n")
	fmt.Println(text2)

	input, err := os.ReadFile(file3flag)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	found := 0
	fileContent := ""
	hit := false
	for i, line := range lines {
		if strings.Contains(line, text1) {
			switch {
			case *linebefore:
				hit = true
				fileContent += (text2 + "\n")
				fileContent += line
				fileContent += "\n"
				fmt.Println("b")
				found = i
				fmt.Println(fmt.Sprint(found))
			case *lineafter:
				hit = true
				fileContent += line
				fileContent += "\n"
				fileContent += (text2 + "\n")
				fmt.Println("n")
				found = i
				fmt.Println(fmt.Sprint(found))
			case *linerpl:
				lines[i] = text2
				fmt.Println("r")
				found = i
				fmt.Println(fmt.Sprint(found))
				output := strings.Join(lines, "\n")
				err = os.WriteFile(file3flag, []byte(output), fileperms)
				if err != nil {
					log.Fatalln(err)
				}
				os.Exit(0)
			default:
				log.Fatalln("Match found but invalid action cmd")
			}
		}
		if !(hit) { // keep going if no changes made
			fileContent += line
			fileContent += "\n"
		}
		hit = false
	}
	// Save changes without extra trailing newline
	err = os.WriteFile(file3flag, []byte(strings.TrimRight(fileContent, "\n")), fileperms)
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(0)
}
