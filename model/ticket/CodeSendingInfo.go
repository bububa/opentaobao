package ticket

// CodeSendingInfo 结构体
type CodeSendingInfo struct {
	// 发码方式。1、电子票自动发码 需设置电子凭证信息，2、电子票手工发码，3、实体票
	CodeMode int64 `json:"code_mode,omitempty" xml:"code_mode,omitempty"`
	// 电子凭证信息
	ElecInfo *ItemEleCertInfo `json:"elec_info,omitempty" xml:"elec_info,omitempty"`
	// 是否需要买家邮箱
	HasEmail bool `json:"has_email,omitempty" xml:"has_email,omitempty"`
}
