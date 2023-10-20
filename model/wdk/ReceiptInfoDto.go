package wdk

import (
	"sync"
)

// ReceiptInfoDto 结构体
type ReceiptInfoDto struct {
	// 商品列表
	ItemInfoList []ItemInfoDto `json:"item_info_list,omitempty" xml:"item_info_list>item_info_dto,omitempty"`
	// 订单票联  user:顾客联，seller：商家联
	TicketCouPon []string `json:"ticket_cou_pon,omitempty" xml:"ticket_cou_pon>string,omitempty"`
	// 批次名称
	BatchName string `json:"batch_name,omitempty" xml:"batch_name,omitempty"`
	// 淘宝订单号
	TbOrderId string `json:"tb_order_id,omitempty" xml:"tb_order_id,omitempty"`
	// 取消退款金额
	CancelAmount string `json:"cancel_amount,omitempty" xml:"cancel_amount,omitempty"`
	// 渠道号
	OrderNum string `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// 优惠金额
	DiscountAmount string `json:"discount_amount,omitempty" xml:"discount_amount,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 批次号
	BatchId string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 店仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 店仓code
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 自提点地址
	SelfPickPlace string `json:"self_pick_place,omitempty" xml:"self_pick_place,omitempty"`
	// 最晚送达时间
	LatestArriveTime string `json:"latest_arrive_time,omitempty" xml:"latest_arrive_time,omitempty"`
	// 实付金额
	PayOrderAmount string `json:"pay_order_amount,omitempty" xml:"pay_order_amount,omitempty"`
	// 包装费
	PackageFee string `json:"package_fee,omitempty" xml:"package_fee,omitempty"`
	// 差额退款金额
	RefundAmount string `json:"refund_amount,omitempty" xml:"refund_amount,omitempty"`
	// 渠道订单号
	ChannelOrderId string `json:"channel_order_id,omitempty" xml:"channel_order_id,omitempty"`
	// 发票连接
	BillingLink string `json:"billing_link,omitempty" xml:"billing_link,omitempty"`
	// 运费
	Postage string `json:"postage,omitempty" xml:"postage,omitempty"`
	// 商品总额
	TotalOrderAmount string `json:"total_order_amount,omitempty" xml:"total_order_amount,omitempty"`
	// 批次策略
	BatchStrategy string `json:"batch_strategy,omitempty" xml:"batch_strategy,omitempty"`
	// 业务类型
	BusinessType string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 履约单号
	FulfillOrderId string `json:"fulfill_order_id,omitempty" xml:"fulfill_order_id,omitempty"`
	// 储藏方式
	StorageMode int64 `json:"storage_mode,omitempty" xml:"storage_mode,omitempty"`
	// 商品总件数
	AllItemCount int64 `json:"all_item_count,omitempty" xml:"all_item_count,omitempty"`
	// 买家信息
	BuyerInfo *BuyerInfoDto `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
	// 缺货出状态
	OutOfStock bool `json:"out_of_stock,omitempty" xml:"out_of_stock,omitempty"`
}

var poolReceiptInfoDto = sync.Pool{
	New: func() any {
		return new(ReceiptInfoDto)
	},
}

// GetReceiptInfoDto() 从对象池中获取ReceiptInfoDto
func GetReceiptInfoDto() *ReceiptInfoDto {
	return poolReceiptInfoDto.Get().(*ReceiptInfoDto)
}

// ReleaseReceiptInfoDto 释放ReceiptInfoDto
func ReleaseReceiptInfoDto(v *ReceiptInfoDto) {
	v.ItemInfoList = v.ItemInfoList[:0]
	v.TicketCouPon = v.TicketCouPon[:0]
	v.BatchName = ""
	v.TbOrderId = ""
	v.CancelAmount = ""
	v.OrderNum = ""
	v.DiscountAmount = ""
	v.Remark = ""
	v.BatchId = ""
	v.WarehouseName = ""
	v.WarehouseCode = ""
	v.SelfPickPlace = ""
	v.LatestArriveTime = ""
	v.PayOrderAmount = ""
	v.PackageFee = ""
	v.RefundAmount = ""
	v.ChannelOrderId = ""
	v.BillingLink = ""
	v.Postage = ""
	v.TotalOrderAmount = ""
	v.BatchStrategy = ""
	v.BusinessType = ""
	v.FulfillOrderId = ""
	v.StorageMode = 0
	v.AllItemCount = 0
	v.BuyerInfo = nil
	v.OutOfStock = false
	poolReceiptInfoDto.Put(v)
}
