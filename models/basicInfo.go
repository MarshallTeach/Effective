package models

type BasicInfo struct {
	Model
	WebsiteName string `json:"website_name"`
	WebsiteUrl string `json:"website_url"`
	WebsiteLogoUrl string `json:"website_logo_url"`
	WebsiteFaviconUrl string `json:"website_favicon_url"`
	WebsiteCompanyName string `json:"website_company_name"`
	WebsiteContact string `json:"website_contact"`
	WebsiteMobile string `json:"website_mobile"`
	WebsiteHotLine string `json:"website_hot_line"`
	WebsiteTel string `json:"website_tel"`
	WebsiteFax string `json:"website_fax"`
	WebsiteQQ string `json:"website_qq"`
	WebsiteEmail string `json:"website_email"`
	WebsiteAddress string `json:"website_address"`
	WebsiteQRcode string `json:"website_qrcode_url"`
	WebsiteHeaderMsg string `json:"website_header_msg"`
	WebsiteFooterMsg string `json:"website_footer_msg"`
	WebsiteCopyright string `json:"website_copyright"`
	WebsiteIcp string `json:"website_icp"`
	WebsiteSupport string `json:"website_support"`
	WebsiteDisclaim string `json:"website_disclaim"`
}

func EditBasicInfo(interface{}) (bool, error) {

}