package repositories

import (
	"errors"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"gorm.io/gorm"
)

type CompanyRepository interface {
	CreateCompany(company *domain.Company) (*domain.Company, error)
	GetCompanyByEmail(email string) (*domain.Company, error)
	GetCompanyById(id int) (*domain.Company, error)
	UpdateCompany(id int, dto domain.UpdateCompanyDto) (*domain.Company, error)
}

type CompanyRepositoryDb struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &CompanyRepositoryDb{db: db}
}

func (r CompanyRepositoryDb) CreateCompany(company *domain.Company) (*domain.Company, error) {
	res := r.db.Create(company)
	if res.Error != nil {
		return nil, res.Error
	}
	return company, nil
}

func (r CompanyRepositoryDb) GetCompanyByEmail(email string) (*domain.Company, error) {
	company := domain.Company{}

	res := r.db.Find(&company, "email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("specified company not found")
	}

	return &company, nil
}

func (r CompanyRepositoryDb) GetCompanyById(id int) (*domain.Company, error) {
	var company = domain.Company{Id: id}

	res := r.db.Find(&company)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("company not found")
	}

	return &company, nil
}

func (r CompanyRepositoryDb) UpdateCompany(id int, dto domain.UpdateCompanyDto) (*domain.Company, error) {
	company, err := r.GetCompanyById(id)
	if err != nil {
		return nil, err
	}

	if dto.Contacts != nil {
		company.Contacts = *dto.Contacts
	}

	if dto.Description != nil {
		company.Description = *dto.Description
	}

	if dto.LogoUrl != nil {
		company.LogoUrl = *dto.LogoUrl
	}

	if dto.CompanySize != nil {
		company.CompanySize = *dto.CompanySize
	}

	res := r.db.Save(&company)

	if res.Error != nil {
		return nil, res.Error
	}

	return company, err
}
