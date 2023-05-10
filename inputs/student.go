package inputs

type AddStudent struct {
	FullName string `json:"full_name" binding:"required"`
	SIDN     int    `json:"sidn" binding:"required"`
	Password string `json:"password" binding:"required"`
}
