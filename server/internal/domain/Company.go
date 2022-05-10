package domain

type Company struct {
	Id           int    `json:"id,omitempty" gorm:"primaryKey"`
	Email        string `json:"email" gorm:"unique"`
	CompanyName  string `json:"company_name" gorm:"unique"`
	PasswordHash string `json:"-"`
	PasswordSalt string `json:"-"`
	CompanySize  string `json:"company_size"`
	LogoUrl      string `json:"logo_url,omitempty"`
	Contacts     string `json:"contacts,omitempty"`
	Description  string `json:"description,omitempty"`
	//Vacancies    []string `json:"vacancies,omitempty" gorm:"foreignKey:CompanyId"`
}

type CreateCompanyDto struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
}

type LoginCompanyDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateCompanyDto struct {
	LogoUrl     *string `json:"logo_url,omitempty"`
	Contacts    *string `json:"contacts,omitempty"`
	Description *string `json:"description,omitempty"`
	CompanySize *string `json:"company_size,omitempty"`
}
