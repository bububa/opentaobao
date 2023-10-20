package axintrade

import (
	"sync"
)

// TopOrderDetailApiResDto 结构体
type TopOrderDetailApiResDto struct {
	// 分销商订单号
	OuterOrderId string `json:"outer_order_id,omitempty" xml:"outer_order_id,omitempty"`
	// 订单状态名称
	OrderStatusDesc string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// 创建订单时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 确认码信息（若有）
	ConfirmInfo string `json:"confirm_info,omitempty" xml:"confirm_info,omitempty"`
	// 产品名称
	ProductName string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// 产品描述
	ProductDescription string `json:"product_description,omitempty" xml:"product_description,omitempty"`
	// 出行日期
	ServiceStartTime string `json:"service_start_time,omitempty" xml:"service_start_time,omitempty"`
	// 结束日期
	ServiceEndTime string `json:"service_end_time,omitempty" xml:"service_end_time,omitempty"`
	// 游玩日期（若有）
	UseTime string `json:"use_time,omitempty" xml:"use_time,omitempty"`
	// 采购单号
	PurchaseSubOrderId int64 `json:"purchase_sub_order_id,omitempty" xml:"purchase_sub_order_id,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 产品单价，单位为分
	ProductPrice int64 `json:"product_price,omitempty" xml:"product_price,omitempty"`
	// 订单总金额，单位为分
	TotalPrice int64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// 退款金额，单位为分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 支付ID（若有）
	FundOrderId int64 `json:"fund_order_id,omitempty" xml:"fund_order_id,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
}

var poolTopOrderDetailApiResDto = sync.Pool{
	New: func() any {
		return new(TopOrderDetailApiResDto)
	},
}

// GetTopOrderDetailApiResDto() 从对象池中获取TopOrderDetailApiResDto
func GetTopOrderDetailApiResDto() *TopOrderDetailApiResDto {
	return poolTopOrderDetailApiResDto.Get().(*TopOrderDetailApiResDto)
}

// ReleaseTopOrderDetailApiResDto 释放TopOrderDetailApiResDto
func ReleaseTopOrderDetailApiResDto(v *TopOrderDetailApiResDto) {
	v.OuterOrderId = ""
	v.OrderStatusDesc = ""
	v.GmtCreate = ""
	v.ConfirmInfo = ""
	v.ProductName = ""
	v.ProductDescription = ""
	v.ServiceStartTime = ""
	v.ServiceEndTime = ""
	v.UseTime = ""
	v.PurchaseSubOrderId = 0
	v.OrderStatus = 0
	v.ProductPrice = 0
	v.TotalPrice = 0
	v.RefundFee = 0
	v.FundOrderId = 0
	v.ProductId = 0
	v.BuyAmount = 0
	poolTopOrderDetailApiResDto.Put(v)
}
