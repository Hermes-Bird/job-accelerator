package domain

type WorkExperience = string

const (
	NoWorkExperience        WorkExperience = "0"
	OneYearExperience       WorkExperience = "1y"
	TwoYearsExperience      WorkExperience = "2y"
	ThreeYearsExperience    WorkExperience = "3y"
	FourYearsExperience     WorkExperience = "4y"
	FiveYearsExperience     WorkExperience = "5y"
	FivePlusYearsExperience WorkExperience = "5y+"
)

type Vacancy struct {
	Id                 int            `json:"id" gorm:"primaryKey"`
	Description        string         `json:"description,omitempty"`
	Salary             int            `json:"salary,omitempty"`
	RequiredExperience WorkExperience `json:"required_experience,omitempty" gorm:"type:ENUM('0y', '1y', '2y', '3y', '4y', '5y', '5y+')"`
	KeySkills          []KeySkill     `json:"key_skills" gorm:"many2many:vacancy_skill"`
	RegionId           int            `json:"region_id"`
	Region             Region         `json:"region" gorm:"foreignKey:RegionId"`
	CompanyId          int            `json:"company_id"`
	Company            Company        `json:"company" gorm:"foreignKey:CompanyId"`
}

type VacancyDto struct {
	Description        *string         `json:"description"`
	Salary             *int            `json:"salary"`
	RequiredExperience *WorkExperience `json:"required_experience"`
	KeySkills          []KeySkill      `json:"key_skills"`
	RegionId           *int            `json:"region_id"`
}
