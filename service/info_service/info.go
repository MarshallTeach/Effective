package info_service

import (
	"Effective/models"
	"fmt"
)

type BasicInfo struct {
	models.Model
	WebsiteName string
	WebsiteUrl string
	WebsiteLogoUrl string
	WebsiteFaviconUrl string
	WebsiteCompanyName string
	WebsiteContact string
	WebsiteMobile string
	WebsiteHotLine string
	WebsiteTel string
	WebsiteFax string
	WebsiteQQ string
	WebsiteEmail string
	WebsiteAddress string
	WebsiteQRcode string
	WebsiteHeaderMsg string
	WebsiteFooterMsg string
	WebsiteCopyright string
	WebsiteIcp string
	WebsiteSupport string
	WebsiteDisclaim string
}

func (b *BasicInfo) Add() error {
	basicInfo := map[string]interface{}{
		"website_name" : b.WebsiteName,
		"website_url" : b.WebsiteUrl,
		"website_logo_url" : b.WebsiteLogoUrl,
		"website_favicon_url" : b.WebsiteFaviconUrl,
		"website_company_name" : b.WebsiteCompanyName,
		"website_contact" : b.WebsiteContact,
		"website_mobile" : b.WebsiteMobile,
		"website_hot_line" : b.WebsiteHotLine,
		"website_tel" : b.WebsiteTel,
		"website_fax" : b.WebsiteFax,
		"website_qq" : b.WebsiteQQ,
		"website_email" : b.WebsiteEmail,
		"website_address" : b.WebsiteAddress,
		"website_qrcode_url" : b.WebsiteQRcode,
		"website_header_msg" : b.WebsiteHeaderMsg,
		"website_footer_msg" : b.WebsiteFooterMsg,
		"website_copyright" : b.WebsiteCopyright,
		"website_icp" : b.WebsiteIcp,
		"website_support" : b.WebsiteSupport,
		"website_disclaim" : b.WebsiteDisclaim,
	}

	if err := models.AddBasicInfo(basicInfo); err != nil {
		return err
	}

	return nil
}

func (b *BasicInfo) Get() (*models.BasicInfo, error) {
/*	var cacheInfo *models.BasicInfo
	fmt.Println("b.ID", b.ID)
	cache := cache_service.Info{ID: b.ID}
	key := cache.GetInfoKey()
	fmt.Println("KEY", key)
	fmt.Println("CACHE", &cacheInfo)*/
	/*if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheInfo)
			return cacheInfo, nil
		}
	}*/

	info, err := models.GetInfo(b.ID)
	fmt.Println("infoservice info", info)
	if err != nil {
		return nil, err
	}

	/*gredis.Set(key, *info, 3600)*/
	return info, nil
}
