package main

type Course struct {
	Name      string
	UUID      string
	Main_Path bool
}
type Courses struct {
	UUID                         string   `json:"UUID"`
	ArchivedAt                   any      `json:"ArchivedAt"`
	Slug                         string   `json:"Slug"`
	Title                        string   `json:"Title"`
	ShortDescription             string   `json:"ShortDescription"`
	Description                  string   `json:"Description"`
	ImageURL                     string   `json:"ImageURL"`
	ThumbnailURL                 string   `json:"ThumbnailURL"`
	Difficulty                   float64  `json:"Difficulty"`
	PrerequisiteCourseUUIDS      []any    `json:"PrerequisiteCourseUUIDS"`
	EstimatedCompletionTimeHours int      `json:"EstimatedCompletionTimeHours"`
	TypeDescription              string   `json:"TypeDescription"`
	LastUpdated                  string   `json:"LastUpdated"`
	SlugAliases                  []any    `json:"SlugAliases"`
	AuthorUUIDs                  []string `json:"AuthorUUIDs"`
	MaintainerUUIDs              []string `json:"MaintainerUUIDs"`
	Alternatives                 struct {
	} `json:"Alternatives"`
	Draft              bool `json:"Draft"`
	NumRequiredLessons int  `json:"NumRequiredLessons"`
	NumOptionalLessons int  `json:"NumOptionalLessons"`
	Chapters           []struct {
		UUID            string `json:"UUID"`
		Slug            string `json:"Slug"`
		Title           string `json:"Title"`
		Description     string `json:"Description"`
		RequiredLessons []struct {
			UUID           string `json:"UUID"`
			Slug           string `json:"Slug"`
			Type           string `json:"Type"`
			CourseUUID     string `json:"CourseUUID"`
			CourseTitle    string `json:"CourseTitle"`
			CourseImageURL string `json:"CourseImageURL"`
			ChapterUUID    string `json:"ChapterUUID"`
			IsFree         bool   `json:"IsFree"`
			LastMod        string `json:"LastMod"`
			CompletionType string `json:"CompletionType"`
			Title          string `json:"Title"`
		} `json:"RequiredLessons"`
		OptionalLessons []struct {
			UUID           string `json:"UUID"`
			Slug           string `json:"Slug"`
			Type           string `json:"Type"`
			CourseUUID     string `json:"CourseUUID"`
			CourseTitle    string `json:"CourseTitle"`
			CourseImageURL string `json:"CourseImageURL"`
			ChapterUUID    string `json:"ChapterUUID"`
			IsFree         bool   `json:"IsFree"`
			LastMod        string `json:"LastMod"`
			CompletionType string `json:"CompletionType"`
			Title          string `json:"Title"`
		} `json:"OptionalLessons"`
		NumRequiredLessons int    `json:"NumRequiredLessons"`
		NumOptionalLessons int    `json:"NumOptionalLessons"`
		CourseUUID         string `json:"CourseUUID"`
	} `json:"Chapters"`
	Language string `json:"Language"`
}

func make_course() map[uint]Course {
	courses := map[uint]Course{
		21: {
			Name:      "Learn Crytography",
			UUID:      "6321ddbf-49eb-4748-9737-6bc12e8bb705",
			Main_Path: false,
		},
		1: {
			Name:      "Learn Python",
			UUID:      "f9a25dfb-3e00-4727-ac78-36de82315355",
			Main_Path: true,
		},
		2: {
			Name:      "Learn Shells and Terminals",
			UUID:      "bc7a07ef-ab87-42ab-80de-e7261f2c58a0",
			Main_Path: true,
		},
		3: {
			Name:      "Learn Git",
			UUID:      "933d6dd0-b21a-488e-8ece-469bbef28652",
			Main_Path: true,
		},
		4: {
			Name:      "Build a Bookbot",
			UUID:      "094fd7d4-ec78-4202-96ca-c5f79fc332d2",
			Main_Path: true,
		},
		5: {
			Name:      "Learn Object Oriented Programming",
			UUID:      "f9a48bbc-d1ff-4388-bf0c-23c6e3c60ae0",
			Main_Path: true,
		},
		6: {
			Name:      "Learn Functional Programming",
			UUID:      "b1459f0c-21eb-41e5-b7f3-562ef69d344c",
			Main_Path: true,
		},
		7: {
			Name:      "Build a Static Site Generator",
			UUID:      "d38e78e9-ae52-458e-8494-ec7ecbdab14f",
			Main_Path: true,
		},
		8: {
			Name: "Learn Algorithms",
			UUID: "884342fc-5469-47b4-8125-8bfc897428a8",
		},
		9: {
			Name: "Learn Data Structures",
			UUID: "7bbb53ed-2106-4f6b-b885-e7645c2ff9d8",
		},
		10: {
			Name: "Build a Maze Solver",
			UUID: "2b266bb4-2262-49c0-b6d1-75cd8c5e8be8",
		},
		11: {
			Name: "Learn JavaScript",
			UUID: "2af5c197-21eb-48b4-bd90-b0d59adb311e",
		},
		12: {
			Name: "Learn HTTP",
			UUID: "5d804c54-887a-4c1c-b8c7-b6436f3a132e",
		},
		13: {
			Name: "Build a Web Scraper",
			UUID: "59fbb2aa-7d67-4e88-bac8-42f49798a9f5",
		},
		14: {
			Name: "Learn Go",
			UUID: "3b39d0f6-f944-4f1b-832d-a1daba32eda4",
		},
		15: {
			Name: "Build a Pokedex",
			UUID: "b6ac3462-d76f-453b-bd5d-5d7fe07cdadb",
		},
		16: {
			Name: "Learn Webservers",
			UUID: "81b7293c-60aa-40c7-a158-7c87428f6031",
		},
		17: {
			Name: "Learn SQL",
			UUID: "bc0dc34b-025a-4d97-b7a0-382aa21533aa",
		},
		18: {
			Name:      "Build an RSS Aggregator",
			UUID:      "3a8d6b13-c064-424d-bd09-5e09ceaddfea",
			Main_Path: true,
		},
		19: {
			Name:      "Learn Docker",
			UUID:      "2d740eb6-3234-419e-9a23-08ec9e9889b7",
			Main_Path: true,
		},
		20: {
			Name:      "Learn CI-CD",
			UUID:      "1d057329-0a4a-486f-a8cc-2fccada6b307",
			Main_Path: true,
		},
		22: {
			Name:      "Learn Advanced Algorithms",
			UUID:      "aaad49fb-0dc5-43c6-992c-96d3f83ee663",
			Main_Path: false,
		},
		23: {
			Name:      "Learn Kubernetes",
			UUID:      "6e6236f7-6f6b-45e3-859a-5fd0084754aa",
			Main_Path: false,
		},
		24: {
			Name:      "Learn Pub-Sub Arcitecture",
			UUID:      "93174165-cfaf-4201-a5b6-7da1864c9792",
			Main_Path: false,
		},
		25: {
			Name:      "Learn Git 2",
			UUID:      "bfc68e84-1d2d-461e-a8cd-ce94182d8731",
			Main_Path: false,
		},
	}

	return courses
}
