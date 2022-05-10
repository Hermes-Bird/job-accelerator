package services

import "github.com/Hermes-Bird/job-accelerator/internal/domain"

type VacancyService interface {
	CreateVacancy(dto domain.VacancyDto) (*domain.Vacancy, error)
	GetVacancyById(id int) (*domain.Vacancy, error)
	UpdateVacancy(id int, dto domain.VacancyDto) (*domain.Vacancy, error)
	DeleteVacancyById(id int) error
}
