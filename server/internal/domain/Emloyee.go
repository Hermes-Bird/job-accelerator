package domain

import (
	"time"
)

type EmployeeJobDescription struct {
	Id               int       `json:"id,omitempty" gorm:"primaryKey"`
	EmployeeId       int       `json:"user_id,omitempty" gorm:"index"`
	StartDate        time.Time `json:"start_date"`
	EndDate          time.Time `json:"end_date"`
	Organization     string    `json:"organization"`
	Responsibilities string    `json:"responsibilities"`
}

type EmployeeEducation struct {
	Id             int       `json:"id,omitempty" gorm:"primaryKey"`
	EmployeeId     int       `json:"user_id,omitempty" gorm:"index"`
	Specialization string    `json:"specialization,omitempty"`
	EndDate        time.Time `json:"end_date"`
}

type Employee struct {
	Id              int                      `json:"id,omitempty" gorm:"primaryKey"`
	PasswordHash    string                   `json:"-"`
	PasswordSalt    string                   `json:"-"`
	Email           string                   `json:"email" gorm:"unique"`
	BirthDate       time.Time                `json:"birth_date"`
	Sex             string                   `json:"sex"`
	FirstName       string                   `json:"first_name"`
	LastName        string                   `json:"last_name"`
	Contacts        string                   `json:"contacts,omitempty"`
	RegionId        int                      `json:"region_id,omitempty"`
	Region          *Region                  `json:"region,omitempty" gorm:"foreignKey:RegionId"`
	Description     string                   `json:"description"`
	JobDescriptions []EmployeeJobDescription `json:"job_descriptions" gorm:"foreignKey:EmployeeId;constraint:OnUpdate:CASCADE"`
	KeySkills       []KeySkill               `json:"key_skills" gorm:"many2many:employee_skills"`
	Languages       []Language               `json:"languages" gorm:"many2many:employee_languages"`
	Educations      []EmployeeEducation      `json:"educations" gorm:"foreignKey:EmployeeId"`
	Salary          string                   `json:"salary,omitempty"`
}

type CreateEmployeeDto struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Sex       string `json:"sex"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type LoginEmployeeDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type EmployeeUpdateDto struct {
	BirthDate       *time.Time               `json:"birth_date,omitempty"`
	Sex             *string                  `json:"sex,omitempty" validate:"required"`
	FirstName       *string                  `json:"first_name,omitempty"`
	LastName        *string                  `json:"last_name,omitempty"`
	Contacts        *string                  `json:"contacts,omitempty"`
	RegionId        *int                     `json:"region_id,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	KeySkills       []KeySkill               `json:"key_skills" validate:"dive,required"`
	JobDescriptions []EmployeeJobDescription `json:"job_descriptions"`
	Languages       []Language               `json:"languages,omitempty"`
	Educations      []EmployeeEducation      `json:"educations"`
	Salary          *string                  `json:"salary,omitempty"`
}
