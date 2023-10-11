package ascp

// WorkerPriceInfo 结构体
type WorkerPriceInfo struct {
	// 师傅id
	WorkerId string `json:"worker_id,omitempty" xml:"worker_id,omitempty"`
	// 师傅姓名
	WorkerName string `json:"worker_name,omitempty" xml:"worker_name,omitempty"`
	// 师傅电话
	WorkerMobile string `json:"worker_mobile,omitempty" xml:"worker_mobile,omitempty"`
	// 评分
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 保价金额：单位：分
	PriceAmount string `json:"price_amount,omitempty" xml:"price_amount,omitempty"`
	// 平台完工量
	CompletionQuantity int64 `json:"completion_quantity,omitempty" xml:"completion_quantity,omitempty"`
}
