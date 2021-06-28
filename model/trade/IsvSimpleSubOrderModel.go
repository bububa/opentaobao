package trade

// IsvSimpleSubOrderModel 
/* model for simplify = false
type IsvSimpleSubOrderModel struct {

    // 1688的商品ID(offerId)
    
    OfferId   string `json:"offer_id,omitempty"`
    

    // 购买数量
    
    Quantity   string `json:"quantity,omitempty"`
    

    // 1688的单品货号ID(skuId)，如果有的话，必须填写
    
    SkuId   string `json:"sku_id,omitempty"`
    

    // 商品的类目(Key)，可不填写
    
    CargoKey   string `json:"cargo_key,omitempty"`
    

}
*/

// IsvSimpleSubOrderModel 
type IsvSimpleSubOrderModel struct {

    // 1688的商品ID(offerId)
    OfferId   string `json:"offer_id,omitempty"`

    // 购买数量
    Quantity   string `json:"quantity,omitempty"`

    // 1688的单品货号ID(skuId)，如果有的话，必须填写
    SkuId   string `json:"sku_id,omitempty"`

    // 商品的类目(Key)，可不填写
    CargoKey   string `json:"cargo_key,omitempty"`

}
