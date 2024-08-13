package app

import (
	"github.com/PTLam25/microserver-course-1/controllers"
)

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
