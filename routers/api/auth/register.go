package auth

import (
	"Effective/models"
	"Effective/pkg/app"
	"Effective/pkg/e"
	"Effective/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

type register struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; Match(/^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,10}$/)"`
}

func Register(c *gin.Context)  {
	var appRes = app.Gin{c}
	id := uuid.NewV1().String()
	username := c.Query("username")
	password := c.Query("password") //由数字和字母组成，并且要同时含有数字和字母，且长度要在8-16位之间。

	valid := validation.Validation{}
	a := register{
		Username: username,
		Password: password,
	}
	ok, err := valid.Valid(&a)
	if err != nil {
		log.Printf("rules wrong: %s", err)
		return
	}

	if !ok {
		code := e.ERROR_REGISTER_FORMAT_FAIL
		appRes.Response(http.StatusBadRequest, code, nil)
		for _, err := range valid.Errors {
			log.Printf("Fail to valid username&password, err:%s, msg: %s", err.Key, err.Message)
		}
		return
	}

	exist, _ := models.IsUsernameExist(username)
	if exist {
		c.JSON(http.StatusOK, gin.H{
			"code" : e.ERROR_USER_EXIST,
			"msg" : e.GetMsg(e.ERROR_USER_EXIST),
		})

		return
	}

	data := make(map[string]interface{})
	data["username"] = username
	data["password"] = util.EncodeMD5(password)
	data["id"] = id
	if err := models.RegisterUsers(data); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code" : e.ERROR_USER_EXIST,
			"msg" : e.GetMsg(e.ERROR_USER_EXIST),
		})

		return
	}
	code := e.SUCCESS
	appRes.Response(http.StatusOK, code, nil)
}