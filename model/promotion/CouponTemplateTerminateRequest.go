package promotion

// CouponTemplateTerminateRequest 结构体
type CouponTemplateTerminateRequest struct {
	// ump模板ID
	SourceId int64 `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// 用户信息
	UserInfo *UserInfo `json:"user_info,omitempty" xml:"user_info,omitempty"`
}
