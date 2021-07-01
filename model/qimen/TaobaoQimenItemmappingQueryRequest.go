package qimen

// TaobaoQimenItemmappingQueryRequest 
type TaobaoQimenItemmappingQueryRequest struct {
    // 奇门仓储字段,C123,string(50),,
    QueryType   string `json:"queryType,omitempty" xml:"queryType,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    ShopItemId   string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    SkuId   string `json:"skuId,omitempty" xml:"skuId,omitempty"`
}
