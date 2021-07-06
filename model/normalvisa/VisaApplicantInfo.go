package normalvisa

// VisaApplicantInfo 结构体
type VisaApplicantInfo struct {
	// 申请人id
	ApplyId string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 子订单号
	SubTcOrderId int64 `json:"sub_tc_order_id,omitempty" xml:"sub_tc_order_id,omitempty"`
}
