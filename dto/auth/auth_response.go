package authdto

type LoginResponse struct {
	Token string `gorm:"type: varchar(255)" json:"token"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	FullName string `gorm:"type: varchar(255)" json:"full_name"`
}

type CheckAuthResponse struct {
	ID    int    `gorm:"type: int" json:"id"`
	Name  string `gorm:"type: varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
}

type AuthResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}
