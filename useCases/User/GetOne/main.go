package getOneUser

var repository = NewGetOneUserRepository()
var service = NewGetOneUserService(repository)
var Controller = NewGetOneUserController(service)