package qimen

// InventoryRule 
type InventoryRule struct {

    // 奇门仓储字段,C1223,string(50),,
    
    ActionType   string `json:"actionType,omitempty" xml:"actionType,omitempty"`
    

    // 奇门仓储字段,C1223,string(50),,
    
    ItemCode   string `json:"itemCode,omitempty" xml:"itemCode,omitempty"`
    

    // channelRatioRules
    
    ChannelRatioRules   []ChannelRatioRule `json:"channelRatioRules,omitempty" xml:"channelRatioRules,omitempty"`
    

}
