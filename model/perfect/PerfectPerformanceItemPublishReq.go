package perfect

import (
	"sync"
)

// PerfectPerformanceItemPublishReq 结构体
type PerfectPerformanceItemPublishReq struct {
	// 商品sku列表
	ItemSkuInfos []PerfectItemSkuInfoDto `json:"item_sku_infos,omitempty" xml:"item_sku_infos>perfect_item_sku_info_dto,omitempty"`
	// 商品编码
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 商品描述信息
	DescribeInfo *PerfectItemDescribeInfoDto `json:"describe_info,omitempty" xml:"describe_info,omitempty"`
	// 商品基本信息
	ItemBaseInfo *PerfectItemBaseInfoDto `json:"item_base_info,omitempty" xml:"item_base_info,omitempty"`
	// 商品交易信息
	ItemTradeInfo *PerfectItemTradeInfoDto `json:"item_trade_info,omitempty" xml:"item_trade_info,omitempty"`
	// 商品物流信息
	LogisticsInfo *PerfectItemLogisticsInfoDto `json:"logistics_info,omitempty" xml:"logistics_info,omitempty"`
	// 产品信息
	ProductInfo *PerfectItemProductInfoDto `json:"product_info,omitempty" xml:"product_info,omitempty"`
}

var poolPerfectPerformanceItemPublishReq = sync.Pool{
	New: func() any {
		return new(PerfectPerformanceItemPublishReq)
	},
}

// GetPerfectPerformanceItemPublishReq() 从对象池中获取PerfectPerformanceItemPublishReq
func GetPerfectPerformanceItemPublishReq() *PerfectPerformanceItemPublishReq {
	return poolPerfectPerformanceItemPublishReq.Get().(*PerfectPerformanceItemPublishReq)
}

// ReleasePerfectPerformanceItemPublishReq 释放PerfectPerformanceItemPublishReq
func ReleasePerfectPerformanceItemPublishReq(v *PerfectPerformanceItemPublishReq) {
	v.ItemSkuInfos = v.ItemSkuInfos[:0]
	v.ItemCode = ""
	v.DescribeInfo = nil
	v.ItemBaseInfo = nil
	v.ItemTradeInfo = nil
	v.LogisticsInfo = nil
	v.ProductInfo = nil
	poolPerfectPerformanceItemPublishReq.Put(v)
}
