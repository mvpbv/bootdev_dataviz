package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Lesson_Json struct {
	Lesson struct {
		UUID                     string `json:"UUID"`
		Slug                     string `json:"Slug"`
		Type                     string `json:"Type"`
		CourseUUID               string `json:"CourseUUID"`
		CourseTitle              string `json:"CourseTitle"`
		CourseImageURL           string `json:"CourseImageURL"`
		ChapterUUID              string `json:"ChapterUUID"`
		IsFree                   bool   `json:"IsFree"`
		LastMod                  string `json:"LastMod"`
		CompletionType           string `json:"CompletionType"`
		Title                    string `json:"Title"`
		LessonDataCodeCompletion struct {
			Readme       string `json:"Readme"`
			ProgLang     string `json:"ProgLang"`
			StarterFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"StarterFiles"`
			SolutionFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"SolutionFiles"`
			CodeExpectedOutput string `json:"CodeExpectedOutput"`
		} `json:"LessonDataCodeCompletion"`
		LessonDataCodeCompletionSQL struct {
			Readme       string `json:"Readme"`
			ProgLang     string `json:"ProgLang"`
			StarterFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"StarterFiles"`
			SolutionFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"SolutionFiles"`
		} `json:"LessonDataCodeCompletionSQL"`
		LessonDataMultipleChoice struct {
			Readme   string `json:"Readme"`
			Question struct {
				Question string   `json:"Question"`
				Answers  []string `json:"Answers"`
				Answer   string   `json:"Answer"`
			} `json:"Question"`
			ContainsCompleteDir bool `json:"ContainsCompleteDir"`
		} `json:"LessonDataMultipleChoice"`
		LessonDataHTTPTests struct {
			Readme    string `json:"Readme"`
			HTTPTests struct {
				BaseURL  string `json:"BaseURL"`
				Requests []struct {
					Request struct {
						Method    string `json:"Method"`
						Path      string `json:"Path"`
						BasicAuth any    `json:"BasicAuth"`
						BodyJSON  any    `json:"BodyJSON"`
						Headers   any    `json:"Headers"`
						Actions   struct {
							DelayRequestByMs any `json:"DelayRequestByMs"`
						} `json:"Actions"`
					} `json:"Request"`
					Tests []struct {
						StatusCode     int `json:"StatusCode"`
						BodyContains   any `json:"BodyContains"`
						HeadersContain any `json:"HeadersContain"`
						JSONValue      any `json:"JSONValue"`
					} `json:"Tests"`
					ResponseVariables any `json:"ResponseVariables"`
				} `json:"Requests"`
				ContainsCompleteDir bool `json:"ContainsCompleteDir"`
			} `json:"HTTPTests"`
		} `json:"LessonDataHTTPTests"`
		LessonDataGitHubChecks struct {
			Readme       string `json:"Readme"`
			GitHubChecks struct {
				ContainsCompleteDir bool `json:"ContainsCompleteDir"`
				Steps               []struct {
					StepType                 string `json:"StepType"`
					BranchExistsData         any    `json:"BranchExistsData"`
					PRExistsData             any    `json:"PRExistsData"`
					WorkflowStatusData       any    `json:"WorkflowStatusData"`
					WorkflowLogsContainsData any    `json:"WorkflowLogsContainsData"`
					FileContainsData         any    `json:"FileContainsData"`
				} `json:"Steps"`
				ShowcaseRepo bool `json:"ShowcaseRepo"`
			} `json:"GitHubChecks"`
		} `json:"LessonDataGitHubChecks"`
		LessonDataManual struct {
			Readme     string `json:"Readme"`
			ManualData struct {
				ContainsCompleteDir bool `json:"ContainsCompleteDir"`
			} `json:"ManualData"`
		} `json:"LessonDataManual"`
		LessonDataCodeTests struct {
			Readme       string `json:"Readme"`
			ProgLang     string `json:"ProgLang"`
			StarterFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"StarterFiles"`
			SolutionFiles []struct {
				Name       string `json:"Name"`
				Content    string `json:"Content"`
				IsHidden   bool   `json:"IsHidden"`
				IsReadonly bool   `json:"IsReadonly"`
			} `json:"SolutionFiles"`
		} `json:"LessonDataCodeTests"`
		LessonDataTextInput struct {
			Readme        string `json:"Readme"`
			TextInputData struct {
				ContainsAll  any      `json:"ContainsAll"`
				ContainsNone any      `json:"ContainsNone"`
				MatchesOne   []string `json:"MatchesOne"`
			} `json:"TextInputData"`
		} `json:"LessonDataTextInput"`
		LessonDataCLICommand struct {
			Readme         string `json:"Readme"`
			CLICommandData struct {
				ContainsCompleteDir bool `json:"ContainsCompleteDir"`
				Commands            []struct {
					Command string `json:"Command"`
					Tests   []struct {
						ExitCode           int `json:"ExitCode"`
						StdoutContainsAll  any `json:"StdoutContainsAll"`
						StdoutContainsNone any `json:"StdoutContainsNone"`
						StdoutLinesGt      any `json:"StdoutLinesGt"`
					} `json:"Tests"`
				} `json:"Commands"`
			} `json:"CLICommandData"`
		} `json:"LessonDataCLICommand"`
	} `json:"Lesson"`
	LessonDifficulty int `json:"LessonDifficulty"`
}

type Analytics_Lesson struct {
	UUID             string `json:"UUID"`
	Title            string `json:"Title"`
	Lesson_Index     int    `json:"ChapterIndex"`
	Difficulty       int    `json:"Difficulty"`
	Optional         bool   `json:"CompletionType"`
	Course_Title     string `json:"CourseTitle"`
	Readme_Len       int    `json:"ReadmeLen"`
	Solution_Len     int    `json:"SolutionLen"`
	Programming_Lang string `json:"ProgrammingLang"`
	Chapter          string `json:"ChapterTitle"`
	Type             string `json:"Type"`
	Chapter_Index    int    `json:"ChapterId"`
}

func gather_lesson_details(l *Analytics_Lesson, path string) {
	url := path + l.UUID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("error creating http request: %v\n", err)
		return

	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %v\n", err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error reading response body: %v\n", err)
		return
	}
	obj := Lesson_Json{}
	err = json.Unmarshal(body, &obj)
	if err != nil {
		fmt.Printf("error unmarshalling json: %v\n", err)
		return
	}
	good_title := strings.ReplaceAll(obj.Lesson.Title, ",", ";")
	l.Title = good_title
	good_index := strings.Split(obj.Lesson.Slug, "-")[0]
	l.Lesson_Index, err = strconv.Atoi(good_index)
	if err != nil {
		fmt.Printf("error converting slug to index: %v\n", err)
		return
	}
	l.Difficulty = obj.LessonDifficulty
	l.Difficulty = obj.LessonDifficulty
	if obj.Lesson.CompletionType == "completion_type_standard" {
		l.Optional = false
	} else if obj.Lesson.CompletionType == "completion_type_challenge" {
		l.Optional = true
	} else {
		fmt.Println("Found New Completion Type with UUID: ", l.UUID)
	}
	l.Course_Title = obj.Lesson.CourseTitle
	if obj.Lesson.Type == "type_code_tests" {
		l.Readme_Len = len(obj.Lesson.LessonDataCodeTests.Readme)
		l.Solution_Len = len(obj.Lesson.LessonDataCodeTests.SolutionFiles)
		l.Programming_Lang = obj.Lesson.LessonDataCodeTests.ProgLang
	} else if obj.Lesson.Type == "type_code" {
		l.Readme_Len = len(obj.Lesson.LessonDataCodeCompletion.Readme)
		l.Solution_Len = len(obj.Lesson.LessonDataCodeCompletion.SolutionFiles)
		l.Programming_Lang = obj.Lesson.LessonDataCodeCompletion.ProgLang
	} else if obj.Lesson.Type == "type_code_sql" {
		l.Readme_Len = len(obj.Lesson.LessonDataCodeCompletionSQL.Readme)
		l.Solution_Len = len(obj.Lesson.LessonDataCodeCompletionSQL.SolutionFiles)
		l.Programming_Lang = obj.Lesson.LessonDataCodeCompletionSQL.ProgLang
	} else if obj.Lesson.Type == "type_manual" {
		l.Readme_Len = len(obj.Lesson.LessonDataManual.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else if obj.Lesson.Type == "type_cli_command" {
		l.Readme_Len = len(obj.Lesson.LessonDataCLICommand.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else if obj.Lesson.Type == "type_choice" {
		l.Readme_Len = len(obj.Lesson.LessonDataMultipleChoice.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else if obj.Lesson.Type == "type_http_tests" {
		l.Readme_Len = len(obj.Lesson.LessonDataHTTPTests.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else if obj.Lesson.Type == "type_text_input" {
		l.Readme_Len = len(obj.Lesson.LessonDataTextInput.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else if obj.Lesson.Type == "type_github_checks" {
		l.Readme_Len = len(obj.Lesson.LessonDataGitHubChecks.Readme)
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	} else {
		fmt.Println("Found New Lesson Type with UUID: ", l.UUID)
		l.Readme_Len = 0
		l.Solution_Len = 0
		l.Programming_Lang = "N/A"
	}
	l.Type = obj.Lesson.Type
	file_writer(l)

}
