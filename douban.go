package moviesubject

import (
	"fmt"
	"github.com/heibizi/go-douban"
	"strconv"
	"strings"
)

type doubanService struct {
	client *douban.ApiClient
}

func getDoubanYearFrom(title string) int {
	i, err := strconv.Atoi(strings.TrimSpace(strings.Split(title, "/")[0]))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return i
}
