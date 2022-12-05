package entity

type Users struct {
	Id    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Name  string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Phone string `gorm:"->;<-;not null" json:"-"`
}
