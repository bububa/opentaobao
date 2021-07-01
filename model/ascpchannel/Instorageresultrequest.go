package ascpchannel

// Instorageresultrequest 结构体
type Instorageresultrequest struct {
	// 逆向物流单号
	LgOrderCode string `json:"lg_order_code,omitempty" xml:"lg_order_code,omitempty"`
	// 操作信息
	OperateInfo *Operateinfo `json:"operate_info,omitempty" xml:"operate_info,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 入库时间
	OrderConfirmTime string `json:"order_confirm_time,omitempty" xml:"order_confirm_time,omitempty"`
}
