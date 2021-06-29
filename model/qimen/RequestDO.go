package qimen

// RequestDO 
type RequestDO struct {
    // 奇门仓储字段,C123,string(50),,
    OwnerCode   string `json:"ownerCode,omitempty" xml:"ownerCode,omitempty"`
    // 奇门仓储字段
    WarehouseCodes   []string `json:"warehouseCodes,omitempty" xml:"warehouseCodes>string,omitempty"`
    // 奇门仓储字段
    ItemCodes   []string `json:"itemCodes,omitempty" xml:"itemCodes>string,omitempty"`
    // 奇门仓储字段
    ChannelCodes   []string `json:"channelCodes,omitempty" xml:"channelCodes>string,omitempty"`
    // 奇门仓储字段,C123,string(50),,
    ItemId   string `json:"itemId,omitempty" xml:"itemId,omitempty"`
    // inventoryRules
    InventoryRules   []InventoryRule `json:"inventoryRules,omitempty" xml:"inventoryRules>inventory_rule,omitempty"`
    // 奇门仓储字段,C123,string(50),必填,
    ActionType   string `json:"actionType,omitempty" xml:"actionType,omitempty"`
    // 奇门仓储字段,C123,string(50),必填,
    ShopNick   string `json:"shopNick,omitempty" xml:"shopNick,omitempty"`
    // 奇门仓储字段,C123,string(50),必填,
    ItemSource   string `json:"itemSource,omitempty" xml:"itemSource,omitempty"`
    // 奇门仓储字段,C123,string(50),必填,
    ShopItemId   string `json:"shopItemId,omitempty" xml:"shopItemId,omitempty"`
    // 奇门仓储字段,C123,string(50),必填,
    SkuId   string `json:"skuId,omitempty" xml:"skuId,omitempty"`
    // 商品编码,S1234,string(50),必填,
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
}
