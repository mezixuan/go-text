package models

type UserBasic struct {
	Id       int
	Name     string
	Identity string
	Password string
	Email    string
}

func (table UserBasic) Tablename() string {
	return "user_basic"
}
