package idle

// AppraiseOrderSynDto 结构体
type AppraiseOrderSynDto struct {
	// biz_order_id
	BizOrderId string `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	// (主状态,子状态,状态说明)示例如下： (&#34;1&#34;, &#34;1&#34;, &#34;买家拍下未付款&#34;) (&#34;2&#34;, &#34;1&#34;, &#34;买家拍下已付款&#34;) (&#34;3&#34;, &#34;1&#34;, &#34;卖家已发货&#34;) 等，详情参考对接文档
	OrderStatus string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// 见order_status字段说明
	OrderSubStatus string `json:"order_sub_status,omitempty" xml:"order_sub_status,omitempty"`
	// 3.0 必填
	TriggerEvent string `json:"trigger_event,omitempty" xml:"trigger_event,omitempty"`
	// 根据订单状态不同，传递不同的内容
	Attribute *Attribute `json:"attribute,omitempty" xml:"attribute,omitempty"`
}
