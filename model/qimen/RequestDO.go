package qimen

// RequestDO 
type RequestDO struct {

    // 奇门仓储字段,C123,string(50),,
    OwnerCode   string `json:"ownerCode,omitempty"`

    // 奇门仓储字段
    WarehouseCodes   []String `json:"warehouseCodes,omitempty"`

    // 奇门仓储字段
    ItemCodes   []String `json:"itemCodes,omitempty"`

    // 奇门仓储字段
    ChannelCodes   []String `json:"channelCodes,omitempty"`

    // 奇门仓储字段,C123,string(50),,
    ItemId   string `json:"itemId,omitempty"`

    // inventoryRules
    InventoryRules   []InventoryRule `json:"inventoryRules,omitempty"`

    // 奇门仓储字段,C123,string(50),必填,
    ActionType   string `json:"actionType,omitempty"`

    // 奇门仓储字段,C123,string(50),必填,
    ShopNick   string `json:"shopNick,omitempty"`

    // 奇门仓储字段,C123,string(50),必填,
    ItemSource   string `json:"itemSource,omitempty"`

    // 奇门仓储字段,C123,string(50),必填,
    ShopItemId   string `json:"shopItemId,omitempty"`

    // 奇门仓储字段,C123,string(50),必填,
    SkuId   string `json:"skuId,omitempty"`

    // 商品编码,S1234,string(50),必填,
    ItemCode   string `json:"itemCode,omitempty"`

}