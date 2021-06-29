package alihouse

// EbbasItemDto 
type EbbasItemDto struct {

    // 商品id
    
    ItemId   string `json:"item_id,omitempty" xml:"item_id,omitempty"`
    

    // 楼盘上下架状态
    
    OnlineStatus   *model.File `json:"online_status,omitempty" xml:"online_status,omitempty"`
    

    // 外部id
    
    OuterId   string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
    

}
