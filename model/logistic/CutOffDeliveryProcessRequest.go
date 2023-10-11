package logistic

// CutOffDeliveryProcessRequest 结构体
type CutOffDeliveryProcessRequest struct {
	// 面单号
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 业务单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 配拦截发起时间
	InterceptTime string `json:"intercept_time,omitempty" xml:"intercept_time,omitempty"`
	// 配拦截发起原因
	InterceptReason string `json:"intercept_reason,omitempty" xml:"intercept_reason,omitempty"`
	// 快递公司编码
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
}
