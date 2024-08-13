package controllers

// Контроллеры - это входная точка для любого запроса в нашем приложения.
// Они отвечают за валидации входящих данных и передача их сервисам на обработку. А потом возвращают ответ клиенту.
// В контроллерах НЕ должна содержаться бизнес логика. Бизнес логика содержатся в сервисах.
import (
	"fmt"
	"github.com/PTLam25/microserver-course-1/services"
	"github.com/PTLam25/microserver-course-1/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	// 1) валидируем данные
	userIdParam := c.Param("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}

		// добавить 400 статус в заголовок ответа и отправить текст ошибк
		utils.RespondError(c, apiErr)
		return
	}

	// 2) отправляем данные на обработку сервису
	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		// добавить 404 статус в заголовок ответа и отправить текст ошибки
		utils.RespondError(c, apiErr)
		return
	}

	// 4) вернуть ответ клиенту
	utils.Respond(c, http.StatusBadRequest, user)
}
