package wdk

// VerificateCallbackDto 结构体
type VerificateCallbackDto struct {
	// 核销说明, 核销失败则填写核销失败原因
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 核销时间 YYYY-MM-DD HH:MI:SS
	VerificateTime string `json:"verificate_time,omitempty" xml:"verificate_time,omitempty"`
	// 核销账单ID
	BillOrderId string `json:"bill_order_id,omitempty" xml:"bill_order_id,omitempty"`
	// 业务经营店ID
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 商家编码
	MerchantCode string `json:"merchant_code,omitempty" xml:"merchant_code,omitempty"`
	// 核销状态 1=核销完成 2=核销失败
	VerificateStatus int64 `json:"verificate_status,omitempty" xml:"verificate_status,omitempty"`
	// 核销账单类型  1=正向 / 2=逆向
	BillType int64 `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
}
