package moviesubject

import "fmt"

func (s *MovieSubject) ResetOrder() error {
	var defaultSubjects []Subject
	tx := s.db.Model(&Subject{})
	err := tx.Where("is_default = ?", true).Find(&defaultSubjects).Error
	if err != nil {
		return fmt.Errorf("查询默认主题失败: %v", err)
	}
	for i, subject := range defaultSubjects {
		err = tx.Where("id = ?", subject.Id).Update("display_order", i).Error
		if err != nil {
			return fmt.Errorf("重置排序失败: %v", err)
		}
	}
	return nil
}
