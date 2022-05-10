package domain

type Language struct {
	Id           int    `json:"id" gorm:"primaryKey"`
	LanguageName string `json:"language_name"`
}
