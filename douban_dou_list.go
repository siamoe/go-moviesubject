package moviesubject

func (s *doubanService) douListItems(subject Subject, pageIndex int, pageSize int) (Result, error) {
	r, err := s.client.DouListItems(subject.Code, pageIndex*pageSize, pageSize)
	if err != nil {
		return Result{}, err
	}
	var list []Media
	for _, item := range r.Items {
		list = append(list, Media{
			Id:       item.ID,
			Title:    item.Title,
			Type:     item.Type,
			Year:     getDoubanYearFrom(item.Subtitle),
			Vote:     item.Rating.Value,
			Image:    item.CoverURL,
			Overview: item.Subtitle,
		})
	}
	return Result{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Total:     r.Total,
		List:      list,
	}, nil
}
