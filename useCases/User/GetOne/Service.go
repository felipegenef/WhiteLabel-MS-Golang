package getOneUser



type GetOneUserService struct{
	UserRepository GetOneUserRepository
}

func (c *GetOneUserService) Execute(id string) (UserDTO,error) {
	var user,err= c.UserRepository.FindOne(id);
	if err !=nil{
		return user,err
	}
	return user,nil
}

func NewGetOneUserService(userRepository GetOneUserRepository) GetOneUserService{
	return GetOneUserService{
		UserRepository: userRepository,
	}
}