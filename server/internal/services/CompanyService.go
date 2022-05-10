package services

import (
	"errors"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/repositories"
)

type CompanyService interface {
	CreateCompany(dto domain.CreateCompanyDto) (*domain.Company, error)
	CheckCompanyCreds(dto domain.LoginCompanyDto) (int, error)
	UpdateCompany(id int, dto domain.UpdateCompanyDto) (*domain.Company, error)
	GetCompanyById(id int) (*domain.Company, error)
}

type CompanyServiceImpl struct {
	companyRepo repositories.CompanyRepository
	authService AuthService
}

func NewCompanyService(companyRepo repositories.CompanyRepository, auth AuthService) *CompanyServiceImpl {
	return &CompanyServiceImpl{
		companyRepo: companyRepo,
		authService: auth,
	}
}

func (c CompanyServiceImpl) CreateCompany(dto domain.CreateCompanyDto) (*domain.Company, error) {
	passwordPayload, err := c.authService.HashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	newCompany := domain.Company{
		Email:        dto.Email,
		CompanyName:  dto.CompanyName,
		PasswordHash: passwordPayload.PasswordHash,
		PasswordSalt: passwordPayload.PasswordSalt,
	}

	company, err := c.companyRepo.CreateCompany(&newCompany)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (c CompanyServiceImpl) CheckCompanyCreds(dto domain.LoginCompanyDto) (int, error) {
	company, err := c.companyRepo.GetCompanyByEmail(dto.Email)
	if err != nil {
		return 0, err
	}

	testHash, err := c.authService.HashPasswordWithSalt(dto.Password, company.PasswordSalt)
	if err != nil {
		return 0, err
	}

	if testHash != company.PasswordHash {
		return 0, errors.New("wrong user credentials")
	}

	return company.Id, nil
}

func (c CompanyServiceImpl) UpdateCompany(id int, dto domain.UpdateCompanyDto) (*domain.Company, error) {
	return c.companyRepo.UpdateCompany(id, dto)
}

func (c CompanyServiceImpl) GetCompanyById(id int) (*domain.Company, error) {
	return c.companyRepo.GetCompanyById(id)
}
