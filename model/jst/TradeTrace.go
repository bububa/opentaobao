package jst

// TradeTrace 结构体
type TradeTrace struct {
	// 动作发生的时间
	ActionTime string `json:"action_time,omitempty" xml:"action_time,omitempty"`
	// 应用标题
	AppTitle string `json:"app_title,omitempty" xml:"app_title,omitempty"`
	// 子订单的id列表,以逗号分割
	OrderIds string `json:"order_ids,omitempty" xml:"order_ids,omitempty"`
	// 备注字段
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 卖家的淘宝nick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 回流的订单状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 交易tid
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
}
