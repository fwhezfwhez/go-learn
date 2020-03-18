package userRouter

import (
	"github.com/gin-gonic/gin"
	"go-learn/projects/online-edu/user/userService"
)

func HTTPUserRouter(r gin.IRoutes) {
	r.POST("/user-info/", userService.HTTPAddUserInfo)
	r.PATCH("/user-info/:id/", userService.HTTPUpdateUserInfo)
	r.GET("/user-info/", userService.HTTPListUserInfo)
	r.GET("/user-info/:id/", userService.HTTPGetUserInfo)
	r.DELETE("/user-info/:id", userService.HTTPDeleteUserInfo)
}
