package ascpchannel

import (
	"sync"
)

// WaybillGenRequest 结构体
type WaybillGenRequest struct {
	// 货品信息
	WaybillGenItemList []WaybillGenItemList `json:"waybill_gen_item_list,omitempty" xml:"waybill_gen_item_list>waybill_gen_item_list,omitempty"`
	// 服务类型
	WaybillGenServ *WaybillGenServ `json:"waybill_gen_serv,omitempty" xml:"waybill_gen_serv,omitempty"`
	// 发货人信息
	WaybillGenSender *WaybillGenSender `json:"waybill_gen_sender,omitempty" xml:"waybill_gen_sender,omitempty"`
	// 收货人信息
	WaybillGenReceiver *WaybillGenReceiver `json:"waybill_gen_receiver,omitempty" xml:"waybill_gen_receiver,omitempty"`
	// 开单信息
	WaybillGenInfo *WaybillGenInfo `json:"waybill_gen_info,omitempty" xml:"waybill_gen_info,omitempty"`
}

var poolWaybillGenRequest = sync.Pool{
	New: func() any {
		return new(WaybillGenRequest)
	},
}

// GetWaybillGenRequest() 从对象池中获取WaybillGenRequest
func GetWaybillGenRequest() *WaybillGenRequest {
	return poolWaybillGenRequest.Get().(*WaybillGenRequest)
}

// ReleaseWaybillGenRequest 释放WaybillGenRequest
func ReleaseWaybillGenRequest(v *WaybillGenRequest) {
	v.WaybillGenItemList = v.WaybillGenItemList[:0]
	v.WaybillGenServ = nil
	v.WaybillGenSender = nil
	v.WaybillGenReceiver = nil
	v.WaybillGenInfo = nil
	poolWaybillGenRequest.Put(v)
}
