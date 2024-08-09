package moviesubject

import "fmt"

func (s *MovieSubject) Add(subject SubjectAdd) error {
	err := s.db.Model(&Subject{}).Create(&Subject{
		Code:         subject.Code,
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
