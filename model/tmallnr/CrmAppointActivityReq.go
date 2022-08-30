package tmallnr

// CrmAppointActivityReq 结构体
type CrmAppointActivityReq struct {
	// 微信名称
	WechatName string `json:"wechat_name,omitempty" xml:"wechat_name,omitempty"`
	// 业务code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 会员ID
	OpenId string `json:"open_id,omitempty" xml:"open_id,omitempty"`
	// 导购员ID
	GuiderId int64 `json:"guider_id,omitempty" xml:"guider_id,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 活动ID
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}
