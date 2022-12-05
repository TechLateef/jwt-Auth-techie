package dto

type CreateUserDto struct {
	Id    uint64 `uri:"id" form:"id" binding:"required,uuid"`
	Name  string `uri:"name" form:"name" binding:"required"`
	Email string `uri:"email" form:"email" binding:"required"`
	Phone string `uri:"->;<-;not null" form:"phone" binding:"required"`
}

type UpdateUserDto struct {
	Id       uint64 `uri:"id" form:"id" binding:"required,uuid"`
	Name     string `uri:"name" form:"name" binding:"required"`
	Email    string `uri:"email" form:"email" binding:"required"`
	UserID   uint64 `json:"user_id,omitempty" form:"user_id,omitempty"`
	Password string `uri:"password, omitempty" form:"password, omitempty"`
}
