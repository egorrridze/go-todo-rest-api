package todo

type User struct {
	Id       int    `json:"-"`
	//Name     string `json:"name"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	//Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`

}
