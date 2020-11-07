package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hikaru-sh/bookstore_users_api/domain/users"
	"github.com/hikaru-sh/bookstore_users_api/services"
	"github.com/hikaru-sh/bookstore_users_api/utils/errors"
)

var(
	counter int
)
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
    restErr := errors.NewBadRequestError("invalid json body.")
		c.JSON(restErr.Status, restErr)
		return
	}
	fmt.Println(user)

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
  c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number.")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
  c.JSON(http.StatusCreated, user)
}

