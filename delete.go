package moviesubject

import "fmt"

func (s *MovieSubject) Delete(id int64) error {
	err := s.db.Where("id = ?", id).Delete(&Subject{}).Error
	if err != nil {
		return fmt.Errorf("删除主题失败: %v", err)
	}
	return nil
}
