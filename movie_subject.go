package moviesubject

import (
	"errors"
	"fmt"
	"github.com/heibizi/go-douban"
	"gorm.io/gorm"
	"net/url"
	"path"
	"strings"
)

type MovieSubject struct {
	db *gorm.DB
}

func NewMovieSubject(db *gorm.DB) *MovieSubject {
	return &MovieSubject{
		db: db,
	}
}

// Init 初始化主题 只有 reset 为 true 或没有默认主题时才会初始化
func (s *MovieSubject) Init(req InitRequest) error {
	var defaultSubjectCount int64
	tx := s.db.Model(&Subject{})
	err := tx.Where("is_default = ?", true).Count(&defaultSubjectCount).Error
	if err != nil {
		return fmt.Errorf("查询默认主题失败: %v", err)
	}
	// 重置或者没有默认主题
	if req.Reset || defaultSubjectCount == 0 {
		// 先清空
		err = tx.Delete(&Subject{}).Error
		if err != nil {
			return fmt.Errorf("清空主题失败: %v", err)
		}
		var subjects []Subject
		for i, init := range req.Subjects {
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

func (s *MovieSubject) List() ([]Subject, error) {
	var subjects []Subject
	err := s.db.Find(&subjects).Error
	if err != nil {
		return nil, fmt.Errorf("获取主题列表失败: %v", err)
	}
	return subjects, nil
}

func (s *MovieSubject) Order(req OrderRequest) error {
	for i, id := range req.Ids {
		err := s.db.Model(&Subject{}).Where("id = ?", id).Update("display_order", i).Error
		if err != nil {
			return fmt.Errorf("排序失败: %v", err)
		}
	}
	return nil
}

func (s *MovieSubject) Add(req AddRequest) error {
	URL, err := url.Parse(req.SubjectIdOrUrl)
	if err == nil {
		req.SubjectIdOrUrl = path.Base(URL.Path)
	}
	err = s.db.Model(&Subject{}).Create(&Subject{
		SubjectId:    req.SubjectIdOrUrl,
		Name:         req.Name,
		Category:     string(req.Category),
		IsDefault:    false,
		DisplayOrder: req.Order,
	}).Error
	if err != nil {
		return fmt.Errorf("添加主题失败: %v", err)
	}
	return nil
}

func (s *MovieSubject) Delete(req DeleteRequest) error {
	err := s.db.Where("id = ?", req.Id).Delete(&Subject{}).Error
	if err != nil {
		return fmt.Errorf("删除主题失败: %v", err)
	}
	return nil
}

func (s *MovieSubject) Rename(req RenameRequest) error {
	err := s.db.Model(&Subject{}).Where("id = ?", req.Id).Update("name", req.Name).Error
	if err != nil {
		return fmt.Errorf("重命名失败: %v", err)
	}
	return nil
}

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

func (s *MovieSubject) Items(req ItemsRequest) (result ItemResult, err error) {
	if req.Id != nil {
		return s.items(*req.Id, req.Category, req.PageIndex, req.PageSize)
	}
	if req.SubjectId != nil {
		return s.itemsWith(*req.SubjectId, req.Category, req.PageIndex, req.PageSize)
	}
	return result, errors.New("参数错误")
}

func (s *MovieSubject) items(id int64, category Category, pageIndex int, pageSize int) (result ItemResult, err error) {
	var subject Subject
	err = s.db.Model(&Subject{}).Where("id = ?", id).First(&subject).Error
	if err != nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, fmt.Errorf("获取主题失败: %v", err)
	}
	return s.itemsWith(subject.SubjectId, category, pageIndex, pageSize)
}

func (s *MovieSubject) itemsWith(subjectId string, category Category, pageIndex int, pageSize int) (result ItemResult, err error) {
	c := douban.NewApiClient()
	var items []Item
	if category == DouList {
		r, err := c.DouListItems(subjectId, pageIndex*pageSize, pageSize)
		if err != nil {
			return result, err
		}
		for _, item := range r.Items {
			items = append(items, Item{
				Id:        item.ID,
				Title:     item.Title,
				MediaType: getMediaType(item.Type),
				Year:      strings.TrimSpace(strings.Split(item.Subtitle, "/")[0]),
				Vote:      item.Rating.Value,
				Image:     item.CoverURL,
				Overview:  item.Subtitle,
			})
		}
		result.List = items
		result.Total = r.Total
	} else if category == SubjectCollection {
		r, err := c.SubjectCollectionItems(subjectId, pageIndex*pageSize, pageSize)
		if err != nil {
			return result, err
		}
		for _, item := range r.SubjectCollectionItems {
			items = append(items, Item{
				Id:        item.ID,
				Title:     item.Title,
				MediaType: getMediaType(item.Type),
				Year:      strings.TrimSpace(strings.Split(item.CardSubtitle, "/")[0]),
				Vote:      item.Rating.Value,
				Image:     item.CoverURL,
				Overview:  item.CardSubtitle,
			})
		}
		result.Total = r.Total
	} else {
		return result, fmt.Errorf("未知的分类: %v", category)
	}
	result.PageIndex = pageIndex
	result.PageSize = pageSize
	result.List = items
	return result, err
}

func getMediaType(mediaType string) douban.MediaType {
	if mediaType == douban.Movie.Code {
		return douban.Movie
	}
	return douban.Tv
}
