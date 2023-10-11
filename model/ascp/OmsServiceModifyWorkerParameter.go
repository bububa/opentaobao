package ascp

// OmsServiceModifyWorkerParameter 结构体
type OmsServiceModifyWorkerParameter struct {
	// 师傅id
	WorkerId string `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
	// 师傅姓名
	WorkerName string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
	// 师傅电话
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// BFC单号
	WdsCoordinationOrderId string `json:"wds_coordination_order_id,omitempty" xml:"wds_coordination_order_id,omitempty"`
	// 服务工单di
	WorkcardId string `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
	// 服务商id
	TpId string `json:"tp_id,omitempty" xml:"tp_id,omitempty"`
	// 商家id
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 评分
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 完工数量
	CompletionQuantity int64 `json:"completion_quantity,omitempty" xml:"completion_quantity,omitempty"`
}
