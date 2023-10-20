package tbk

import (
	"sync"
)

// TaobaoTbkRelationRefundResult 结构体
type TaobaoTbkRelationRefundResult struct {
	// 第三方应该返还的补贴
	TkSubsidyFeeRefund3rdPub string `json:"tk_subsidy_fee_refund3rd_pub,omitempty" xml:"tk_subsidy_fee_refund3rd_pub,omitempty"`
	// 第三方应该返还的佣金
	TkCommissionFeeRefund3rdPub string `json:"tk_commission_fee_refund3rd_pub,omitempty" xml:"tk_commission_fee_refund3rd_pub,omitempty"`
	// 第二方应该返还的补贴(不包括技术服务费)
	TkSubsidyFeeRefundPub string `json:"tk_subsidy_fee_refund_pub,omitempty" xml:"tk_subsidy_fee_refund_pub,omitempty"`
	// 第二方应该返还的佣金(不包括技术服务费)
	TkCommissionFeeRefundPub string `json:"tk_commission_fee_refund_pub,omitempty" xml:"tk_commission_fee_refund_pub,omitempty"`
	// 维权完成时间
	TkRefundSuitTime string `json:"tk_refund_suit_time,omitempty" xml:"tk_refund_suit_time,omitempty"`
	// 维权创建时间
	TkRefundTime string `json:"tk_refund_time,omitempty" xml:"tk_refund_time,omitempty"`
	// 订单结算时间
	EarningTime string `json:"earning_time,omitempty" xml:"earning_time,omitempty"`
	// 订单创建时间
	TbTradeCreateTime string `json:"tb_trade_create_time,omitempty" xml:"tb_trade_create_time,omitempty"`
	// 宝贝标题
	TbAuctionTitle string `json:"tb_auction_title,omitempty" xml:"tb_auction_title,omitempty"`
	// 维权金额
	RefundFee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 结算金额
	TbTradeFinishPrice string `json:"tb_trade_finish_price,omitempty" xml:"tb_trade_finish_price,omitempty"`
	// 应返商家金额(二方)
	TkPubShowReturnFee string `json:"tk_pub_show_return_fee,omitempty" xml:"tk_pub_show_return_fee,omitempty"`
	// 应返商家金额(三方)
	Tk3rdPubShowReturnFee string `json:"tk3rd_pub_show_return_fee,omitempty" xml:"tk3rd_pub_show_return_fee,omitempty"`
	// （口碑订单）口碑父订单号
	AlscPid string `json:"alsc_pid,omitempty" xml:"alsc_pid,omitempty"`
	// （口碑订单）口碑子订单号
	AlscId string `json:"alsc_id,omitempty" xml:"alsc_id,omitempty"`
	// 更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 淘宝订单编号
	TbTradeParentId int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// 会员关系id
	SpecialId int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道关系id
	RelationId int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 第三方推广者memberid
	Tk3rdPubId int64 `json:"tk3rd_pub_id,omitempty" xml:"tk3rd_pub_id,omitempty"`
	// 推广者memberid
	TkPubId int64 `json:"tk_pub_id,omitempty" xml:"tk_pub_id,omitempty"`
	// 维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理      11,从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15
	RefundStatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 淘宝子订单编号
	TbTradeId int64 `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 1 表示2方，2表示3方
	RefundType int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}

var poolTaobaoTbkRelationRefundResult = sync.Pool{
	New: func() any {
		return new(TaobaoTbkRelationRefundResult)
	},
}

// GetTaobaoTbkRelationRefundResult() 从对象池中获取TaobaoTbkRelationRefundResult
func GetTaobaoTbkRelationRefundResult() *TaobaoTbkRelationRefundResult {
	return poolTaobaoTbkRelationRefundResult.Get().(*TaobaoTbkRelationRefundResult)
}

// ReleaseTaobaoTbkRelationRefundResult 释放TaobaoTbkRelationRefundResult
func ReleaseTaobaoTbkRelationRefundResult(v *TaobaoTbkRelationRefundResult) {
	v.TkSubsidyFeeRefund3rdPub = ""
	v.TkCommissionFeeRefund3rdPub = ""
	v.TkSubsidyFeeRefundPub = ""
	v.TkCommissionFeeRefundPub = ""
	v.TkRefundSuitTime = ""
	v.TkRefundTime = ""
	v.EarningTime = ""
	v.TbTradeCreateTime = ""
	v.TbAuctionTitle = ""
	v.RefundFee = ""
	v.TbTradeFinishPrice = ""
	v.TkPubShowReturnFee = ""
	v.Tk3rdPubShowReturnFee = ""
	v.AlscPid = ""
	v.AlscId = ""
	v.ModifiedTime = ""
	v.TbTradeParentId = 0
	v.SpecialId = 0
	v.RelationId = 0
	v.Tk3rdPubId = 0
	v.TkPubId = 0
	v.RefundStatus = 0
	v.TbTradeId = 0
	v.RefundType = 0
	poolTaobaoTbkRelationRefundResult.Put(v)
}
