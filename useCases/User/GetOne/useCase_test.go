package getOneUser

import (
	"errors"
	"testing"
)


type MockGetOneUserRepository struct{

}

func (m *MockGetOneUserRepository) FindOne(id string) (UserDTO, error) {
	// Implementação fictícia para o mock, você pode ajustar conforme necessário para seus testes.
	if id == "1" {
		return UserDTO{Id: "1", Name: "John Doe"}, nil
	}
	return UserDTO{}, errors.New("Usuário não encontrado")
}

func TestGetOneUserService_Execute(t *testing.T) {
	// configure mock
	service := NewGetOneUserService(&MockGetOneUserRepository{})

	// Success useCase
	result, err := service.Execute("1")
	if err != nil {
		t.Errorf("Error on TestGetOneUserService %v", err)
	}

	// Verify sucess useCase
	expectedResult := UserDTO{Id: "1", Name: "John Doe"}
	if result != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, result)
	}

	// Error UseCase
	_, err = service.Execute("2")
	if err == nil {
		t.Error("Expected Error, got Success")
	}

}
