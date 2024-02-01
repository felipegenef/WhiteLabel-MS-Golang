package getOneUser

type UserDTO struct{
	Id string
	Name string
}


func NewUserDTO(id string, name string) UserDTO{
	return UserDTO{id,name}
}