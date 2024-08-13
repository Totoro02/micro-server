package services

// Сервисы содержат в себя бизнес логику для обработки данных от контроллеров. Сервисы могут общаться с другими сервисами (вызывать у себя в коде), а также к DAO для связи с БД, для получения данных.
import (
	"github.com/PTLam25/microserver-course-1/domain"
	"github.com/PTLam25/microserver-course-1/utils"
	"net/http"
)

type itemService struct { // создаем новый тип, чтобы через него вызывать методы в пакете
}

var (
	ItemService itemService // Декларация переменной, который будет использовать вне пакета
)

func (is *itemService) GetItem(itemId int64) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement_me",
		StatusCode: http.StatusInternalServerError,
		Code:       "internal_error",
	}
}
