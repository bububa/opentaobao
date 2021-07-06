package wdk

// BatchStockPublishDto 结构体
type BatchStockPublishDto struct {
	// 子参数列表
	StockPublishDtos []StockPublishDto `json:"stock_publish_dtos,omitempty" xml:"stock_publish_dtos>stock_publish_dto,omitempty"`
	// 外部单据号(幂等)，理解为一次请求提交
	BillNo string `json:"bill_no,omitempty" xml:"bill_no,omitempty"`
	// 发布来源，取商家编码
	PublishSource string `json:"publish_source,omitempty" xml:"publish_source,omitempty"`
	// 仓编码（废弃）
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 经营店编码（必须）
	ShopCode string `json:"shop_code,omitempty" xml:"shop_code,omitempty"`
	// 渠道店编号（废弃）
	ChannelSourceId string `json:"channel_source_id,omitempty" xml:"channel_source_id,omitempty"`
	// 操作者
	Operator string `json:"operator,omitempty" xml:"operator,omitempty"`
	// 更新类型，全量1，增量2
	UpdateType int64 `json:"update_type,omitempty" xml:"update_type,omitempty"`
	// 外部单据类型，9001大润发 9002欧尚 9003三江 9005 文峰 9007美特好
	BillType int64 `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 渠道店类型（废弃）
	ChannelSourceType int64 `json:"channel_source_type,omitempty" xml:"channel_source_type,omitempty"`
	// 是否已扣除未批次数，用于全量发布
	UnBatchedOrderStockSubtracted bool `json:"un_batched_order_stock_subtracted,omitempty" xml:"un_batched_order_stock_subtracted,omitempty"`
}
