package qimen

import (
	"sync"
)

// TaobaoQimenStockoutConfirmStruct 结构体
type TaobaoQimenStockoutConfirmStruct struct {
	// packages
	Packages []Package `json:"packages,omitempty" xml:"packages>package,omitempty"`
	// orderLines
	OrderLines []OrderLine `json:"orderLines,omitempty" xml:"orderLines>order_line,omitempty"`
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// deliveryOrder
	DeliveryOrder *DeliveryOrder `json:"deliveryOrder,omitempty" xml:"deliveryOrder,omitempty"`
	// 扩展属性
	ExtendProps *TaobaoQimenStockoutConfirmMap `json:"extendProps,omitempty" xml:"extendProps,omitempty"`
}

var poolTaobaoQimenStockoutConfirmStruct = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutConfirmStruct)
	},
}

// GetTaobaoQimenStockoutConfirmStruct() 从对象池中获取TaobaoQimenStockoutConfirmStruct
func GetTaobaoQimenStockoutConfirmStruct() *TaobaoQimenStockoutConfirmStruct {
	return poolTaobaoQimenStockoutConfirmStruct.Get().(*TaobaoQimenStockoutConfirmStruct)
}

// ReleaseTaobaoQimenStockoutConfirmStruct 释放TaobaoQimenStockoutConfirmStruct
func ReleaseTaobaoQimenStockoutConfirmStruct(v *TaobaoQimenStockoutConfirmStruct) {
	v.Packages = v.Packages[:0]
	v.OrderLines = v.OrderLines[:0]
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.DeliveryOrder = nil
	v.ExtendProps = nil
	poolTaobaoQimenStockoutConfirmStruct.Put(v)
}
