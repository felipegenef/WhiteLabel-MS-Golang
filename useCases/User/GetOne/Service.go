package getOneUser



type GetOneUserService struct{}

func (c *GetOneUserService) Execute(id string) (UserDTO,error) {
	var user= NewUserDTO(id,"Jhon")
	return user,nil
}

func NewGetOneUserService() GetOneUserService{
	return GetOneUserService{}
}