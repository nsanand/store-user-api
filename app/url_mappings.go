package app

import (
	"github.com/nsanand/store-user-api/controllers/ping"
	"github.com/nsanand/store-user-api/controllers/users"
)

func mapUrls()  {
	router.GET("/ping", ping.Ping)
	router.GET("/users/search", users.SearchUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
