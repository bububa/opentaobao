package pur

// MallItemDTO 
type MallItemDTO struct {
    // 币种
    CurrencyCode   string `json:"currency_code,omitempty" xml:"currency_code,omitempty"`
    // 商品预览图
    ImgUrl   string `json:"img_url,omitempty" xml:"img_url,omitempty"`
    // 商品描述
    ItemDescription   string `json:"item_description,omitempty" xml:"item_description,omitempty"`
    // 商品id
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    // 商品名称
    ItemName   string `json:"item_name,omitempty" xml:"item_name,omitempty"`
    // 外部品类id
    MallCategoryId   string `json:"mall_category_id,omitempty" xml:"mall_category_id,omitempty"`
    // 数量
    Quantity   int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
    // skuid
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
    // 外部订单行id
    SubPurReqId   string `json:"sub_pur_req_id,omitempty" xml:"sub_pur_req_id,omitempty"`
    // 供应商id
    SupplierId   string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
    // 税率
    TaxRate   string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
    // 单价
    UnitPrice   string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
    // 计价单位
    Uom   string `json:"uom,omitempty" xml:"uom,omitempty"`
    // 合同单号
    ContractCode   string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
    // 产品详情链接
    ItemUrl   string `json:"item_url,omitempty" xml:"item_url,omitempty"`
}
