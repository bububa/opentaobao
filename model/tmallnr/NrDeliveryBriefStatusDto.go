package tmallnr

// NrDeliveryBriefStatusDto 结构体
type NrDeliveryBriefStatusDto struct {
	// 状态产生时间
	LogisticsTime string `json:"logistics_time,omitempty" xml:"logistics_time,omitempty"`
	// 取件失败的code
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// 取件失败的原因
	FailReason string `json:"fail_reason,omitempty" xml:"fail_reason,omitempty"`
	// 快递员电话
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 快递员姓名
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 服务商的cp
	SpName string `json:"sp_name,omitempty" xml:"sp_name,omitempty"`
	// logisticsStatusName状态的说明包含[下单,接单,取件,妥投,拒收，取消]
	LogisticsStatusName string `json:"logistics_status_name,omitempty" xml:"logistics_status_name,omitempty"`
	// 包含[CREATE,GRASP,GOT,DELIVERYED]
	LogisticsStatus string `json:"logistics_status,omitempty" xml:"logistics_status,omitempty"`
}
