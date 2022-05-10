package domain

type Region struct {
	Id         int    `json:"id,omitempty" gorm:"primaryKey"`
	RegionName string `json:"region_name"`
}
