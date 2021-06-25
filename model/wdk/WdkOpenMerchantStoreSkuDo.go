package wdk

// WdkOpenMerchantStoreSkuDo 
type WdkOpenMerchantStoreSkuDo struct {

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 经营店ID
    StoreId   string `json:"store_id,omitempty"`

    // 供应商code
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 商品简称
    ShortTitle   string `json:"short_title,omitempty"`

    // 商品状态；A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡
    LifeStatus   string `json:"life_status,omitempty"`

    // 条码
    Barcode   string `json:"barcode,omitempty"`

    // 商品（渠道）价格
    SkuPrice   int64 `json:"sku_price,omitempty"`

    // 会员价
    MemberPrice   int64 `json:"member_price,omitempty"`

    // 业务类型
    BusinessType   int64 `json:"business_type,omitempty"`

    // APP可售
    OnlineSaleFlag   int64 `json:"online_sale_flag,omitempty"`

    // 渠道类型
    ChannelCode   int64 `json:"channel_code,omitempty"`

    // 线上渠道店code
    ShopId   string `json:"shop_id,omitempty"`

    // 商品销售单位对应的含量表达；APP展示重要字段，体现售卖单位中含有的商品数量，通常描述为“550g/份”等样式。填字符串。
    SaleSpec   string `json:"sale_spec,omitempty"`

    // 加工时间
    ProcessingTime   int64 `json:"processing_time,omitempty"`

    // 是否测试商品
    TestFlag   int64 `json:"test_flag,omitempty"`

    // 是否服务商品
    ServiceFlag   int64 `json:"service_flag,omitempty"`

    // 标准类目编码
    BackCatCode   string `json:"back_cat_code,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

}
