package car

import (
	"sync"
)

// TransferOrderDetailCallbackRsp 结构体
type TransferOrderDetailCallbackRsp struct {
	// 订单航班信息
	TransferFlightInfo *TransferFlightInfo `json:"transfer_flight_info,omitempty" xml:"transfer_flight_info,omitempty"`
	// 买家信息
	TransferBuyerInfo *TransferBuyerInfo `json:"transfer_buyer_info,omitempty" xml:"transfer_buyer_info,omitempty"`
	// 订单基础信息
	TransferOrderInfo *TransferOrderInfo `json:"transfer_order_info,omitempty" xml:"transfer_order_info,omitempty"`
	// 订单发票信息
	TransferInvoiceInfo *TransferInvoiceInfo `json:"transfer_invoice_info,omitempty" xml:"transfer_invoice_info,omitempty"`
	// 订单金额信息
	TransferOrderAmountInfo *TransferOrderAmountInfo `json:"transfer_order_amount_info,omitempty" xml:"transfer_order_amount_info,omitempty"`
	// 司机信息
	TransferDriveInfo *TransferDriveInfo `json:"transfer_drive_info,omitempty" xml:"transfer_drive_info,omitempty"`
	// 用车信息
	TransferUseCarInfo *TransferUseCarInfo `json:"transfer_use_car_info,omitempty" xml:"transfer_use_car_info,omitempty"`
}

var poolTransferOrderDetailCallbackRsp = sync.Pool{
	New: func() any {
		return new(TransferOrderDetailCallbackRsp)
	},
}

// GetTransferOrderDetailCallbackRsp() 从对象池中获取TransferOrderDetailCallbackRsp
func GetTransferOrderDetailCallbackRsp() *TransferOrderDetailCallbackRsp {
	return poolTransferOrderDetailCallbackRsp.Get().(*TransferOrderDetailCallbackRsp)
}

// ReleaseTransferOrderDetailCallbackRsp 释放TransferOrderDetailCallbackRsp
func ReleaseTransferOrderDetailCallbackRsp(v *TransferOrderDetailCallbackRsp) {
	v.TransferFlightInfo = nil
	v.TransferBuyerInfo = nil
	v.TransferOrderInfo = nil
	v.TransferInvoiceInfo = nil
	v.TransferOrderAmountInfo = nil
	v.TransferDriveInfo = nil
	v.TransferUseCarInfo = nil
	poolTransferOrderDetailCallbackRsp.Put(v)
}
