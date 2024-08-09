package moviesubject

type (
	Subject struct {
		Id           int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"` // 主键
		Code         string `gorm:"column:code" json:"code"`                      // 编码
		Name         string `gorm:"column:name" json:"name"`                      // 名称
		Category     string `gorm:"column:category" json:"category"`              // 类别
		IsDefault    bool   `gorm:"column:is_default" json:"isDefault"`           // 是否默认
		DisplayOrder int    `gorm:"column:display_order" json:"displayOrder"`     // 显示顺序
	}
)

func (Subject) TableName() string {
	return "movie_subject"
}
