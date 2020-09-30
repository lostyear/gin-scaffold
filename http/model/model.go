package model

import (
	"fmt"

	"github.com/lostyear/gin-scaffold/commons"
	"gorm.io/gorm"
)

type TestModel struct {
	gorm.Model
	Name string
	Data string
}

func (m *TestModel) GetById(id int) TestModel {
	var data TestModel
	db := commons.GetDatabaseConn()
	db.Take(&data, id)
	return data
}

func (m *TestModel) GetByName(name string) TestModel {
	var data TestModel
	data.Name = name
	db := commons.GetDatabaseConn()
	db.Where(&data).First(&data)
	return data
}

func (m *TestModel) GetByInfo() TestModel {
	var data TestModel
	db := commons.GetDatabaseConn()
	db.Where(m).First(&data)
	// m=&data
	return data
}

func (m *TestModel) GetList() []TestModel {
	var data []TestModel
	db := commons.GetDatabaseConn()
	db.Where(m).Find(&data)
	return data
}

func (m *TestModel) Exist() bool {
	var count int64
	db := commons.GetDatabaseConn()
	result := db.Where(m).Count(&count)
	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return count > 0
}

func (m *TestModel) New() error {
	db := commons.GetDatabaseConn()
	result := db.Create(m)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected != 1 {
		return fmt.Errorf(
			"insert %#v failed with affectrd rows are %d",
			m,
			result.RowsAffected,
		)
	}
	return nil
}

func (m *TestModel) Update() error {
	db := commons.GetDatabaseConn()
	db.Save(m)
	return nil
}

func (m *TestModel) Delete() error {
	db := commons.GetDatabaseConn()
	db.Delete(m, m.ID)
	return nil
}
