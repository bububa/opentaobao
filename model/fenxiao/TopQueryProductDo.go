package fenxiao

// TopQueryProductDO 
type TopQueryProductDO struct {
    // 要查询的产品id 列表
    Ids   []int64 `json:"ids,omitempty" xml:"ids>int64,omitempty"`
    // 分页大小
    PageSize   int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
    // 产品线id
    ProductLineId   int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
    // 当前要查看的页数
    PageNum   int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
    // 产品商家编码
    ProductOuterId   string `json:"product_outer_id,omitempty" xml:"product_outer_id,omitempty"`
    // 关联的前端宝贝id列表
    ItemIds   []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
    // sku商家编码
    SkuOuterId   string `json:"sku_outer_id,omitempty" xml:"sku_outer_id,omitempty"`
    // 渠道[21 零售plus]
    Channel   int64 `json:"channel,omitempty" xml:"channel,omitempty"`
}
