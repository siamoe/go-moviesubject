package moviesubject

import (
	"fmt"
)

func (s *MovieSubject) ResetOrder() error {
	var defaultSubjects []Subject
	err := s.db.Model(&Subject{}).Where("is_default = ?", true).Find(&defaultSubjects).Error
	if err != nil {
		return fmt.Errorf("查询默认主题失败: %v", err)
	}
	for i, subject := range defaultSubjects {
		subject.DisplayOrder = i
		err = s.db.Model(subject).Save(subject).Error
		if err != nil {
			return fmt.Errorf("重置排序失败: %v", err)
		}
	}
	return nil
}
