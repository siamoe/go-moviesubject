package moviesubject

import "fmt"

func (s *MovieSubject) Order(ids []int64) error {
	for i, id := range ids {
		err := s.db.Model(&Subject{}).Where("id = ?", id).Update("display_order", i).Error
		if err != nil {
			return fmt.Errorf("排序失败: %v", err)
		}
	}
	return nil
}
