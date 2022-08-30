package alihealthoutflow

// PayTaxValidationRequest 结构体
type PayTaxValidationRequest struct {
	// 信息获取凭证
	Ticket string `json:"ticket,omitempty" xml:"ticket,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 医生id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
}
