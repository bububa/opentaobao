package wdk

// ModelList 
type ModelList struct {

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 商品在机构内的生命周期，商品状态；A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡
    LifeStatus   string `json:"life_status,omitempty"`

    // 条码：多个条码用英文逗号分割
    Barcodes   string `json:"barcodes,omitempty"`

    // 价格：单位元
    SalePrice   string `json:"sale_price,omitempty"`

    // 会员价：单位元
    MemberPrice   string `json:"member_price,omitempty"`

    // 售卖单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 是否称重品：1称重品0非称重品
    WeightFlag   string `json:"weight_flag,omitempty"`

    // 商家类目编码
    MerchantCatCode   string `json:"merchant_cat_code,omitempty"`

    // 子公司编码
    OrgNo   string `json:"org_no,omitempty"`

    // 门店编码
    OuCode   string `json:"ou_code,omitempty"`

    // 渠道店编码
    ShopId   string `json:"shop_id,omitempty"`

    // 渠道店类型：4淘宝
    ChannelCodes   string `json:"channel_codes,omitempty"`

    // 税收编码(查询返回使用)
    TaxClassNo   string `json:"tax_class_no,omitempty"`

    // 业务类型 1：盒饭  2：生鲜
    BusinessType   int64 `json:"business_type,omitempty"`

    // 是否测试商品
    TestFlag   int64 `json:"test_flag,omitempty"`

    // 是否服务商品
    ServiceFlag   int64 `json:"service_flag,omitempty"`

    // 修改时间
    ModifiedTime   string `json:"modified_time,omitempty"`

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 默认供应商
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 短标题
    ShortTitle   string `json:"short_title,omitempty"`

    // 是否线上可售 是 0：否 1：是
    OnlineSaleFlag   int64 `json:"online_sale_flag,omitempty"`

    // 销售规格描述
    SaleSpec   string `json:"sale_spec,omitempty"`

    // 加工时间 单位：分钟
    ProcessingTime   int64 `json:"processing_time,omitempty"`

    // 后台平台类目编码
    BackCatCode   string `json:"back_cat_code,omitempty"`

    // 进项税率
    InputTaxRate   string `json:"input_tax_rate,omitempty"`

    // 销项税率
    TaxRate   string `json:"tax_rate,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

    // 品牌编码
    BrandCode   string `json:"brand_code,omitempty"`

    // 保质期天数，商品的保质期（单位：天）,0表示没有保质期
    ShelfLife   string `json:"shelf_life,omitempty"`

}
