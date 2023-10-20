package tbk

// TaobaotbkscrelationrefundResult 结构体
type TaobaotbkscrelationrefundResult struct {
	// 第三方应该返还的补贴
	Tksubsidyfeerefund3rdpub string `json:"tk_subsidy_fee_refund3rd_pub,omitempty" xml:"tk_subsidy_fee_refund3rd_pub,omitempty"`
	// 第三方应该返还的佣金
	Tkcommissionfeerefund3rdpub string `json:"tk_commission_fee_refund3rd_pub,omitempty" xml:"tk_commission_fee_refund3rd_pub,omitempty"`
	// 第二方应该返还的补贴(不包括技术服务费)
	Tksubsidyfeerefundpub string `json:"tk_subsidy_fee_refund_pub,omitempty" xml:"tk_subsidy_fee_refund_pub,omitempty"`
	// 第二方应该返还的佣金(不包括技术服务费)
	Tkcommissionfeerefundpub string `json:"tk_commission_fee_refund_pub,omitempty" xml:"tk_commission_fee_refund_pub,omitempty"`
	// 维权完成时间
	Tkrefundsuittime string `json:"tk_refund_suit_time,omitempty" xml:"tk_refund_suit_time,omitempty"`
	// 维权创建时间
	Tkrefundtime string `json:"tk_refund_time,omitempty" xml:"tk_refund_time,omitempty"`
	// 订单结算时间
	Earningtime string `json:"earning_time,omitempty" xml:"earning_time,omitempty"`
	// 订单创建时间
	Tbtradecreatetime string `json:"tb_trade_create_time,omitempty" xml:"tb_trade_create_time,omitempty"`
	// 宝贝标题
	Tbauctiontitle string `json:"tb_auction_title,omitempty" xml:"tb_auction_title,omitempty"`
	// 维权金额
	Refundfee string `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	// 结算金额
	Tbtradefinishprice string `json:"tb_trade_finish_price,omitempty" xml:"tb_trade_finish_price,omitempty"`
	// 应返商家金额(二方)
	Tkpubshowreturnfee string `json:"tk_pub_show_return_fee,omitempty" xml:"tk_pub_show_return_fee,omitempty"`
	// 应返商家金额(三方)
	Tk3rdpubshowreturnfee string `json:"tk3rd_pub_show_return_fee,omitempty" xml:"tk3rd_pub_show_return_fee,omitempty"`
	// （口碑订单）口碑父订单号
	Alscpid string `json:"alsc_pid,omitempty" xml:"alsc_pid,omitempty"`
	// （口碑订单）口碑子订单号
	Alscid string `json:"alsc_id,omitempty" xml:"alsc_id,omitempty"`
	// 更新时间
	Modifiedtime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 淘宝订单编号
	Tbtradeparentid int64 `json:"tb_trade_parent_id,omitempty" xml:"tb_trade_parent_id,omitempty"`
	// 会员关系id
	Specialid int64 `json:"special_id,omitempty" xml:"special_id,omitempty"`
	// 渠道关系id
	Relationid int64 `json:"relation_id,omitempty" xml:"relation_id,omitempty"`
	// 第三方推广者memberid
	Tk3rdpubid int64 `json:"tk3rd_pub_id,omitempty" xml:"tk3rd_pub_id,omitempty"`
	// 推广者memberid
	Tkpubid int64 `json:"tk_pub_id,omitempty" xml:"tk_pub_id,omitempty"`
	// 维权创建(淘客结算回执) 4,维权成功(淘客结算回执) 2,维权失败(淘客结算回执) 3,发生多次维权，待处理      11,从淘客处补扣（钱已结给淘客） 等待扣款 12,从淘客处补扣（钱已结给淘客） 扣款成功 13,从卖家处补扣（钱已结给卖家） 等待扣款 14,从卖家处补扣（钱已结给卖家） 扣款成功 15
	Refundstatus int64 `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// 淘宝子订单编号
	Tbtradeid int64 `json:"tb_trade_id,omitempty" xml:"tb_trade_id,omitempty"`
	// 1 表示2方，2表示3方
	Refundtype int64 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
}
