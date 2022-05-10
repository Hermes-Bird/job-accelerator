package repositories

import "github.com/Hermes-Bird/job-accelerator/internal/domain"

type CommonRepository interface {
	getLanguages() []domain.Language
	getRegions() []domain.Region
}
