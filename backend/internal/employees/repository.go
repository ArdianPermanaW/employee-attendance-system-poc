package employees

import "gorm.io/gorm"

type Repository interface {
	Create(emp *Employee) error
	FindAll() ([]Employee, error)
	FindByID(id uint) (*Employee, error)
	Update(emp *Employee) error
	Delete(id uint) error
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db}
}

func (r *repo) Create(emp *Employee) error {
	return r.db.Create(emp).Error
}

func (r *repo) FindAll() ([]Employee, error) {
	var emps []Employee
	err := r.db.Find(&emps).Error
	return emps, err
}

func (r *repo) FindByID(id uint) (*Employee, error) {
	var emp Employee
	err := r.db.First(&emp, id).Error
	return &emp, err
}

func (r *repo) Update(emp *Employee) error {
	return r.db.Save(emp).Error
}

func (r *repo) Delete(id uint) error {
	return r.db.Delete(&Employee{}, id).Error
}
