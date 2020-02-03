package basic

import (
	"Effective/models"
	"Effective/pkg/app"
	"Effective/pkg/e"
	"Effective/service/info_service"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

type AddBasicInfo struct {
	WebsiteName string `form:"website_name" valid:"MaxSize(255)"`
	WebsiteUrl string `form:"website_url" valid:"MaxSize(255)"`
	WebsiteLogoUrl string `form:"website_logo_url" valid:"MaxSize(255)"`
	WebsiteFaviconUrl string `form:"website_favicon_url" valid:"MaxSize(255)"`
	WebsiteCompanyName string `form:"website_company_name" valid:"MaxSize(255)"`
	WebsiteContact string `form:"website_contact" valid:"MaxSize(255)"`
	WebsiteMobile string `form:"website_mobile" valid:"MaxSize(255)"`
	WebsiteHotLine string `form:"website_hot_line" valid:"MaxSize(255)"`
	WebsiteTel string `form:"website_tel" valid:"MaxSize(255)"`
	WebsiteFax string `form:"website_fax" valid:"MaxSize(255)"`
	WebsiteQQ string `form:"website_qq" valid:"MaxSize(255)"`
	WebsiteEmail string `form:"website_email" valid:"MaxSize(255)"`
	WebsiteAddress string `form:"website_address" valid:"MaxSize(255)"`
	WebsiteQRcode string `form:"website_qrcode_url" valid:"MaxSize(255)"`
	WebsiteHeaderMsg string `form:"website_header_msg" valid:"MaxSize(255)"`
	WebsiteFooterMsg string `form:"website_footer_msg" valid:"MaxSize(255)"`
	WebsiteCopyright string `form:"website_copyright" valid:"MaxSize(255)"`
	WebsiteIcp string `form:"website_icp" valid:"MaxSize(255)"`
	WebsiteSupport string `form:"website_support" valid:"MaxSize(255)"`
	WebsiteDisclaim string `form:"website_disclaim" valid:"MaxSize(255)"`
}

func EditBasicInfo(c *gin.Context)  {
	var (
		appG = app.Gin{c}
		form AddBasicInfo
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	infoService := info_service.BasicInfo{
		WebsiteName: form.WebsiteName,
		WebsiteUrl : form.WebsiteUrl,
		WebsiteLogoUrl : form.WebsiteLogoUrl,
		WebsiteFaviconUrl : form.WebsiteFaviconUrl,
		WebsiteCompanyName : form.WebsiteCompanyName,
		WebsiteContact : form.WebsiteContact,
		WebsiteMobile : form.WebsiteMobile,
		WebsiteHotLine : form.WebsiteHotLine,
		WebsiteTel : form.WebsiteTel,
		WebsiteFax : form.WebsiteFax,
		WebsiteQQ : form.WebsiteQQ,
		WebsiteEmail : form.WebsiteEmail,
		WebsiteAddress : form.WebsiteAddress,
		WebsiteQRcode : form.WebsiteQRcode,
		WebsiteHeaderMsg : form.WebsiteHeaderMsg,
		WebsiteFooterMsg : form.WebsiteFooterMsg,
		WebsiteCopyright : form.WebsiteCopyright,
		WebsiteIcp : form.WebsiteIcp,
		WebsiteSupport : form.WebsiteSupport,
		WebsiteDisclaim : form.WebsiteDisclaim,
	}

	if err := infoService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_BASIC_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func GetBasicInfo(c *gin.Context)  {
	appG := app.Gin{c}
	id := com.StrTo(c.Query("id")).MustInt()
	fmt.Println(id)
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	infoService := info_service.BasicInfo{
		Model: models.Model{ID: id},
	}

	info, err := infoService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_BASIC_INFO_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, info)
}