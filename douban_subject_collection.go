package moviesubject

func (s *doubanService) subjectCollectionItems(subject Subject, pageIndex int, pageSize int) (Result, error) {
	r, err := s.client.SubjectCollectionItems(subject.Code, pageIndex*pageSize, pageSize)
	if err != nil {
		return Result{}, err
	}
	var list []Media
	for _, item := range r.SubjectCollectionItems {
		list = append(list, Media{
			Id:       item.ID,
			Title:    item.Title,
			Type:     item.Type,
			Year:     getDoubanYearFrom(item.CardSubtitle),
			Vote:     item.Rating.Value,
			Image:    item.CoverURL,
			Overview: item.CardSubtitle,
		})
	}
	return Result{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     r.Total,
		List:      list,
	}, nil
}
