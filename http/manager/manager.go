package manager

import "github.com/lostyear/gin-scaffold/http/model"

type Manager struct {
	Name string
	Data string
}

func (m *Manager) GetAll() []model.TestModel {
	list := (&model.TestModel{}).GetList()
	return list
}

func (m *Manager) CreateNew() model.TestModel {
	data := model.TestModel{
		Name: m.Name,
		Data: m.Data,
	}
	data.New()
	return data
}

func (m *Manager) Update(id int) model.TestModel {
	data := model.TestModel{
		Name: m.Name,
		Data: m.Data,
	}
	data.ID = uint(id)
	data.Update()
	return data
}

func (m *Manager) Delete(id int) model.TestModel {
	data := model.TestModel{
		Name: m.Name,
		Data: m.Data,
	}
	data.ID = uint(id)
	data.Delete()
	return data
}
