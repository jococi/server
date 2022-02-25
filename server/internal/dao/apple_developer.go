package dao

import (
	"supersign/internal/model"

	"gorm.io/gorm"
)

func newAppleDeveloper(db *gorm.DB) *appleDeveloper {
	return &appleDeveloper{db}
}

type appleDeveloper struct {
	db *gorm.DB
}

var _ model.AppleDeveloperStore = (*appleDeveloper)(nil)

func (a *appleDeveloper) Create(appleDeveloper *model.AppleDeveloper) error {
	return a.db.Create(appleDeveloper).Error
}

func (a *appleDeveloper) Del(iss string) error {
	return a.db.Delete(&model.AppleDeveloper{Iss: iss}).Error
}

func (a *appleDeveloper) AddCount(iss string, num int) error {
	return a.db.Model(&model.AppleDeveloper{Iss: iss}).
		UpdateColumn("count", gorm.Expr("count + ?", num)).Error
}

func (a *appleDeveloper) UpdateCount(iss string, count int) error {
	return a.db.Model(&model.AppleDeveloper{Iss: iss}).
		Update("count", count).Error
}

func (a *appleDeveloper) UpdateLimit(iss string, limit int) error {
	return a.db.Model(&model.AppleDeveloper{Iss: iss}).
		Update("limit", limit).Error
}

func (a *appleDeveloper) Enable(iss string, enable bool) error {
	return a.db.Model(&model.AppleDeveloper{Iss: iss}).
		Update("enable", enable).Error
}

func (a *appleDeveloper) Query(iss string) (*model.AppleDeveloper, error) {
	appleDeveloper := &model.AppleDeveloper{Iss: iss}
	err := a.db.Take(appleDeveloper).Error
	if err != nil {
		return nil, err
	}
	return appleDeveloper, nil
}

func (a *appleDeveloper) GetUsable() (*model.AppleDeveloper, error) {
	var appleDeveloper model.AppleDeveloper
	err := a.db.Where("limit - count > 0 And count < ? And enable = ?", 100, true).
		Take(&appleDeveloper).Error
	if err != nil {
		return nil, err
	}
	return &appleDeveloper, nil
}

func (a *appleDeveloper) List(page, pageSize *int) ([]model.AppleDeveloper, error) {
	var appleDevelopers []model.AppleDeveloper
	err := a.db.Scopes(paginate(page, pageSize)).Find(&appleDevelopers).Error
	if err != nil {
		return nil, err
	}
	return appleDevelopers, nil
}
