package alitripmerchant

// JoinBrandDto 结构体
type JoinBrandDto struct {
	// 品牌logo
	BrandLogo string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
	// 品牌跳转url
	LogoRedirectUrl string `json:"logo_redirect_url,omitempty" xml:"logo_redirect_url,omitempty"`
}
