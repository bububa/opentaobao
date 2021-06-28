package baichuan

// IsvItemSubDo 
/* model for simplify = false
type IsvItemSubDo struct {

    // 商品状态
    
    ItemStatus   *model.File `json:"item_status,omitempty"`
    

    // 商品id
    
    ItemId   int64 `json:"item_id,omitempty"`
    

    // 添加时间
    
    AddTime   string `json:"add_time,omitempty"`
    

    // 数据库索引id
    
    Id   int64 `json:"id,omitempty"`
    

}
*/

// IsvItemSubDo 
type IsvItemSubDo struct {

    // 商品状态
    ItemStatus   *model.File `json:"item_status,omitempty"`

    // 商品id
    ItemId   int64 `json:"item_id,omitempty"`

    // 添加时间
    AddTime   string `json:"add_time,omitempty"`

    // 数据库索引id
    Id   int64 `json:"id,omitempty"`

}
