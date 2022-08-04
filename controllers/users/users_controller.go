package users

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ralkmim/bookstore_users-api/domain/users"
	"github.com/ralkmim/bookstore_users-api/services"
	"github.com/ralkmim/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	/* fmt.Println(user)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: Handle error
		return
	}

	if err := json.Unmarshal(bytes, &user); err != nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	} */
	//THE PREVIOUS LINES OF CODE IS EQUIVALENT TO THE NEXT LINES OF CODE

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
 