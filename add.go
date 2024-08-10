package moviesubject

import (
	"fmt"
	"net/url"
	"path"
)

func (s *MovieSubject) Add(subject SubjectAdd) error {
	URL, err := url.Parse(subject.SubjectIdOrUrl)
	if err == nil {
		subject.SubjectIdOrUrl = path.Base(URL.Path)
	}
	err = s.db.Model(&Subject{}).Create(&Subject{
		SubjectId:    subject.SubjectIdOrUrl,
		Name:         subject.Name,
		Category:     string(subject.Category),
		IsDefault:    false,
		DisplayOrder: subject.Order,
	}).Error
	if err != nil {
		return fmt.Errorf("添加主题失败: %v", err)
	}
	return nil
}
