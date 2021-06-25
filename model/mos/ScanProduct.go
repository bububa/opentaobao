package mos

// ScanProduct 
type ScanProduct struct {

    // 货号
    ArtNo   string `json:"art_no,omitempty"`

    // 条码
    BarCode   string `json:"bar_code,omitempty"`

    // 集团专柜编码
    GroupShopCode   string `json:"group_shop_code,omitempty"`

    // 收银编码
    IntimeCodes   string `json:"intime_codes,omitempty"`

    // 商品名称
    Name   string `json:"name,omitempty"`

    // 销售价
    Price   string `json:"price,omitempty"`

    // 销售属性
    SalePropertys   []SaleProperty `json:"sale_propertys,omitempty"`

    // 专柜Code
    ShopCode   string `json:"shop_code,omitempty"`

    // 商品Id
    SkuId   string `json:"sku_id,omitempty"`

    // 商品来源
    SourceType   string `json:"source_type,omitempty"`

    // 商品来源
    StoreCode   string `json:"store_code,omitempty"`

    // 吊牌价
    TagPrice   string `json:"tag_price,omitempty"`

    // 默认收银编码
    DefaultIntimeCode   string `json:"default_intime_code,omitempty"`

    // 唯一码
    UniqueCode   string `json:"unique_code,omitempty"`

    // 商品标签
    Tag   string `json:"tag,omitempty"`

}
