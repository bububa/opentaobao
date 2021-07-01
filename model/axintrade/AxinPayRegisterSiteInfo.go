package axintrade

// AxinPayRegisterSiteInfo 结构体
type AxinPayRegisterSiteInfo struct {
	// 站点地址
	SiteUrl string `json:"site_url,omitempty" xml:"site_url,omitempty"`
	// 网站名称
	SiteName string `json:"site_name,omitempty" xml:"site_name,omitempty"`
	// 网站：01APP : 02服务窗:03公众号:04其他:05支付宝小程序:06
	SiteType string `json:"site_type,omitempty" xml:"site_type,omitempty"`
}
