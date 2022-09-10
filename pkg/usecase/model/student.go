package model

import (
	"time"

	"github.com/yagikota/clean_architecture_wtih_go/pkg/domain/model"
)

type StudentID int

type StudentSlice []*Student

type Student struct {
	ID    StudentID `json:"id"`
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Class int       `json:"class"`
}

func StudentFromDomainModel(m *model.Student) *Student {
	s := &Student{
		ID:    StudentID(m.ID),
		Name:  m.Name,
		Class: m.Class,
	}

	age, err := calcAge(m.Birthday, time.Now())
	if err != nil {
		return nil
	}
	s.Age = age

	return s
}

// 誕生日から年齢を計算
func calcAge(birthday, currTime time.Time) (int, error) {
	// タイムゾーンをJSTに設定
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := currTime.In(jst)

	thisYear, thisMonth, thisDay := nowJST.Date()
	age := thisYear - birthday.Year()

	// 誕生日を迎えていない時の処理
	if thisMonth < birthday.Month() || thisMonth == birthday.Month() && thisDay < birthday.Day() {
		age -= 1
	}

	return age, nil
}
