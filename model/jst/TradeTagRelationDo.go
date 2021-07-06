package jst

// TradeTagRelationDo 结构体
type TradeTagRelationDo struct {
	// 该标签操作的历史记录
	HistoryTradeTagRelations []HistoryTradeRelationDo `json:"history_trade_tag_relations,omitempty" xml:"history_trade_tag_relations>history_trade_relation_do,omitempty"`
	// 记录的最新修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 记录的创建时间
	GmtCreated string `json:"gmt_created,omitempty" xml:"gmt_created,omitempty"`
	// 标签名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
	// 标签值，json格式
	TagValue string `json:"tag_value,omitempty" xml:"tag_value,omitempty"`
	// 标签类型       1：官方标签      2：自定义标签     3：主站只读标签
	TagType int64 `json:"tag_type,omitempty" xml:"tag_type,omitempty"`
	// 订单id
	Tid int64 `json:"tid,omitempty" xml:"tid,omitempty"`
	// 该标签在消费者端是否显示,0:不显示,1：显示
	Visible int64 `json:"visible,omitempty" xml:"visible,omitempty"`
}
