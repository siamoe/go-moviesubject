package moviesubject

import "fmt"

// Init 初始化主题 只有 reset 为 true 或没有默认主题时才会初始化
func (s *MovieSubject) Init(inits []SubjectInit, reset bool) error {
	var defaultSubjectCount int64
	tx := s.db.Model(&Subject{})
	err := tx.Where("is_default = ?", true).Count(&defaultSubjectCount).Error
	if err != nil {
		return fmt.Errorf("查询默认主题失败: %v", err)
	}
	// 重置或者没有默认主题
	if reset || defaultSubjectCount == 0 {
		// 先清空
		err = tx.Delete(&Subject{}).Error
		if err != nil {
			return fmt.Errorf("清空主题失败: %v", err)
		}
		var subjects []Subject
		for i, init := range inits {
			subjects = append(subjects, Subject{
				SubjectId:    init.SubjectId,
				Name:         init.Name,
				Category:     string(init.Category),
				IsDefault:    true,
				DisplayOrder: i,
			})
		}
		// 初始化
		err = tx.CreateInBatches(subjects, 100).Error
		if err != nil {
			return fmt.Errorf("初始化主题失败: %v", err)
		}
	}
	return nil
}
