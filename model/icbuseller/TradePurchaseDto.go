package icbuseller

import (
	"sync"
)

// TradePurchaseDto 结构体
type TradePurchaseDto struct {
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务开始时间
	ServiceBegin string `json:"service_begin,omitempty" xml:"service_begin,omitempty"`
	// 服务结束时间
	ServiceEnd string `json:"service_end,omitempty" xml:"service_end,omitempty"`
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 服务状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 订单编号
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 购买人登录账号
	BuyerLoginId string `json:"buyer_login_id,omitempty" xml:"buyer_login_id,omitempty"`
	// 服务商id
	VendorId string `json:"vendor_id,omitempty" xml:"vendor_id,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// skucode
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 成交价（无效字段）
	TransactionPrice int64 `json:"transaction_price,omitempty" xml:"transaction_price,omitempty"`
	// 成交单价（无效字段）
	TransactionUnitPrice int64 `json:"transaction_unit_price,omitempty" xml:"transaction_unit_price,omitempty"`
}

var poolTradePurchaseDto = sync.Pool{
	New: func() any {
		return new(TradePurchaseDto)
	},
}

// GetTradePurchaseDto() 从对象池中获取TradePurchaseDto
func GetTradePurchaseDto() *TradePurchaseDto {
	return poolTradePurchaseDto.Get().(*TradePurchaseDto)
}

// ReleaseTradePurchaseDto 释放TradePurchaseDto
func ReleaseTradePurchaseDto(v *TradePurchaseDto) {
	v.ServiceCode = ""
	v.ServiceBegin = ""
	v.ServiceEnd = ""
	v.ServiceType = ""
	v.Status = ""
	v.OrderNo = ""
	v.BuyerLoginId = ""
	v.VendorId = ""
	v.CreateTime = ""
	v.SkuCode = ""
	v.Quantity = 0
	v.TransactionPrice = 0
	v.TransactionUnitPrice = 0
	poolTradePurchaseDto.Put(v)
}
