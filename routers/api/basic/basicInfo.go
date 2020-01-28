package basic

import (
	"Effective/pkg/app"
	"Effective/pkg/e"
	"github.com/gin-gonic/gin"
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
}