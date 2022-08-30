package alsc

// RecordOpenSyncRequest 结构体
type RecordOpenSyncRequest struct {
	// 业务方
	BizSource string `json:"biz_source,omitempty" xml:"biz_source,omitempty"`
	// 同步请求扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 加密openId的appKey
	OpenAppKey string `json:"open_app_key,omitempty" xml:"open_app_key,omitempty"`
	// 开放平台用户id
	OpenUserId string `json:"open_user_id,omitempty" xml:"open_user_id,omitempty"`
	// 用户账号类型
	OwnerType string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	// 同步消费记录模型
	RecordOpenInfo *AlscConsumeRecordOpenInfo `json:"record_open_info,omitempty" xml:"record_open_info,omitempty"`
}
