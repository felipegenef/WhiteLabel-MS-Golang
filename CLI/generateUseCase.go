package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What's the domain name? ")
	domainName, _ := reader.ReadString('\n')
	domainName = strings.TrimSpace(domainName)

	fmt.Print("What's the useCase name? ")
	useCaseName, _ := reader.ReadString('\n')
	useCaseName = strings.TrimSpace(useCaseName)

	useCasesPath := "./useCases"
	domainPath := filepath.Join(useCasesPath, domainName)
	useCasePath := filepath.Join(useCasesPath, domainName, useCaseName)

	// Criar a pasta ./useCases se não existir
	if _, err := os.Stat(useCasesPath); os.IsNotExist(err) {
		os.Mkdir(useCasesPath, 0755)
	}

	// Criar a pasta ./useCases/domainName se não existir
	if _, err := os.Stat(domainPath); os.IsNotExist(err) {
		os.Mkdir(domainPath, 0755)
	}

	os.Mkdir(useCasePath, 0755)

	// Criar os arquivos dentro da pasta do domainName
	controllerContent := fmt.Sprintf(`package %s

import (
	global "example.com/goLangMicroservice/Global/Interfaces"
	"github.com/gofiber/fiber/v2"
)

type %sController struct {
	service *%sService
}


type ErrorResponse struct {
	Message string ` + "`json:\"message\"`" + `
}

func (controller *%sController) Handle(ctx *fiber.Ctx) error {
	
	%s, err := controller.service.Execute(ctx.Params("id"))
	if err != nil {
		response := ErrorResponse{
			Message: "%s not found",
		}
		return ctx.Status(404).JSON(response)
	}



	return ctx.Status(200).JSON(%s)
}

func New%sController(service *%sService) global.Controller {
	return &%sController{
		service,
	}
}

`, capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, domainName, domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName))

	serviceContent := fmt.Sprintf(`package %s



	type %sService struct{
		%sRepository I%sRepository
	}
	
	
	func (c *%sService) Execute(id string) (DTO,error) {
		var %s,err= c.%sRepository.FindOne(id);
		if err !=nil{
			return %s,err
		}
		return %s,nil
	}
	
	func New%sService(%sRepository I%sRepository) *%sService{
		
		return &%sService{
			%sRepository: %sRepository,
		}
	}	
`, capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, domainName, domainName, domainName, domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName)

	dtosContent := fmt.Sprintf(`package %s

type DTO struct {
	Id string ` + "`json:\"id\"`" + `
	Name string ` + "`json:\"name\"`" + `
}
func NewDTO(id string, name string) DTO{
	return DTO{id,name}
}
`,capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName))

	repositoriesContent := fmt.Sprintf(`package %s

	import (
		"go.mongodb.org/mongo-driver/mongo"
	)
	
	
	type I%sRepository interface {
		FindOne(id string) (DTO, error)
	
	
	}
	type %sRepository struct{
		collection *mongo.Collection
	}
	
	func (repo *%sRepository) FindOne(id string) (DTO,error) {
		// var %s DTO
		var %s= NewDTO(id,"Jhon")
		// err:=db.EnsureConnected()
		// if err !=nil{
		// 	return %s, err
		// }
		// ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()
	
		// err=repo.collection.FindOne(ctx,bson.M{"id": id}).Decode(&%s)
		// if err !=nil{
		// 	return %s, err
		// }
	
		return %s,nil
	}
	
	func New%sRepository() I%sRepository{
		// connection,err:=db.GetConnectionCache()
		// if err !=nil{
		// 	panic("Could not get connection")
		// }
		// collection:=db.GetCollection(connection,"Users")
	
		return &%sRepository{
			// collection: collection,
		}
	}
`, capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, domainName, domainName, domainName, domainName, domainName, domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName))

	indexContent := fmt.Sprintf(`package %s

	var repository = New%sRepository()
	var service = New%sService(repository)
	var Controller = New%sController(service)
`, capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName))

	useCaseTestContent := fmt.Sprintf(`package %s

	import (
		"errors"
		"testing"
	)
	
	
	type Mock%sRepository struct{
	
	}
	
	func (m *Mock%sRepository) FindOne(id string) (DTO, error) {
		// Implementação fictícia para o mock, você pode ajustar conforme necessário para seus testes.
		if id == "1" {
			return DTO{Id: "1", Name: "John Doe"}, nil
		}
		return DTO{}, errors.New("%s não encontrado")
	}
	
	func Test%sService_Execute(t *testing.T) {
		// configure mock
		service := New%sService(&Mock%sRepository{})
	
		// Success useCase
		result, err := service.Execute("1")
		if err != nil {
			t.Errorf("Error on Test%sService %%v", err)
		}
	
		// Verify sucess useCase
		expectedResult := DTO{Id: "1", Name: "John Doe"}
		if result != expectedResult {
			t.Errorf("Expected %%v, got %%v", expectedResult, result)
		}
	
		// Error UseCase
		_, err = service.Execute("2")
		if err == nil {
			t.Error("Expected Error, got Success")
		}
	
	}
	
`, capitalizeFirstLetter(useCaseName)+capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), domainName, capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName), capitalizeFirstLetter(domainName))

	controllerFile, _ := os.Create(filepath.Join(useCasePath, "Controller.go"))
	defer controllerFile.Close()
	controllerFile.WriteString(controllerContent)

	serviceFile, _ := os.Create(filepath.Join(useCasePath, "Service.go"))
	defer serviceFile.Close()
	serviceFile.WriteString(serviceContent)

	dtosFile, _ := os.Create(filepath.Join(useCasePath, "DTOs.go"))
	defer dtosFile.Close()
	dtosFile.WriteString(dtosContent)

	repositoriesFile, _ := os.Create(filepath.Join(useCasePath, "Repositories.go"))
	defer repositoriesFile.Close()
	repositoriesFile.WriteString(repositoriesContent)

	indexFile, _ := os.Create(filepath.Join(useCasePath, "main.go"))
	defer indexFile.Close()
	indexFile.WriteString(indexContent)

	testFile, _ := os.Create(filepath.Join(useCasePath, "useCase_test.go"))
	defer testFile.Close()
	testFile.WriteString(useCaseTestContent)

	fmt.Println("Files created successfully.")
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
