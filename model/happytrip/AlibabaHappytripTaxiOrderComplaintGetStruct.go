package happytrip

// AlibabahappytriptaxiordercomplaintgetStruct 结构体
type AlibabahappytriptaxiordercomplaintgetStruct struct {
	// 供应商订单号
	OrderId string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// 供应商工单号
	CaseId string `json:"case_id,omitempty" xml:"case_id,omitempty"`
	// 处理结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
	// 处理状态，0：处理中，1：处理完成
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
}
