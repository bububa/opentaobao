package qimen

import (
	"sync"
)

// TaobaoQimenStockoutCreateResponse 结构体
type TaobaoQimenStockoutCreateResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 订单创建时间(YYYY-MM-DD HH:MM:SS)
	CreateTime string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// 出库单仓储系统编码
	DeliveryOrderId string `json:"deliveryOrderId,omitempty" xml:"deliveryOrderId,omitempty"`
}

var poolTaobaoQimenStockoutCreateResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenStockoutCreateResponse)
	},
}

// GetTaobaoQimenStockoutCreateResponse() 从对象池中获取TaobaoQimenStockoutCreateResponse
func GetTaobaoQimenStockoutCreateResponse() *TaobaoQimenStockoutCreateResponse {
	return poolTaobaoQimenStockoutCreateResponse.Get().(*TaobaoQimenStockoutCreateResponse)
}

// ReleaseTaobaoQimenStockoutCreateResponse 释放TaobaoQimenStockoutCreateResponse
func ReleaseTaobaoQimenStockoutCreateResponse(v *TaobaoQimenStockoutCreateResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	v.CreateTime = ""
	v.DeliveryOrderId = ""
	poolTaobaoQimenStockoutCreateResponse.Put(v)
}
