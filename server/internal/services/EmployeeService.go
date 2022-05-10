package services

import (
	"errors"
	"github.com/Hermes-Bird/job-accelerator/internal/domain"
	"github.com/Hermes-Bird/job-accelerator/internal/repositories"
)

type EmployeeService interface {
	CreateEmployee(dto domain.CreateEmployeeDto) (*domain.Employee, error)
	GetEmployeeById(id int) (*domain.Employee, error)
	CheckEmployeeCreds(dto domain.LoginEmployeeDto) (int, error)
	UpdateEmployee(id int, dto domain.EmployeeUpdateDto) (*domain.Employee, error)
}

type EmployeeServiceImpl struct {
	empRepo     repositories.EmployeeRepository
	authService AuthService
}

func NewEmployeeService(empRepo repositories.EmployeeRepository, authService AuthService) EmployeeService {
	return &EmployeeServiceImpl{
		authService: authService,
		empRepo:     empRepo,
	}
}

func (s EmployeeServiceImpl) GetEmployeeById(id int) (*domain.Employee, error) {
	return s.empRepo.GetEmployeeById(id)
}

func (s EmployeeServiceImpl) CheckEmployeeCreds(dto domain.LoginEmployeeDto) (int, error) {
	employee, err := s.empRepo.GetEmployeeByEmail(dto.Email)
	if err != nil {
		return 0, err
	}

	testHash, err := s.authService.HashPasswordWithSalt(dto.Password, employee.PasswordSalt)
	if err != nil {
		return 0, err
	}

	if testHash != employee.PasswordHash {
		return 0, errors.New("wrong user credentials")
	}

	return employee.Id, nil
}

func (s EmployeeServiceImpl) CreateEmployee(dto domain.CreateEmployeeDto) (*domain.Employee, error) {
	passwordPayload, err := s.authService.HashPassword(dto.Password)
	newEmployee := domain.Employee{
		PasswordHash: passwordPayload.PasswordHash,
		PasswordSalt: passwordPayload.PasswordSalt,
		Email:        dto.Email,
		Sex:          dto.Sex,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
	}

	err = s.empRepo.SaveEmployee(&newEmployee)

	if err != nil {
		return nil, err
	}

	return &newEmployee, nil
}

func (s EmployeeServiceImpl) UpdateEmployee(id int, dto domain.EmployeeUpdateDto) (*domain.Employee, error) {
	return s.empRepo.UpdateEmployee(id, dto)
}
