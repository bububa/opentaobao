package waybill

import (
	"sync"
)

// WaybillApplyCancelRequest 结构体
type WaybillApplyCancelRequest struct {
	// 交易订单列表
	TradeOrderList []string `json:"trade_order_list,omitempty" xml:"trade_order_list>string,omitempty"`
	// CP快递公司编码
	CpCode string `json:"cp_code,omitempty" xml:"cp_code,omitempty"`
	// 电子面单号码
	WaybillCode string `json:"waybill_code,omitempty" xml:"waybill_code,omitempty"`
	// ERP订单号或包裹号
	PackageId string `json:"package_id,omitempty" xml:"package_id,omitempty"`
	// 面单使用者编号
	RealUserId int64 `json:"real_user_id,omitempty" xml:"real_user_id,omitempty"`
}

var poolWaybillApplyCancelRequest = sync.Pool{
	New: func() any {
		return new(WaybillApplyCancelRequest)
	},
}

// GetWaybillApplyCancelRequest() 从对象池中获取WaybillApplyCancelRequest
func GetWaybillApplyCancelRequest() *WaybillApplyCancelRequest {
	return poolWaybillApplyCancelRequest.Get().(*WaybillApplyCancelRequest)
}

// ReleaseWaybillApplyCancelRequest 释放WaybillApplyCancelRequest
func ReleaseWaybillApplyCancelRequest(v *WaybillApplyCancelRequest) {
	v.TradeOrderList = v.TradeOrderList[:0]
	v.CpCode = ""
	v.WaybillCode = ""
	v.PackageId = ""
	v.RealUserId = 0
	poolWaybillApplyCancelRequest.Put(v)
}
