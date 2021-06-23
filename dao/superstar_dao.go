package dao

import (
	"iris/models"
	"log"

	"github.com/go-xorm/xorm"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(en *xorm.Engine) *SuperstarDao {
	return &SuperstarDao{
		engine: en,
	}
}
func (s *SuperstarDao) Get(id int) *models.StarInfo {
	var userInfo models.StarInfo
	ok, err := s.engine.ID(id).Get(&userInfo)
	if ok && err == nil {
		return &userInfo
	} else {
		userInfo.Id = 0
		return &userInfo
	}
}

func (s *SuperstarDao) GetAll() []models.StarInfo {
	var stars []models.StarInfo
	err := s.engine.Find(&stars)
	if err != nil {
		log.Println(err)
		return stars
	} else {
		return stars
	}
}

func (s *SuperstarDao) Search(country string) []models.StarInfo {
	var stars []models.StarInfo
	err := s.engine.Where("country = ?", country).Desc("id").Find(&stars)
	if err != nil {
		log.Println(err)
		return stars
	} else {
		return stars
	}
}

func (s *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := s.engine.Id(data.Id).Update(data)
	return err
}

func (s *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (s *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := s.engine.Insert(data)
	return err
}
