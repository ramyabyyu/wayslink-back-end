package authdto

type RegisterRequest struct {
	Email    string `json:"emailregister" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"passwordregister" gorm:"type: varchar(255)" validate:"required"`
	FullName string `json:"fullnameregister" gorm:"type: varchar(255)" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"emaillogin" gorm:"type: varchar(255)" validate:"required"`
	Password string `json:"passwordlogin" gorm:"type: varchar(255)" validate:"required"`
}