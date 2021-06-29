package ascpchannel

// ProductSkuLinkDTO 
type ProductSkuLinkDTO struct {
    // 分销商商品 skuid
    OutSkuId   string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
    // 供应商产品 skuId
    SkuId   string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
}
