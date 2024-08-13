package services

// Сервисы содержат в себя бизнес логику для обработки данных от контроллеров. Сервисы могут общаться с другими сервисами (вызывать у себя в коде), а также к DAO для связи с БД, для получения данных.
import (
	"github.com/PTLam25/microserver-course-1/domain"
	"github.com/PTLam25/microserver-course-1/utils"
)

type userService struct { // создаем новый тип, чтобы через него вызывать методы в пакете
}

var (
	UserService userService // Декларация переменной, который будет использовать вне пакета
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	// тут может быть вызовы к другим сервисам
	return domain.UserDao.GetUser(userId)
}
