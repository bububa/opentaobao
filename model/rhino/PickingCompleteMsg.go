package rhino

import (
	"sync"
)

// PickingCompleteMsg 结构体
type PickingCompleteMsg struct {
	// 运货单信息package_infos
	Waybills []ClothingOutboundWaybill `json:"waybills,omitempty" xml:"waybills>clothing_outbound_waybill,omitempty"`
	// SAP交货单号torder_code
	SapInvoiceId string `json:"sap_invoice_id,omitempty" xml:"sap_invoice_id,omitempty"`
	// 仓库出库单完成时间shipment_confirm_time
	OutboundDate string `json:"outbound_date,omitempty" xml:"outbound_date,omitempty"`
	// 出库单编码shipment_id
	WmsInvoiceId string `json:"wms_invoice_id,omitempty" xml:"wms_invoice_id,omitempty"`
	// 货主编号company
	FactoryId int64 `json:"factory_id,omitempty" xml:"factory_id,omitempty"`
	// 仓库编码
	WareHouseId int64 `json:"ware_house_id,omitempty" xml:"ware_house_id,omitempty"`
	// 出库单单据类型，1-正常出库2-手工出库
	OutboundType int64 `json:"outbound_type,omitempty" xml:"outbound_type,omitempty"`
}

var poolPickingCompleteMsg = sync.Pool{
	New: func() any {
		return new(PickingCompleteMsg)
	},
}

// GetPickingCompleteMsg() 从对象池中获取PickingCompleteMsg
func GetPickingCompleteMsg() *PickingCompleteMsg {
	return poolPickingCompleteMsg.Get().(*PickingCompleteMsg)
}

// ReleasePickingCompleteMsg 释放PickingCompleteMsg
func ReleasePickingCompleteMsg(v *PickingCompleteMsg) {
	v.Waybills = v.Waybills[:0]
	v.SapInvoiceId = ""
	v.OutboundDate = ""
	v.WmsInvoiceId = ""
	v.FactoryId = 0
	v.WareHouseId = 0
	v.OutboundType = 0
	poolPickingCompleteMsg.Put(v)
}
