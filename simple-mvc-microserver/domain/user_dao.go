package domain

// DOA отвечает за связь модели С БД для CRUD.
import (
	"fmt"
	"github.com/PTLam25/microserver-course-1/utils"
	"log"
	"net/http"
)

func init() {
	// init() вызывается только при 1-й инициализация пакета (1-ого вызова импорта в приложения), потом он хранится в кеше, поэтому при следующем обращения init() не вызовется
	UserDao = &userDao{}
}

type userServiceInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}
type userDao struct{}

var (
	// типо БД, в котором есть пользователи
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Lam", LastName: "Pham", Email: "test@gmail.com"},
	}

	UserDao userServiceInterface
)

func (ud *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We are accessing DB")
	//  получить данные с БД
	if user := users[userId]; user != nil {
		return user, nil

	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
