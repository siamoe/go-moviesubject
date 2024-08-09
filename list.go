package moviesubject

import "fmt"

func (s *MovieSubject) List() ([]Subject, error) {
	var subjects []Subject
	err := s.db.Find(&subjects).Error
	if err != nil {
		return nil, fmt.Errorf("获取主题列表失败: %v", err)
	}
	return subjects, nil
}
