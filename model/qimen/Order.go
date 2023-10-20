package qimen

import (
	"sync"
)

// Order 结构体
type Order struct {
	// 订单包裹信息
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// 单据列表
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 拆单情况
	DeliveryOrders []DeliveryOrder `json:"deliveryOrders,omitempty" xml:"deliveryOrders>delivery_order,omitempty"`
	// 出错的出库单号
	DeliveryOrderCode string `json:"deliveryOrderCode,omitempty" xml:"deliveryOrderCode,omitempty"`
	// 出错信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 单据编号
	OrderCode string `json:"orderCode,omitempty" xml:"orderCode,omitempty"`
	// 仓库编码
	WarehouseCode string `json:"warehouseCode,omitempty" xml:"warehouseCode,omitempty"`
	// 货主编码
	OwnerCode string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
	// 单据类型(JYCK=一般交易出库单;HHCK=换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;B2BRK=B2B入库;DBRK=调拨入库;QTRK=其他入库;XTRK=销退入库;THRK=退货入库;HHRK=换货入库;CNJG=仓内加工单;CGTH=采购退货出库单;)
	OrderType string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	// 外部业务编码(消息ID;用于去重;ISV对于同一请求;分配一个唯一性的编码。用来保证因为网络等原因导致重复传输;请求不会被重复处理;条件必填;条件为一单需要多次确认时)
	OutBizCode string `json:"outBizCode,omitempty" xml:"outBizCode,omitempty"`
	// 仓储系统单据号
	OrderId string `json:"orderId,omitempty" xml:"orderId,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 出库单仓储系统编码
	DeliveryOrderId string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
	// 波次中的次序号
	Num string `json:"num,omitempty" xml:"num,omitempty"`
	// 发货单信息
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 单据详情
	OrderInfo *OrderInfo `json:"orderInfo,omitempty" xml:"orderInfo,omitempty"`
}

var poolOrder = sync.Pool{
	New: func() any {
		return new(Order)
	},
}

// GetOrder() 从对象池中获取Order
func GetOrder() *Order {
	return poolOrder.Get().(*Order)
}

// ReleaseOrder 释放Order
func ReleaseOrder(v *Order) {
	v.Packages = v.Packages[:0]
	v.OrderLines = v.OrderLines[:0]
	v.DeliveryOrders = v.DeliveryOrders[:0]
	v.DeliveryOrderCode = ""
	v.Message = ""
	v.Flag = ""
	v.Code = ""
	v.OrderCode = ""
	v.WarehouseCode = ""
	v.OwnerCode = ""
	v.OrderType = ""
	v.OutBizCode = ""
	v.OrderId = ""
	v.Remark = ""
	v.DeliveryOrderId = ""
	v.Num = ""
	v.DeliveryOrder = nil
	v.OrderInfo = nil
	poolOrder.Put(v)
}
