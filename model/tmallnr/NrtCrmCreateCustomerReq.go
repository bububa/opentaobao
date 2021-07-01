package tmallnr

// NrtCrmCreateCustomerReq 结构体
type NrtCrmCreateCustomerReq struct {
	// 活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 业务身份：居然/红星
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 会员openId
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 姓名
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 摊位id
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 导购员id
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
}
