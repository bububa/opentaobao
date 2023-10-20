package idleisv

import (
	"sync"
)

// AppraiseIsvRefundDto 结构体
type AppraiseIsvRefundDto struct {
	// 买家申请退款原因
	BuyerApplyReason string `json:"buyer_apply_reason,omitempty" xml:"buyer_apply_reason,omitempty"`
	// 买家退款说明,买家申请退款二级原因
	BuyerApplySubReason string `json:"buyer_apply_sub_reason,omitempty" xml:"buyer_apply_sub_reason,omitempty"`
	// 买家昵称（不唯一且用户可以自己更改）
	BuyerNick string `json:"buyer_nick,omitempty" xml:"buyer_nick,omitempty"`
	// 买家申请退款描述
	BuyerApplyDesc string `json:"buyer_apply_desc,omitempty" xml:"buyer_apply_desc,omitempty"`
	// 退货-快递公司
	RefundPostCompany string `json:"refund_post_company,omitempty" xml:"refund_post_company,omitempty"`
	// 退货-快递单号
	RefundPostNo string `json:"refund_post_no,omitempty" xml:"refund_post_no,omitempty"`
	// 卖家同意退货说明
	SellerAgreeMsg string `json:"seller_agree_msg,omitempty" xml:"seller_agree_msg,omitempty"`
	// 卖家昵称（不唯一且用户可以自己更改）
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 卖家拒绝退款说明
	SellerRefuseMsg string `json:"seller_refuse_msg,omitempty" xml:"seller_refuse_msg,omitempty"`
	// 卖家拒绝退款原因
	SellerRefuseReason string `json:"seller_refuse_reason,omitempty" xml:"seller_refuse_reason,omitempty"`
	// 加密的买家id（唯一且不会改变）
	EncryptionBuyerId string `json:"encryption_buyer_id,omitempty" xml:"encryption_buyer_id,omitempty"`
	// 加密的卖家id（唯一且不会改变）
	EncryptionSellerId string `json:"encryption_seller_id,omitempty" xml:"encryption_seller_id,omitempty"`
	// 订单号
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 商品购买数量
	BuyAmount int64 `json:"buy_amount,omitempty" xml:"buy_amount,omitempty"`
	// 货物状态
	GoodsStatus int64 `json:"goods_status,omitempty" xml:"goods_status,omitempty"`
	// 商品信息
	Item *AppraiseIsvItemDto `json:"item,omitempty" xml:"item,omitempty"`
	// 订单状态
	OrderStatus int64 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 订单实付金额,单位分
	Payment int64 `json:"payment,omitempty" xml:"payment,omitempty"`
	// 退款申请时间,时间戳,单位分
	RefundCreateTime int64 `json:"refund_create_time,omitempty" xml:"refund_create_time,omitempty"`
	// 退还金额(退还给买家的金额),单位分
	RefundFee int64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 退款订单号
	RefundId int64 `json:"refund_id,omitempty" xml:"refund_id,omitempty"`
	// 退款订单状态
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 退款最新修改时间,时间戳,单位分
	RefundModifyTime int64 `json:"refund_modify_time,omitempty" xml:"refund_modify_time,omitempty"`
	// 退款超时信息
	TimeoutData *IsvRefundTimeoutDto `json:"timeout_data,omitempty" xml:"timeout_data,omitempty"`
	// 买家是否需要退货
	NeedReturnGoods bool `json:"need_return_goods,omitempty" xml:"need_return_goods,omitempty"`
}

var poolAppraiseIsvRefundDto = sync.Pool{
	New: func() any {
		return new(AppraiseIsvRefundDto)
	},
}

// GetAppraiseIsvRefundDto() 从对象池中获取AppraiseIsvRefundDto
func GetAppraiseIsvRefundDto() *AppraiseIsvRefundDto {
	return poolAppraiseIsvRefundDto.Get().(*AppraiseIsvRefundDto)
}

// ReleaseAppraiseIsvRefundDto 释放AppraiseIsvRefundDto
func ReleaseAppraiseIsvRefundDto(v *AppraiseIsvRefundDto) {
	v.BuyerApplyReason = ""
	v.BuyerApplySubReason = ""
	v.BuyerNick = ""
	v.BuyerApplyDesc = ""
	v.RefundPostCompany = ""
	v.RefundPostNo = ""
	v.SellerAgreeMsg = ""
	v.SellerNick = ""
	v.SellerRefuseMsg = ""
	v.SellerRefuseReason = ""
	v.EncryptionBuyerId = ""
	v.EncryptionSellerId = ""
	v.BizOrderId = 0
	v.BuyAmount = 0
	v.GoodsStatus = 0
	v.Item = nil
	v.OrderStatus = 0
	v.Payment = 0
	v.RefundCreateTime = 0
	v.RefundFee = 0
	v.RefundId = 0
	v.RefundStatus = 0
	v.RefundModifyTime = 0
	v.TimeoutData = nil
	v.NeedReturnGoods = false
	poolAppraiseIsvRefundDto.Put(v)
}
