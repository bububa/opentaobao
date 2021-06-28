package wdk

// WdkOpenMerchantSkuDo 
/* model for simplify = false
type WdkOpenMerchantSkuDo struct {

    // 商家编码
    
    MerchantCode   string `json:"merchant_code,omitempty"`
    

    // 机构编码
    
    OrgainzaNo   string `json:"orgainza_no,omitempty"`
    

    // 商品编码
    
    SkuCode   string `json:"sku_code,omitempty"`
    

    // 商品名称
    
    SkuName   string `json:"sku_name,omitempty"`
    

    // 商品简称
    
    ShortTitle   string `json:"short_title,omitempty"`
    

    // 条码
    
    Barcodes  struct {
        String  []string `json:"string,omitempty"`
    } `json:"barcodes,omitempty"`
    

    // 商品生命周期状态(A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡)
    
    LifeStatus   string `json:"life_status,omitempty"`
    

    // 平台类目编码
    
    BackCatCode   string `json:"back_cat_code,omitempty"`
    

    // 商家类目编码
    
    RetailerCatCode   string `json:"retailer_cat_code,omitempty"`
    

    // 商品经营方式(1001-普通商品， 2001-加工成品，2002-加工半成品，3001-原材料，4001-耗材，6001-组合商品)
    
    ItemType   int64 `json:"item_type,omitempty"`
    

    // 商品销项税率
    
    InvoiceContent   string `json:"invoice_content,omitempty"`
    

    // 是否可售: 1 - 可售， 0 - 不可售
    
    SaleFlag   int64 `json:"sale_flag,omitempty"`
    

    // 税率
    
    TaxRate   string `json:"tax_rate,omitempty"`
    

    // 是否加工商品
    
    HangdlingFlag   int64 `json:"hangdling_flag,omitempty"`
    

    // 存货性质（此字段一经录入不能修改）；此字段可传：原材料、办公品、服务项目、成品、半成品。与是否加工字段组合成商品类型字段。商品类型有5种：耗材、原材料、加工半成品、加工产成品、普通商品。若存货性质是成品，是否加工为是，则商品类型为“加工产成品”；若存货性质是成品，是否加工为否，则商品类型为“普通商品”；若存货性质是半成品，是否加工为是，则商品性质为“加工半成品”
    
    GoodsNature   int64 `json:"goods_nature,omitempty"`
    

    // 建议零售价
    
    SuggestedPrice   int64 `json:"suggested_price,omitempty"`
    

    // 品牌编码
    
    BrandCode   string `json:"brand_code,omitempty"`
    

    // 供应商code
    
    SupplierNo   string `json:"supplier_no,omitempty"`
    

    // 生产商名称
    
    ProducerName   string `json:"producer_name,omitempty"`
    

    // 生产商地址
    
    ProducerAddress   string `json:"producer_address,omitempty"`
    

    // 产品标准号
    
    ProductCode   string `json:"product_code,omitempty"`
    

    // 厂方货号
    
    FactoryNo   string `json:"factory_no,omitempty"`
    

    // 成份
    
    Component   string `json:"component,omitempty"`
    

    // 销售规格描述
    
    SaleSpec   string `json:"sale_spec,omitempty"`
    

    // 售卖单位
    
    SaleUnit   string `json:"sale_unit,omitempty"`
    

    // 净含量
    
    Content   string `json:"content,omitempty"`
    

    // 是否APP可售
    
    AllowAppSale   int64 `json:"allow_app_sale,omitempty"`
    

    // 是否大件
    
    BigFlag   int64 `json:"big_flag,omitempty"`
    

    // 是否称重
    
    WeightFlag   int64 `json:"weight_flag,omitempty"`
    

    // 是否进口
    
    ImportFlag   int64 `json:"import_flag,omitempty"`
    

    // 存储条件；填常温、冷藏、冷冻、热链、鲜活
    
    Storage   string `json:"storage,omitempty"`
    

    // 保质天数
    
    Period   int64 `json:"period,omitempty"`
    

    // 淘鲜达产地库中的值；国内产地传值格式：中国|省|市。若不能确定产地，可以传“见产品外包装”（按商家支持，需要提前通知技术配置）。国外产地只需要传国家名
    
    ProducerPlace   string `json:"producer_place,omitempty"`
    

    // 重量（单位统一为g）。称重品（weight_flag为1）该字段不填。
    
    Weight   int64 `json:"weight,omitempty"`
    

    // 长度(深)
    
    Length   string `json:"length,omitempty"`
    

    // 宽度（宽）
    
    Width   string `json:"width,omitempty"`
    

    // 高度（高）
    
    Height   string `json:"height,omitempty"`
    

    // 主图链接
    
    PicUrl   string `json:"pic_url,omitempty"`
    

    // 商品其他图片
    
    SkuPicUrls   string `json:"sku_pic_urls,omitempty"`
    

    // 商品图文详情
    
    RichTxtTfs   string `json:"rich_txt_tfs,omitempty"`
    

    // 商品卖点
    
    SubTitle   string `json:"sub_title,omitempty"`
    

    // 卖点一名称
    
    Title1   string `json:"title1,omitempty"`
    

    // 卖点一内容
    
    Subtitle1   string `json:"subtitle1,omitempty"`
    

    // 卖点二名称
    
    Title2   string `json:"title2,omitempty"`
    

    // 卖点二内容
    
    Subtitle2   string `json:"subtitle2,omitempty"`
    

    // 创建时间
    
    GmtCreate   string `json:"gmt_create,omitempty"`
    

    // 修改时间
    
    GmtModified   string `json:"gmt_modified,omitempty"`
    

    // 品牌名称
    
    BrandName   string `json:"brand_name,omitempty"`
    

}
*/

// WdkOpenMerchantSkuDo 
type WdkOpenMerchantSkuDo struct {

    // 商家编码
    MerchantCode   string `json:"merchant_code,omitempty"`

    // 机构编码
    OrgainzaNo   string `json:"orgainza_no,omitempty"`

    // 商品编码
    SkuCode   string `json:"sku_code,omitempty"`

    // 商品名称
    SkuName   string `json:"sku_name,omitempty"`

    // 商品简称
    ShortTitle   string `json:"short_title,omitempty"`

    // 条码
    Barcodes   []string `json:"barcodes,omitempty"`

    // 商品生命周期状态(A-正常、T-暂时停购、C-淘汰出清、R-清退、D-删除封挡)
    LifeStatus   string `json:"life_status,omitempty"`

    // 平台类目编码
    BackCatCode   string `json:"back_cat_code,omitempty"`

    // 商家类目编码
    RetailerCatCode   string `json:"retailer_cat_code,omitempty"`

    // 商品经营方式(1001-普通商品， 2001-加工成品，2002-加工半成品，3001-原材料，4001-耗材，6001-组合商品)
    ItemType   int64 `json:"item_type,omitempty"`

    // 商品销项税率
    InvoiceContent   string `json:"invoice_content,omitempty"`

    // 是否可售: 1 - 可售， 0 - 不可售
    SaleFlag   int64 `json:"sale_flag,omitempty"`

    // 税率
    TaxRate   string `json:"tax_rate,omitempty"`

    // 是否加工商品
    HangdlingFlag   int64 `json:"hangdling_flag,omitempty"`

    // 存货性质（此字段一经录入不能修改）；此字段可传：原材料、办公品、服务项目、成品、半成品。与是否加工字段组合成商品类型字段。商品类型有5种：耗材、原材料、加工半成品、加工产成品、普通商品。若存货性质是成品，是否加工为是，则商品类型为“加工产成品”；若存货性质是成品，是否加工为否，则商品类型为“普通商品”；若存货性质是半成品，是否加工为是，则商品性质为“加工半成品”
    GoodsNature   int64 `json:"goods_nature,omitempty"`

    // 建议零售价
    SuggestedPrice   int64 `json:"suggested_price,omitempty"`

    // 品牌编码
    BrandCode   string `json:"brand_code,omitempty"`

    // 供应商code
    SupplierNo   string `json:"supplier_no,omitempty"`

    // 生产商名称
    ProducerName   string `json:"producer_name,omitempty"`

    // 生产商地址
    ProducerAddress   string `json:"producer_address,omitempty"`

    // 产品标准号
    ProductCode   string `json:"product_code,omitempty"`

    // 厂方货号
    FactoryNo   string `json:"factory_no,omitempty"`

    // 成份
    Component   string `json:"component,omitempty"`

    // 销售规格描述
    SaleSpec   string `json:"sale_spec,omitempty"`

    // 售卖单位
    SaleUnit   string `json:"sale_unit,omitempty"`

    // 净含量
    Content   string `json:"content,omitempty"`

    // 是否APP可售
    AllowAppSale   int64 `json:"allow_app_sale,omitempty"`

    // 是否大件
    BigFlag   int64 `json:"big_flag,omitempty"`

    // 是否称重
    WeightFlag   int64 `json:"weight_flag,omitempty"`

    // 是否进口
    ImportFlag   int64 `json:"import_flag,omitempty"`

    // 存储条件；填常温、冷藏、冷冻、热链、鲜活
    Storage   string `json:"storage,omitempty"`

    // 保质天数
    Period   int64 `json:"period,omitempty"`

    // 淘鲜达产地库中的值；国内产地传值格式：中国|省|市。若不能确定产地，可以传“见产品外包装”（按商家支持，需要提前通知技术配置）。国外产地只需要传国家名
    ProducerPlace   string `json:"producer_place,omitempty"`

    // 重量（单位统一为g）。称重品（weight_flag为1）该字段不填。
    Weight   int64 `json:"weight,omitempty"`

    // 长度(深)
    Length   string `json:"length,omitempty"`

    // 宽度（宽）
    Width   string `json:"width,omitempty"`

    // 高度（高）
    Height   string `json:"height,omitempty"`

    // 主图链接
    PicUrl   string `json:"pic_url,omitempty"`

    // 商品其他图片
    SkuPicUrls   string `json:"sku_pic_urls,omitempty"`

    // 商品图文详情
    RichTxtTfs   string `json:"rich_txt_tfs,omitempty"`

    // 商品卖点
    SubTitle   string `json:"sub_title,omitempty"`

    // 卖点一名称
    Title1   string `json:"title1,omitempty"`

    // 卖点一内容
    Subtitle1   string `json:"subtitle1,omitempty"`

    // 卖点二名称
    Title2   string `json:"title2,omitempty"`

    // 卖点二内容
    Subtitle2   string `json:"subtitle2,omitempty"`

    // 创建时间
    GmtCreate   string `json:"gmt_create,omitempty"`

    // 修改时间
    GmtModified   string `json:"gmt_modified,omitempty"`

    // 品牌名称
    BrandName   string `json:"brand_name,omitempty"`

}
