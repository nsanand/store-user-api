package users

import "C"
import (
	"github.com/gin-gonic/gin"
	"github.com/nsanand/store-user-api/domain/users"
	"github.com/nsanand/store-user-api/services"
	"github.com/nsanand/store-user-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context)  {
	var user users.User
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// TODO handle error
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	// TODO handle error
	//	return
	//}
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status,  saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(http.StatusOK, user)
	}
}

func SearchUser(c *gin.Context)  {

}
