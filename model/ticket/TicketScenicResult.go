package ticket

// TicketScenicResult 
/* model for simplify = false
type TicketScenicResult struct {

    // 是否成功
    
    Success   bool `json:"success,omitempty"`
    

    // 商家景点编码
    
    OutScenicId   string `json:"out_scenic_id,omitempty"`
    

    // 阿里标准景点库ID
    
    AliScenicId   int64 `json:"ali_scenic_id,omitempty"`
    

    // 扩展字段，预留
    
    Extend   string `json:"extend,omitempty"`
    

}
*/

// TicketScenicResult 
type TicketScenicResult struct {

    // 是否成功
    Success   bool `json:"success,omitempty"`

    // 商家景点编码
    OutScenicId   string `json:"out_scenic_id,omitempty"`

    // 阿里标准景点库ID
    AliScenicId   int64 `json:"ali_scenic_id,omitempty"`

    // 扩展字段，预留
    Extend   string `json:"extend,omitempty"`

}
