package content

import (
	"Effective/models"
	"Effective/pkg/app"
	"Effective/pkg/e"
	"Effective/service/content_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Column struct {
	models.Model
	Name	string `json:"name"`
	ParentID	int `json:"parent_id"`
	CoverImageUrl string `json:"cover_image_url"`
}

func GetColumn(c *gin.Context)  {
	appG := app.Gin{c}

	columnService := content_service.Column{}

	column, err := columnService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_BASIC_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, column)
}