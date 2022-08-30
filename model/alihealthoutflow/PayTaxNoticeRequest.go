package alihealthoutflow

// PayTaxNoticeRequest 结构体
type PayTaxNoticeRequest struct {
	// 审核操作类型
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 医生id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
