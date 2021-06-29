package nlife

// RetailSkuTopDO 
type RetailSkuTopDO struct {
    // outerId
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    // skuId
    SkuId   int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // barcode
    Barcode   string `json:"barcode,omitempty" xml:"barcode,omitempty"`
    // price
    Price   string `json:"price,omitempty" xml:"price,omitempty"`
    // sku属性,多个属性以分号分隔
    Properties   string `json:"properties,omitempty" xml:"properties,omitempty"`
    // sku库存
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // sku属性名称,多个属性以分号分隔
    PropertiesName   string `json:"properties_name,omitempty" xml:"properties_name,omitempty"`
    // 修改时间
    Modified   string `json:"modified,omitempty" xml:"modified,omitempty"`
    // 创建时间
    Created   string `json:"created,omitempty" xml:"created,omitempty"`
    // sku的图片信息
    PicUrl   string `json:"pic_url,omitempty" xml:"pic_url,omitempty"`
    // taobaoItemId
    TaobaoItemId   string `json:"taobao_item_id,omitempty" xml:"taobao_item_id,omitempty"`
    // taobaoSkuId
    TaobaoSkuId   string `json:"taobao_sku_id,omitempty" xml:"taobao_sku_id,omitempty"`
    // sku的挂牌价-单位元,保留2位小数
    TagPrice   string `json:"tag_price,omitempty" xml:"tag_price,omitempty"`
    // sku的真实商品
    SkuRefId   int64 `json:"sku_ref_id,omitempty" xml:"sku_ref_id,omitempty"`
    // 网直供库存
    OnlineQuantity   int64 `json:"online_quantity,omitempty" xml:"online_quantity,omitempty"`
}
