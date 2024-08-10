package moviesubject

import (
	"errors"
	"fmt"
	"github.com/heibizi/go-douban"
	"gorm.io/gorm"
	"strings"
)

func (s *MovieSubject) Items(id int64, category Category, pageIndex int, pageSize int) (result ItemResult, err error) {
	var subject Subject
	err = s.db.Model(&Subject{}).Where("id = ?", id).First(&subject).Error
	if err != nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return result, fmt.Errorf("获取主题失败: %v", err)
	}
	return s.ItemsWith(subject.SubjectId, category, pageIndex, pageSize)
}

func (s *MovieSubject) ItemsWith(subjectId string, category Category, pageIndex int, pageSize int) (result ItemResult, err error) {
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
	} else {
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
