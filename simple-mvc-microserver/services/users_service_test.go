package services

import (
	"github.com/PTLam25/microserver-course-1/domain"
	"github.com/PTLam25/microserver-course-1/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func init() {
	// init() вызывается только при 1-й инициализация пакета (1-ого вызова импорта в приложения), потом он хранится в кеше, поэтому при следующем обращения init() не вызовется

	// перезначем на userDaoMock для тестирование
	domain.UserDao = &userDaoMock{}
}

var (
	UserDaoMock userDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

type userDaoMock struct {
}

func (udm *userDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	log.Println("We are accessing MOCK DB")
	return getUserFunction(userId)
}

func TestGetUserNotFound(t *testing.T) {
	// 1) инициалзиация данных
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
			Code:       "not_found",
		}
	}

	// 2) вызов функции
	user, err := UserService.GetUser(0)

	// 3) валидация
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Lam",
			LastName:  "Pham",
			Email:     "test@gmail.com",
		}, nil
	}

	// 1) вызов функции для теста
	user, err := UserService.GetUser(123)

	// 2) валидация результата функции
	assert.NotNilf(t, user, "expecting a user with id 0")
	assert.Nil(t, err, "not expecting error when a user id is 0")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Lam", user.FirstName)
	assert.EqualValues(t, "Pham", user.LastName)
	assert.EqualValues(t, "test@gmail.com", user.Email)
}
