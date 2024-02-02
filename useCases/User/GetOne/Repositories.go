package getOneUser

import (
	"go.mongodb.org/mongo-driver/mongo"
)


type IGetOneUserRepository interface {
	FindOne(id string) (UserDTO, error)


}
type GetOneUserRepository struct{
	collection *mongo.Collection
}

func (repo *GetOneUserRepository) FindOne(id string) (UserDTO,error) {
	// var user UserDTO
	var user= NewUserDTO(id,"Jhon")
	// err:=db.EnsureConnected()
	// if err !=nil{
	// 	return user, err
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	// defer cancel()

	// err=repo.collection.FindOne(ctx,bson.M{"id": id}).Decode(&user)
	// if err !=nil{
	// 	return user, err
	// }

	return user,nil
}

func NewGetOneUserRepository() IGetOneUserRepository{
	// connection,err:=db.GetConnectionCache()
	// if err !=nil{
	// 	panic("Could not get connection")
	// }
	// collection:=db.GetCollection(connection,"Users")

	return &GetOneUserRepository{
		// collection: collection,
	}
}