package auth

import (
	"Effective/pkg/app"
	"Effective/pkg/e"
	"Effective/pkg/util"
	"Effective/service/auth_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type login struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Login
// @Description Login
// @Produce  json
// @Param username query string false "username"
// @Success 200 {object} app.Response
// @Router /api/auth/login [post]
func Login(c *gin.Context)  {
	appG := app.Gin{c}
	valid := validation.Validation{}

	username := c.Query("username")
	password := c.Query("password")

	a := login{
		Username: username,
		Password: password,
	}
	ok, _ := valid.Valid(&a)
	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{
		Username: username,
		Password: password,
	}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(username, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"token" : token,
	})
}