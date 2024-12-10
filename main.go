package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	url := "https://api.boot.dev/v1/"
	lessons_url := "static/lessons/"
	courses_url := "courses/"
	//test_lesson := ""
	Courses := make_course()
	for _, course := range Courses {
		obj := get_course(url + courses_url + course.UUID)
		Chapters := obj.Chapters
		var Lesson_Links []Analytics_Lesson
		for _, chapter := range Chapters {
			chap_slug := chapter.Slug
			chapter_id_str := strings.Split(chap_slug, "-")[0]
			chapter_id, err := strconv.Atoi(chapter_id_str)
			if err != nil {
				fmt.Printf("error converting slug to index: %v\n", err)
				return
			}
			for _, lesson := range chapter.RequiredLessons {
				ana_lesson := Analytics_Lesson{
					UUID:          lesson.UUID,
					Chapter:       chapter.Title,
					Chapter_Index: chapter_id,
				}
				Lesson_Links = append(Lesson_Links, ana_lesson)
			}
			for _, lesson := range chapter.OptionalLessons {
				ana_lesson := Analytics_Lesson{
					UUID:          lesson.UUID,
					Chapter:       chapter.Title,
					Chapter_Index: chapter_id,
				}
				Lesson_Links = append(Lesson_Links, ana_lesson)
			}
		}
		for _, lesson := range Lesson_Links {
			gather_lesson_details(&lesson, url+lessons_url)
		}
	}
}
func get_course(url string) Courses {
	obj := Courses{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error creating http request: %v\n", err)
		return obj
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %v\n", err)
		return obj
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return obj
	}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Printf("error unmarshalling json: %v\n", err)
		return obj
	}
	return obj
}
