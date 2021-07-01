package lstvending

// VendingTradeGoodsDetailDto 结构体
type VendingTradeGoodsDetailDto struct {
	// 修改时间
	GmtModified int64 `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 创建时间
	GmtCreate int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 外部系统交易流水号
	TradeFlowNo string `json:"trade_flow_no,omitempty" xml:"trade_flow_no,omitempty"`
	// 商品分类
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 货道编码
	CargoRoadNo int64 `json:"cargo_road_no,omitempty" xml:"cargo_road_no,omitempty"`
	// 商品总金额
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 商品标题
	GoodsTitle string `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
	// 商品实际金额
	ActualAmount int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	// 货架编码
	ShelfNo int64 `json:"shelf_no,omitempty" xml:"shelf_no,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 商品单价
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 外部系统商品ID
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 商品最小销售单位，如：包、盒、袋
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 货道业务类型：1普通；2推广
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 货道剩余商品数量
	RemainingQuantity int64 `json:"remaining_quantity,omitempty" xml:"remaining_quantity,omitempty"`
	// 商品ID
	GoodsId int64 `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 货道ID
	CargoSpaceId int64 `json:"cargo_space_id,omitempty" xml:"cargo_space_id,omitempty"`
	// 商品清单记录ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}
