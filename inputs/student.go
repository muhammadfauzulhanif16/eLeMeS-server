package inputs

type AddStudent struct {
	FullName string `json:"full_name" binding:"required"`
	SIN      int    `json:"sin" binding:"required"`
	Password string `json:"password" binding:"required"`
}
