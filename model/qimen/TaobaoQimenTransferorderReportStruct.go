package qimen

import (
	"sync"
)

// TaobaoQimenTransferorderReportStruct 结构体
type TaobaoQimenTransferorderReportStruct struct {
	// 项目集
	Items []Items `json:"items,omitempty" xml:"items>items,omitempty"`
	// 调拨单号,0,string(50),必填,
	TransferOrderCode string `json:"transferOrderCode,omitempty" xml:"transferOrderCode,omitempty"`
	// 调拨出库单号
	TransferOutOrderCode string `json:"transferOutOrderCode,omitempty" xml:"transferOutOrderCode,omitempty"`
	// 调拨入库单号
	TransferInOrderCode string `json:"transferInOrderCode,omitempty" xml:"transferInOrderCode,omitempty"`
	// 确认出库时间
	ConfirmOutTime string `json:"confirmOutTime,omitempty" xml:"confirmOutTime,omitempty"`
	// 确认入库时间
	ConfirmInTime string `json:"confirmInTime,omitempty" xml:"confirmInTime,omitempty"`
	// 调拨单创建时间
	CreateTime string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 调拨出库仓编码
	FromWarehouseCode string `json:"fromWarehouseCode,omitempty" xml:"fromWarehouseCode,omitempty"`
	// 调拨入库仓编码
	ToWarehouseCode string `json:"toWarehouseCode,omitempty" xml:"toWarehouseCode,omitempty"`
	// 111
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// erpOrderCode
	ErpOrderCode string `json:"erpOrderCode,omitempty" xml:"erpOrderCode,omitempty"`
	// orderStatus
	OrderStatus string `json:"orderStatus,omitempty" xml:"orderStatus,omitempty"`
	// 响应结果:success|failure,success,string(10),必填,
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码,0,string(50),,
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息,invalid appkey,string(100),,
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenTransferorderReportStruct = sync.Pool{
	New: func() any {
		return new(TaobaoQimenTransferorderReportStruct)
	},
}

// GetTaobaoQimenTransferorderReportStruct() 从对象池中获取TaobaoQimenTransferorderReportStruct
func GetTaobaoQimenTransferorderReportStruct() *TaobaoQimenTransferorderReportStruct {
	return poolTaobaoQimenTransferorderReportStruct.Get().(*TaobaoQimenTransferorderReportStruct)
}

// ReleaseTaobaoQimenTransferorderReportStruct 释放TaobaoQimenTransferorderReportStruct
func ReleaseTaobaoQimenTransferorderReportStruct(v *TaobaoQimenTransferorderReportStruct) {
	v.Items = v.Items[:0]
	v.TransferOrderCode = ""
	v.TransferOutOrderCode = ""
	v.TransferInOrderCode = ""
	v.ConfirmOutTime = ""
	v.ConfirmInTime = ""
	v.CreateTime = ""
	v.FromWarehouseCode = ""
	v.ToWarehouseCode = ""
	v.OwnerCode = ""
	v.ErpOrderCode = ""
	v.OrderStatus = ""
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenTransferorderReportStruct.Put(v)
}
