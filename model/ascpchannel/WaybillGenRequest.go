package ascpchannel

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
