package getOneUser

type UserDTO struct{
	Id string `json:"id"`
	Name string `json:"name"`
}


func NewUserDTO(id string, name string) UserDTO{
	return UserDTO{id,name}
}