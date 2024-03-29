package trade

import (
	"sync"
)

// TradeFlowGoodsDetailModel 结构体
type TradeFlowGoodsDetailModel struct {
	// 商品最小销售单位，如：包、盒、袋
	Unit string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 外部系统商品ID
	ExternalGoodsId string `json:"external_goods_id,omitempty" xml:"external_goods_id,omitempty"`
	// 商品标题
	GoodsTitle string `json:"goods_title,omitempty" xml:"goods_title,omitempty"`
	// 商品分类
	Category string `json:"category,omitempty" xml:"category,omitempty"`
	// 商品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 货道剩余商品数量
	RemainingQuantity int64 `json:"remaining_quantity,omitempty" xml:"remaining_quantity,omitempty"`
	// 货道业务类型：1普通；2推广实付金额
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 实付金额(单位:分)
	ActualAmount int64 `json:"actual_amount,omitempty" xml:"actual_amount,omitempty"`
	// 商品数量
	Count int64 `json:"count,omitempty" xml:"count,omitempty"`
	// 货架编码，方向：从上到下，编码：从1开始
	ShelfNo int64 `json:"shelf_no,omitempty" xml:"shelf_no,omitempty"`
	// 交易总额(单位:分)
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// 商品单价(单位:分)
	Price int64 `json:"price,omitempty" xml:"price,omitempty"`
	// 货道编码，方向：从左到右，编码：从1开始
	CargoRoadNo int64 `json:"cargo_road_no,omitempty" xml:"cargo_road_no,omitempty"`
}

var poolTradeFlowGoodsDetailModel = sync.Pool{
	New: func() any {
		return new(TradeFlowGoodsDetailModel)
	},
}

// GetTradeFlowGoodsDetailModel() 从对象池中获取TradeFlowGoodsDetailModel
func GetTradeFlowGoodsDetailModel() *TradeFlowGoodsDetailModel {
	return poolTradeFlowGoodsDetailModel.Get().(*TradeFlowGoodsDetailModel)
}

// ReleaseTradeFlowGoodsDetailModel 释放TradeFlowGoodsDetailModel
func ReleaseTradeFlowGoodsDetailModel(v *TradeFlowGoodsDetailModel) {
	v.Unit = ""
	v.ExternalGoodsId = ""
	v.GoodsTitle = ""
	v.Category = ""
	v.Barcode = ""
	v.RemainingQuantity = 0
	v.BizType = 0
	v.ActualAmount = 0
	v.Count = 0
	v.ShelfNo = 0
	v.TotalAmount = 0
	v.Price = 0
	v.CargoRoadNo = 0
	poolTradeFlowGoodsDetailModel.Put(v)
}
