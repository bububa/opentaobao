package qimen

import (
	"sync"
)

// ItemLackReportRequest 结构体
type ItemLackReportRequest struct {
	// 缺货商品列表
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// ERP的发货单编码
	DeliveryOrderCode string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
	// 仓库系统的发货单编码
	DeliveryOrderId string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
	// 缺货回告创建时间(YYYY-MM-DD HH:mm:ss)
	CreateTime string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不 会被重复处理)
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenItemlackReportMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolItemLackReportRequest = sync.Pool{
	New: func() any {
		return new(ItemLackReportRequest)
	},
}

// GetItemLackReportRequest() 从对象池中获取ItemLackReportRequest
func GetItemLackReportRequest() *ItemLackReportRequest {
	return poolItemLackReportRequest.Get().(*ItemLackReportRequest)
}

// ReleaseItemLackReportRequest 释放ItemLackReportRequest
func ReleaseItemLackReportRequest(v *ItemLackReportRequest) {
	v.Items = v.Items[:0]
	v.WarehouseCode = ""
	v.DeliveryOrderCode = ""
	v.DeliveryOrderId = ""
	v.CreateTime = ""
	v.OutBizCode = ""
	v.Remark = ""
	v.ExtendProps = nil
	poolItemLackReportRequest.Put(v)
}
