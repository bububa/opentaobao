package wdkitem

// UpdateStoreSkuLifeStatusRequestBean 
type UpdateStoreSkuLifeStatusRequestBean struct {
    // 机构编码
    OrgCode   string `json:"org_code,omitempty" xml:"org_code,omitempty"`
    // 商品编码
    SkuCode   string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
    // 门店编码
    StoreId   string `json:"store_id,omitempty" xml:"store_id,omitempty"`
    // 商品状态
    LifeStatus   string `json:"life_status,omitempty" xml:"life_status,omitempty"`
    // 渠道编码
    ShopId   string `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
}
