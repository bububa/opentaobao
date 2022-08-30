package alsc

// OrderOpenInfo 结构体
type OrderOpenInfo struct {
	// 行动点模型
	ActionOpenInfos []ActionOpenInfo `json:"action_open_infos,omitempty" xml:"action_open_infos>action_open_info,omitempty"`
	// 主单业务上下文
	BizContext string `json:"biz_context,omitempty" xml:"biz_context,omitempty"`
	// 主单业务状态
	BizStatus string `json:"biz_status,omitempty" xml:"biz_status,omitempty"`
	// 主单状态中文描述
	BizStatusDesc string `json:"biz_status_desc,omitempty" xml:"biz_status_desc,omitempty"`
	// 主单扩展信息
	ExtInfo string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 下单渠道
	OrderChannel string `json:"order_channel,omitempty" xml:"order_channel,omitempty"`
	// 订单详情页跳转链接
	OrderDetailUrl string `json:"order_detail_url,omitempty" xml:"order_detail_url,omitempty"`
	// 下单子渠道
	OrderSubChannel string `json:"order_sub_channel,omitempty" xml:"order_sub_channel,omitempty"`
	// 订单标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// 主单资金明细
	FundOpenInfo *FundOpenInfo `json:"fund_open_info,omitempty" xml:"fund_open_info,omitempty"`
}
