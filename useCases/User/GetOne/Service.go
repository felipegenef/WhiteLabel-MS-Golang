package getOneUser



type GetOneUserService struct{
	UserRepository IGetOneUserRepository
}


func (c *GetOneUserService) Execute(id string) (UserDTO,error) {
	var user,err= c.UserRepository.FindOne(id);
	if err !=nil{
		return user,err
	}
	return user,nil
}

func NewGetOneUserService(userRepository IGetOneUserRepository) *GetOneUserService{
	
	return &GetOneUserService{
		UserRepository: userRepository,
	}
}