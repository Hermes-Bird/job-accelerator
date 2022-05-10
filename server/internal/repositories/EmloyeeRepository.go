package repositories

import (
	"errors"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type EmployeeRepository interface {
	SaveEmployee(employee *domain.Employee) error
	UpdateEmployee(id int, dto domain.EmployeeUpdateDto) (*domain.Employee, error)
	GetEmployeeByEmail(email string) (*domain.Employee, error)
	GetEmployeeById(id int) (*domain.Employee, error)
}

type EmployeeRepositoryDb struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &EmployeeRepositoryDb{db: db}
}

func (r EmployeeRepositoryDb) GetEmployeeById(id int) (*domain.Employee, error) {
	empl := domain.Employee{}
	res := r.db.Preload(clause.Associations).Find(&empl, "id = ?", id)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("employee not found")
	}

	return &empl, nil
}

func (r EmployeeRepositoryDb) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	empl := domain.Employee{Email: email}
	res := r.db.Find(&empl)

	if res.Error != nil {
		return nil, res.Error
	}
	return &empl, nil
}

func (r EmployeeRepositoryDb) SaveEmployee(employee *domain.Employee) error {
	res := r.db.Omit("Region", "RegionId", "BirthDate").Create(&employee)
	return res.Error
}

func (r EmployeeRepositoryDb) UpdateEmployee(id int, employeeUpdate domain.EmployeeUpdateDto) (*domain.Employee, error) {
	empl := domain.Employee{Id: id}
	res := r.db.Find(&empl)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("employee not found")
	}

	if employeeUpdate.LastName != nil {
		empl.LastName = *employeeUpdate.LastName
	}

	if employeeUpdate.FirstName != nil {
		empl.FirstName = *employeeUpdate.FirstName
	}

	if employeeUpdate.BirthDate != nil {
		empl.BirthDate = *employeeUpdate.BirthDate
	}

	if employeeUpdate.Sex != nil {
		empl.Sex = *employeeUpdate.Sex
	}

	if employeeUpdate.Contacts != nil {
		empl.Contacts = *employeeUpdate.Contacts
	}

	if employeeUpdate.RegionId != nil {
		empl.RegionId = *employeeUpdate.RegionId
	}

	if employeeUpdate.Description != nil {
		empl.Description = *employeeUpdate.Description
	}

	if employeeUpdate.Salary != nil {
		empl.Salary = *employeeUpdate.Salary
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		res := tx.Omit("RegionId").Save(&empl)
		if res.Error != nil {
			return res.Error
		}

		if employeeUpdate.KeySkills != nil {
			err := tx.Model(&empl).Association("KeySkills").Replace(&employeeUpdate.KeySkills)
			if err != nil {
				return err
			}
		}

		if employeeUpdate.Languages != nil {
			err := tx.Model(&empl).Association("Languages").Replace(&employeeUpdate.Languages)
			if err != nil {
				return err
			}
		}

		if employeeUpdate.JobDescriptions != nil {
			err := tx.Model(&empl).Session(&gorm.Session{FullSaveAssociations: true}).Association("JobDescriptions").Replace(&employeeUpdate.JobDescriptions)
			if err != nil {
				return err
			}
		}

		if employeeUpdate.Educations != nil {
			err := tx.Model(&empl).Session(&gorm.Session{FullSaveAssociations: true}).Association("Educations").Replace(&employeeUpdate.Educations)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &empl, nil
}
