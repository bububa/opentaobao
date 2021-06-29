package perfect

// PerfectPerformanceItemPublishReq 
type PerfectPerformanceItemPublishReq struct {
    // 商品描述信息
    DescribeInfo   *PerfectItemDescribeInfoDTO `json:"describe_info,omitempty" xml:"describe_info,omitempty"`
    // 商品基本信息
    ItemBaseInfo   *PerfectItemBaseInfoDTO `json:"item_base_info,omitempty" xml:"item_base_info,omitempty"`
    // 商品sku列表
    ItemSkuInfos   []PerfectItemSkuInfoDTO `json:"item_sku_infos,omitempty" xml:"item_sku_infos>perfect_item_sku_info_dto,omitempty"`
    // 商品交易信息
    ItemTradeInfo   *PerfectItemTradeInfoDTO `json:"item_trade_info,omitempty" xml:"item_trade_info,omitempty"`
    // 商品物流信息
    LogisticsInfo   *PerfectItemLogisticsInfoDTO `json:"logistics_info,omitempty" xml:"logistics_info,omitempty"`
    // 产品信息
    ProductInfo   *PerfectItemProductInfoDTO `json:"product_info,omitempty" xml:"product_info,omitempty"`
    // 商品编码
    ItemCode   string `json:"item_code,omitempty" xml:"item_code,omitempty"`
}
