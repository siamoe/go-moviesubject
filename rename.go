package moviesubject

import "fmt"

func (s *MovieSubject) Rename(id int64, name string) error {
	err := s.db.Model(&Subject{}).Where("id = ?", id).Update("name", name).Error
	if err != nil {
		return fmt.Errorf("重命名失败: %v", err)
	}
	return nil
}
