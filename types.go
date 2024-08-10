package moviesubject

import "github.com/heibizi/go-douban"

type (
	SubjectInit struct {
		SubjectId string   `json:"subjectId"` // 豆列或者片单的 id
		Name      string   `json:"name"`      // 名称
		Category  Category `json:"category"`  // 类别
	}
	SubjectAdd struct {
		SubjectIdOrUrl string   `json:"subjectIdOrUrl"` // 豆列或者片单的 id 或者 url
		Name           string   `json:"name"`           // 名称
		Category       Category `json:"category"`       // 类别
		Order          int      `json:"order"`          // 显示顺序
	}
)

type (
	Item struct {
		Id        string           `json:"id"`
		Title     string           `json:"title"`
		MediaType douban.MediaType `json:"type"`
		Year      string           `json:"year"`
		Vote      float64          `json:"vote"`
		Image     string           `json:"image"`
		Overview  string           `json:"overview"`
	}
	ItemResult struct {
		PageIndex int    `json:"pageIndex"`
		PageSize  int    `json:"pageSize"`
		Total     int    `json:"total"`
		List      []Item `json:"items"`
	}
)
