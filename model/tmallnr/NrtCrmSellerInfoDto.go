package tmallnr

// NrtCrmSellerInfoDto 结构体
type NrtCrmSellerInfoDto struct {
	// 卖家logo
	NativeBrandLogo string `json:"native_brand_logo,omitempty" xml:"native_brand_logo,omitempty"`
	// nativeOpenMember
	NativeOpenMember string `json:"native_open_member,omitempty" xml:"native_open_member,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 商家名称
	SellerName string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
}
