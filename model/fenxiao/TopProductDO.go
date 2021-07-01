package fenxiao

// TopProductDo 
type TopProductDo struct {
    // 所在地：市
    City   string `json:"city,omitempty" xml:"city,omitempty"`
    // 代销采购价格，单位：元。
    CostPriceYuan   string `json:"cost_price_yuan,omitempty" xml:"cost_price_yuan,omitempty"`
    // 产品描述路径，通过http请求获取
    DescPath   string `json:"desc_path,omitempty" xml:"desc_path,omitempty"`
    // 创建时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // 更新时间
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 是否有发票，可选值：false（否）、true（是）
    HaveInvoice   int64 `json:"have_invoice,omitempty" xml:"have_invoice,omitempty"`
    // 是否有保修，可选值：false（否）、true（是）
    HaveQuarantee   int64 `json:"have_quarantee,omitempty" xml:"have_quarantee,omitempty"`
    // 关联的前台宝贝id
    AuctionId   int64 `json:"auction_id,omitempty" xml:"auction_id,omitempty"`
    // 下载人数
    ItemsCount   int64 `json:"items_count,omitempty" xml:"items_count,omitempty"`
    // ems费用，单位：元
    PostageEms   string `json:"postage_ems,omitempty" xml:"postage_ems,omitempty"`
    // 快递费用，单位：元
    PostageFast   string `json:"postage_fast,omitempty" xml:"postage_fast,omitempty"`
    // 运费模板ID
    PostageId   int64 `json:"postage_id,omitempty" xml:"postage_id,omitempty"`
    // 平邮费用，单位：元
    PostageOrdinary   string `json:"postage_ordinary,omitempty" xml:"postage_ordinary,omitempty"`
    // 运费类型：1（供应商承担运费）、2（分销商承担运费）可选值：seller（供应商承担运费）、buyer（分销商承担运费）
    PostageType   int64 `json:"postage_type,omitempty" xml:"postage_type,omitempty"`
    // 产品ID
    Pid   int64 `json:"pid,omitempty" xml:"pid,omitempty"`
    // 产品所属产品线id
    ProductLineId   int64 `json:"product_line_id,omitempty" xml:"product_line_id,omitempty"`
    // 商家编码
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // 所在地：省
    Prov   string `json:"prov,omitempty" xml:"prov,omitempty"`
    // 产品库存
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // 最高零售价，单位：元。
    RetailPriceHigh   string `json:"retail_price_high,omitempty" xml:"retail_price_high,omitempty"`
    // 最低零售价，单位：元。
    RetailPriceLow   string `json:"retail_price_low,omitempty" xml:"retail_price_low,omitempty"`
    // scItemId
    ScItemId   int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
    // 分销产品SKU列表
    Skus   []ProductSkuDo `json:"skus,omitempty" xml:"skus>product_sku_do,omitempty"`
    // spuId
    SpuId   int64 `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
    // 市场价：单位元
    StandardPrice   string `json:"standard_price,omitempty" xml:"standard_price,omitempty"`
    // 发布状态： 1 上架，2 下架，3 删除
    Status   int64 `json:"status,omitempty" xml:"status,omitempty"`
    // 产品名称
    Name   string `json:"name,omitempty" xml:"name,omitempty"`
    // 累计采购次数
    OrdersCount   int64 `json:"orders_count,omitempty" xml:"orders_count,omitempty"`
    // 产品图片路径列表，用“,”分隔
    Pictures   string `json:"pictures,omitempty" xml:"pictures,omitempty"`
}
