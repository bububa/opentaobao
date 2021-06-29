package nlife

// LogisticsInfoDetail 
type LogisticsInfoDetail struct {
    // 商品列表：[“货码:数量”]，码可以是条形码(sku级别)也可以是零售加唯一码(货级别)、零售+ itemId+"_"+skuId，唯一码数量一定是1
    GoodsIds   []string `json:"goods_ids,omitempty" xml:"goods_ids>string,omitempty"`
    // 物流单号
    LogisticsNo   string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
    // 物流公司名
    LogisticsCompany   string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
    // 发货时间
    DeliverTime   string `json:"deliver_time,omitempty" xml:"deliver_time,omitempty"`
}
