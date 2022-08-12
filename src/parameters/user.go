package parameters

type UserParameters struct {
	FullName string `json:"name" form:"name" validator:"required"`
	Email    string `json:"email" form:"email" validator:"required,email"`
	Password string `json:"password" form:"password" validator:"required"`
}

type LoginParameters struct {
	Email    string `json:"email" form:"email" validator:"required,email"`
	Password string `json:"password" form:"password" validator:"required"`
}
