package domain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNotFound(t *testing.T) {
	// 1) вызов функции для теста
	user, err := UserDao.GetUser(0)

	// 2) валидация результата функции
	assert.Nil(t, user, "not expecting a user with id 0")
	assert.NotNilf(t, err, "expecting error when a user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, "not_found", err.Code)
}

func TestGetUserFound(t *testing.T) {
	// 1) вызов функции для теста
	user, err := UserDao.GetUser(123)

	// 2) валидация результата функции
	assert.NotNilf(t, user, "expecting a user with id 0")
	assert.Nil(t, err, "not expecting error when a user id is 0")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Lam", user.FirstName)
	assert.EqualValues(t, "Pham", user.LastName)
	assert.EqualValues(t, "test@gmail.com", user.Email)
}
