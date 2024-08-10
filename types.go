package moviesubject

import "github.com/heibizi/go-douban"

type (
	SubjectInit struct {
		SubjectId string   `json:"subjectId"` // 豆列或者片单的 id
		Name      string   `json:"name"`      // 名称
		Category  Category `json:"category"`  // 类别
	}
)

type (
	InitRequest struct {
		Subjects []SubjectInit `json:"subjects"` // 主题列表
		Reset    bool          `json:"reset"`    // 是否重置
	}
	OrderRequest struct {
		Ids []int64 `json:"ids"` // 主题 id 列表
	}
	AddRequest struct {
		SubjectIdOrUrl string   `json:"subjectIdOrUrl"` // 豆列或者片单的 id 或者 url
		Name           string   `json:"name"`           // 名称
		Category       Category `json:"category"`       // 类别
		Order          int      `json:"order"`          // 显示顺序
	}
	DeleteRequest struct {
		Id int64 `json:"id"`
	}
	RenameRequest struct {
		Id   int64  `json:"id"`   // 主题 id
		Name string `json:"name"` // 名称
	}
	ItemsRequest struct {
		Id        *int64   `form:"id,optional" json:"id,optional"`               // 主题 id
		SubjectId *string  `form:"subjectId,optional" json:"subjectId,optional"` // 豆列或者片单的 id，与 id 二选一
		Category  Category `form:"category" json:"category"`                     // 类别
		PageIndex int      `form:"pageIndex" json:"pageIndex"`                   // 页码
		PageSize  int      `form:"pageSize" json:"pageSize"`                     // 每页大小
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
