package omniorder

// Classify 
type Classify struct {

    // 分类ID
    
    ClassifyId   int64 `json:"classify_id,omitempty" xml:"classify_id,omitempty"`
    

    // 内部名称
    
    InnerName   string `json:"inner_name,omitempty" xml:"inner_name,omitempty"`
    

    // 展示名称
    
    DisplayName   string `json:"display_name,omitempty" xml:"display_name,omitempty"`
    

    // 关联的门店数
    
    LinkStore   int64 `json:"link_store,omitempty" xml:"link_store,omitempty"`
    

    // 关联的商品数
    
    LinkItem   int64 `json:"link_item,omitempty" xml:"link_item,omitempty"`
    

    // 是否在线上显示
    
    ShowOnline   bool `json:"show_online,omitempty" xml:"show_online,omitempty"`
    

    // 是否在线下显示
    
    ShowOffline   bool `json:"show_offline,omitempty" xml:"show_offline,omitempty"`
    

    // 商户ID
    
    SellerId   int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
    

}
