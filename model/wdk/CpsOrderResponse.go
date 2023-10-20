package wdk

// CpsOrderResponse 结构体
type CpsOrderResponse struct {
	// 子单列表
	CpsSubOrder []CpsSubOrderBo `json:"cps_sub_order,omitempty" xml:"cps_sub_order>cps_sub_order_bo,omitempty"`
	// 订单创建时间
	CreateTime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 支付成功时间
	PayTime string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// 业务订单ID
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// 订单状态： INIT(0, &#34;初始化&#34;),CREATED(1, &#34;已创建&#34;),PAID_DONE(2, &#34;已付款&#34;),ACCEPT_ORDER(3, &#34;已接单&#34;),PICK_ORDER(4, &#34;已拣货&#34;),SHIPPED(5, &#34;已发货&#34;),TRADE_SUCCESS(6, &#34;交易成功&#34;),TRADE_CLOSE(-1, &#34;交易关闭&#34;),
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 外部用户ID：由pid+siteid+adzoneid+custom四部分拼成，其中前3个参数对应拼链中淘宝联盟mm=1_2_3，第4个参数对应拼链中 custom_parameters； eg：125000304_108400334_28176750484_12345
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 最后更新时间
	ModifiedTime string `json:"modified_time,omitempty" xml:"modified_time,omitempty"`
	// 主订单付款金额（不含运费、优惠、券）
	PayPrice int64 `json:"pay_price,omitempty" xml:"pay_price,omitempty"`
	// 预估总分佣金额（仅是预估金额，实际结算仍然通过xls线下doublecheck为准）
	ShareAmount int64 `json:"share_amount,omitempty" xml:"share_amount,omitempty"`
}
