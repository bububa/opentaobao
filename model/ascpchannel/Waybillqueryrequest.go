package ascpchannel

import (
	"sync"
)

// Waybillqueryrequest 结构体
type Waybillqueryrequest struct {
	// 包裹列表
	Packages []WaybillQueryRequestData `json:"packages,omitempty" xml:"packages>waybill_query_request_data,omitempty"`
	// 操作人id
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 操作人名称
	OperatorName string `json:"operator_name,omitempty" xml:"operator_name,omitempty"`
	// 供应商id
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 物流服务编码
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 发货LP单号
	ConsignLpOrderCode string `json:"consign_lp_order_code,omitempty" xml:"consign_lp_order_code,omitempty"`
	// 自营接口配业务模式，默认为1代表商家仓自营配 (为1时会强制校验发货的配CP和单号必须与取号时一致，且多包裹必须一次性发货)
	BusinessModel string `json:"business_model,omitempty" xml:"business_model,omitempty"`
}

var poolWaybillqueryrequest = sync.Pool{
	New: func() any {
		return new(Waybillqueryrequest)
	},
}

// GetWaybillqueryrequest() 从对象池中获取Waybillqueryrequest
func GetWaybillqueryrequest() *Waybillqueryrequest {
	return poolWaybillqueryrequest.Get().(*Waybillqueryrequest)
}

// ReleaseWaybillqueryrequest 释放Waybillqueryrequest
func ReleaseWaybillqueryrequest(v *Waybillqueryrequest) {
	v.Packages = v.Packages[:0]
	v.Operator = ""
	v.OperatorName = ""
	v.SupplierId = ""
	v.ServiceCode = ""
	v.ConsignLpOrderCode = ""
	v.BusinessModel = ""
	poolWaybillqueryrequest.Put(v)
}
