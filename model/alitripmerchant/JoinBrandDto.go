package alitripmerchant

// JoinBrandDTO 
type JoinBrandDTO struct {
    // 品牌logo
    BrandLogo   string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
    // 品牌跳转url
    LogoRedirectUrl   string `json:"logo_redirect_url,omitempty" xml:"logo_redirect_url,omitempty"`
}
