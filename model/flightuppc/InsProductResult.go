package flightuppc

// InsProductResult 结构体
type InsProductResult struct {
	// 保险产品编码
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 保险产品名称
	ProdName string `json:"prod_name,omitempty" xml:"prod_name,omitempty"`
	// 销售单元
	CsuNo string `json:"csu_no,omitempty" xml:"csu_no,omitempty"`
	// 保险价格
	Premium int64 `json:"premium,omitempty" xml:"premium,omitempty"`
	// 生成的保险订单号
	TcOrderId int64 `json:"tc_order_id,omitempty" xml:"tc_order_id,omitempty"`
	// 保险订单关联的外部订单号
	OutOrderId int64 `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// 保险订单号
	InsOrderId int64 `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	// 保险产品id
	PremiumId int64 `json:"premium_id,omitempty" xml:"premium_id,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
