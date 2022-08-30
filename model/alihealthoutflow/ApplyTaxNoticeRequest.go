package alihealthoutflow

// ApplyTaxNoticeRequest 结构体
type ApplyTaxNoticeRequest struct {
	// 审核事件类型
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// appKey
	AppKey string `json:"app_key,omitempty" xml:"app_key,omitempty"`
	// 来源
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// 申请单编号
	ApplyBatchNum string `json:"apply_batch_num,omitempty" xml:"apply_batch_num,omitempty"`
	// 扩展信息(JSON格式)
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}
