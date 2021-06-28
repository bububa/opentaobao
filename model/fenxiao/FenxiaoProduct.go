package fenxiao

// FenxiaoProduct 
/* model for simplify = false
type FenxiaoProduct struct {

    // 产品ID
    
    Pid   int64 `json:"pid,omitempty"`
    

    // 产品名称
    
    Name   string `json:"name,omitempty"`
    

    // 产品线ID
    
    ProductcatId   int64 `json:"productcat_id,omitempty"`
    

    // 采购价格，单位：元。
    
    CostPrice   string `json:"cost_price,omitempty"`
    

    // 最低零售价，单位：分。
    
    RetailPriceLow   string `json:"retail_price_low,omitempty"`
    

    // 最高零售价，单位：分。
    
    RetailPriceHigh   string `json:"retail_price_high,omitempty"`
    

    // 商家编码
    
    OuterId   string `json:"outer_id,omitempty"`
    

    // 产品库存
    
    Quantity   int64 `json:"quantity,omitempty"`
    

    // 产品图片路径
    
    Pictures   string `json:"pictures,omitempty"`
    

    // 产品描述路径，通过http请求获取
    
    DescPath   string `json:"desc_path,omitempty"`
    

    // 所在地：省
    
    Prov   string `json:"prov,omitempty"`
    

    // 所在地：市
    
    City   string `json:"city,omitempty"`
    

    // 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）
    
    PostageType   string `json:"postage_type,omitempty"`
    

    // 运费模板ID
    
    PostageId   int64 `json:"postage_id,omitempty"`
    

    // 平邮费用，单位：元
    
    PostageOrdinary   string `json:"postage_ordinary,omitempty"`
    

    // 快递费用，单位：元
    
    PostageFast   string `json:"postage_fast,omitempty"`
    

    // ems费用，单位：元
    
    PostageEms   string `json:"postage_ems,omitempty"`
    

    // 是否有发票，可选值：false（否）、true（是）
    
    HaveInvoice   bool `json:"have_invoice,omitempty"`
    

    // 是否有保修，可选值：false（否）、true（是）
    
    HaveQuarantee   bool `json:"have_quarantee,omitempty"`
    

    // 发布状态，可选值：up（上架）、down（下架）
    
    Status   string `json:"status,omitempty"`
    

    // 下载人数
    
    ItemsCount   int64 `json:"items_count,omitempty"`
    

    // 累计采购次数
    
    OrdersCount   int64 `json:"orders_count,omitempty"`
    

    // 创建时间
    
    Created   string `json:"created,omitempty"`
    

    // 更新时间
    
    Modified   string `json:"modified,omitempty"`
    

    // sku列表
    
    Skus  struct {
        SkuList  []SkuList `json:"sku_list,omitempty"`
    } `json:"skus,omitempty"`
    

    // 折扣ID（新增参数）
    
    DiscountId   int64 `json:"discount_id,omitempty"`
    

    // 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
    
    TradeType   string `json:"trade_type,omitempty"`
    

    // 采购基准价，单位：元。
    
    StandardPrice   string `json:"standard_price,omitempty"`
    

    // 铺货时间
    
    UpshelfTime   string `json:"upshelf_time,omitempty"`
    

    // 查询产品列表时，查询入参增加is_authz:yes|no yes:需要授权 no:不需要授权
    
    IsAuthz   string `json:"is_authz,omitempty"`
    

    // 类目id
    
    CategoryId   string `json:"category_id,omitempty"`
    

    // 产品属性，格式为pid:vid;pid:vid
    
    Properties   string `json:"properties,omitempty"`
    

    // 属性别名，格式为pid:vid:alias;pid:vid:alias
    
    PropertyAlias   string `json:"property_alias,omitempty"`
    

    // 自定义属性，格式为pid:value;pid:value
    
    InputProperties   string `json:"input_properties,omitempty"`
    

    // 产品描述的内容
    
    Description   string `json:"description,omitempty"`
    

    // 经销采购价
    
    DealerCostPrice   string `json:"dealer_cost_price,omitempty"`
    

    // 产品分销商信息
    
    Pdus  struct {
        PduList  []PduList `json:"pdu_list,omitempty"`
    } `json:"pdus,omitempty"`
    

    // 关联的后端商品id
    
    ScitemId   int64 `json:"scitem_id,omitempty"`
    

    // 零售基准价，单位：元
    
    StandardRetailPrice   string `json:"standard_retail_price,omitempty"`
    

    // 产品图片
    
    Images  struct {
        ProductImageList  []ProductImageList `json:"product_image_list,omitempty"`
    } `json:"images,omitempty"`
    

    // 根据商品ID查询时，返回这个产品对应的商品ID，只对分销商查询接口有效
    
    QueryItemId   int64 `json:"query_item_id,omitempty"`
    

    // 产品spu id
    
    SpuId   int64 `json:"spu_id,omitempty"`
    

    // 导入的商品ID
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 预扣库存
    
    ReservedQuantity   int64 `json:"reserved_quantity,omitempty"`
    

    // 配额可用库存
    
    QuotaQuantity   int64 `json:"quota_quantity,omitempty"`
    

}
*/

// FenxiaoProduct 
type FenxiaoProduct struct {

    // 产品ID
    Pid   int64 `json:"pid,omitempty"`

    // 产品名称
    Name   string `json:"name,omitempty"`

    // 产品线ID
    ProductcatId   int64 `json:"productcat_id,omitempty"`

    // 采购价格，单位：元。
    CostPrice   string `json:"cost_price,omitempty"`

    // 最低零售价，单位：分。
    RetailPriceLow   string `json:"retail_price_low,omitempty"`

    // 最高零售价，单位：分。
    RetailPriceHigh   string `json:"retail_price_high,omitempty"`

    // 商家编码
    OuterId   string `json:"outer_id,omitempty"`

    // 产品库存
    Quantity   int64 `json:"quantity,omitempty"`

    // 产品图片路径
    Pictures   string `json:"pictures,omitempty"`

    // 产品描述路径，通过http请求获取
    DescPath   string `json:"desc_path,omitempty"`

    // 所在地：省
    Prov   string `json:"prov,omitempty"`

    // 所在地：市
    City   string `json:"city,omitempty"`

    // 运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）
    PostageType   string `json:"postage_type,omitempty"`

    // 运费模板ID
    PostageId   int64 `json:"postage_id,omitempty"`

    // 平邮费用，单位：元
    PostageOrdinary   string `json:"postage_ordinary,omitempty"`

    // 快递费用，单位：元
    PostageFast   string `json:"postage_fast,omitempty"`

    // ems费用，单位：元
    PostageEms   string `json:"postage_ems,omitempty"`

    // 是否有发票，可选值：false（否）、true（是）
    HaveInvoice   bool `json:"have_invoice,omitempty"`

    // 是否有保修，可选值：false（否）、true（是）
    HaveQuarantee   bool `json:"have_quarantee,omitempty"`

    // 发布状态，可选值：up（上架）、down（下架）
    Status   string `json:"status,omitempty"`

    // 下载人数
    ItemsCount   int64 `json:"items_count,omitempty"`

    // 累计采购次数
    OrdersCount   int64 `json:"orders_count,omitempty"`

    // 创建时间
    Created   string `json:"created,omitempty"`

    // 更新时间
    Modified   string `json:"modified,omitempty"`

    // sku列表
    Skus   []SkuList `json:"skus,omitempty"`

    // 折扣ID（新增参数）
    DiscountId   int64 `json:"discount_id,omitempty"`

    // 分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）
    TradeType   string `json:"trade_type,omitempty"`

    // 采购基准价，单位：元。
    StandardPrice   string `json:"standard_price,omitempty"`

    // 铺货时间
    UpshelfTime   string `json:"upshelf_time,omitempty"`

    // 查询产品列表时，查询入参增加is_authz:yes|no yes:需要授权 no:不需要授权
    IsAuthz   string `json:"is_authz,omitempty"`

    // 类目id
    CategoryId   string `json:"category_id,omitempty"`

    // 产品属性，格式为pid:vid;pid:vid
    Properties   string `json:"properties,omitempty"`

    // 属性别名，格式为pid:vid:alias;pid:vid:alias
    PropertyAlias   string `json:"property_alias,omitempty"`

    // 自定义属性，格式为pid:value;pid:value
    InputProperties   string `json:"input_properties,omitempty"`

    // 产品描述的内容
    Description   string `json:"description,omitempty"`

    // 经销采购价
    DealerCostPrice   string `json:"dealer_cost_price,omitempty"`

    // 产品分销商信息
    Pdus   []PduList `json:"pdus,omitempty"`

    // 关联的后端商品id
    ScitemId   int64 `json:"scitem_id,omitempty"`

    // 零售基准价，单位：元
    StandardRetailPrice   string `json:"standard_retail_price,omitempty"`

    // 产品图片
    Images   []ProductImageList `json:"images,omitempty"`

    // 根据商品ID查询时，返回这个产品对应的商品ID，只对分销商查询接口有效
    QueryItemId   int64 `json:"query_item_id,omitempty"`

    // 产品spu id
    SpuId   int64 `json:"spu_id,omitempty"`

    // 导入的商品ID
    ItemId   int64 `json:"item_id,omitempty"`

    // 预扣库存
    ReservedQuantity   int64 `json:"reserved_quantity,omitempty"`

    // 配额可用库存
    QuotaQuantity   int64 `json:"quota_quantity,omitempty"`

}
