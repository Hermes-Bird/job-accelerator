package domain

type KeySkill struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	SkillName string `json:"skill_name" validate:"required"`
}
