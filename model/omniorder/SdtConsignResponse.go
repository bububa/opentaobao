package omniorder

// SdtConsignResponse 结构体
type SdtConsignResponse struct {
	// 接单公司
	CompanyName string `json:"company_name,omitempty" xml:"company_name,omitempty"`
	// 接单网点
	DotName string `json:"dot_name,omitempty" xml:"dot_name,omitempty"`
	// 接单小件员电话
	DeliveryUserPhone string `json:"delivery_user_phone,omitempty" xml:"delivery_user_phone,omitempty"`
	// 接单小件员姓名
	DeliveryUserName string `json:"delivery_user_name,omitempty" xml:"delivery_user_name,omitempty"`
	// 商家id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 包裹id
	PackageId int64 `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 物流订单号list
	LpCodeList string `json:"lp_code_list,omitempty" xml:"lp_code_list,omitempty"`
}
