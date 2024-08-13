package app

// отвечает за конфигурацию и запуск севера

import (
	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	// init() вызывается только при 1-й инициализация пакета (1-ого вызова импорта в приложения),
	//потом он хранится в кеше, поэтому при следующем обращения init() не вызовется

	router = gin.Default()
}

func StartApp() {
	// config urls
	mapUrls()

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
