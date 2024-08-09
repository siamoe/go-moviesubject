package moviesubject_test

import (
	"github.com/heibizi/go-moviesubject"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testing"
)

var movieSubject *moviesubject.MovieSubject

func TestMain(m *testing.M) {
	dsn := os.Getenv("GO_MOVIESUBJECT_DSN")
	// 连接到数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	movieSubject = moviesubject.NewMovieSubject(db)
	m.Run()
}

func TestInit(t *testing.T) {
	err := movieSubject.Init([]moviesubject.SubjectInit{{
		Code:     "4026601",
		Name:     "英国历史",
		Category: moviesubject.DouList,
	}}, true)
	if err != nil {
		t.Errorf("初始化失败: %v", err)
	}
}

func TestList(t *testing.T) {
	subjects, err := movieSubject.List()
	if err != nil {
		t.Error(err)
	}
	t.Logf("主题列表: %v", subjects)
}

func TestOrder(t *testing.T) {
	err := movieSubject.Order([]int64{1, 2, 3, 4, 5})
	if err != nil {
		t.Error(err)
	}
}

func TestAdd(t *testing.T) {
	err := movieSubject.Add(moviesubject.SubjectAdd{
		Code:     "4026601",
		Name:     "英国历史",
		Category: moviesubject.DouList,
		Order:    6,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	err := movieSubject.Delete(4)
	if err != nil {
		t.Error(err)
	}
}

func TestRename(t *testing.T) {
	err := movieSubject.Rename(7, "英国历史_new")
	if err != nil {
		t.Error(err)
	}
}

func TestResetOrder(t *testing.T) {
	err := movieSubject.ResetOrder()
	if err != nil {
		t.Error(err)
	}
}
