package moviesubject

type (
	SubjectInit struct {
		Code     string
		Name     string
		Category Category
	}
	SubjectAdd struct {
		Code     string
		Name     string
		Category Category
		Order    int
	}
)
