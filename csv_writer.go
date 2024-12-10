package main

import (
	"fmt"
	"os"
	"time"
)

func file_writer(l *Analytics_Lesson) {

	path := "output/lessons.csv"
	line := fmt.Sprintf("%s, %s,%s,%d,%t,%s,%d,%d,%s,%s,%s,%d\n",
		time.Now().UTC().Format("2006-01-02 15:04:05"),
		l.UUID,
		l.Title,
		l.Difficulty,
		l.Optional,
		l.Course_Title,
		l.Readme_Len,
		l.Solution_Len,
		l.Programming_Lang,
		l.Chapter,
		l.Type,
		l.Chapter_Index)
	if !file_exists(path) {
		create_file(line, path)
	} else {
		write_line(line, path)
	}
}
func file_exists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func write_line(line, path string) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(line)
	if err != nil {
		fmt.Println(err)
	}
}

func create_file(line, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	file.Close()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString("time,UUID,Title,Difficulty,Optional,Course,Readme_Len,Solution_Len,Programming_Lang,Chapter,Type\n")
	if err != nil {
		fmt.Println(err)
	}
	_, err = f.WriteString(line)
	if err != nil {
		fmt.Println(err)
	}
}
