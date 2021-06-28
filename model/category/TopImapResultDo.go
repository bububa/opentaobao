package category

// TopImapResultDo 
/* model for simplify = false
type TopImapResultDo struct {

    // true表示调用成功
    
    Success   bool `json:"success,omitempty"`
    

    // 返回的pv对列表
    
    TopPvPairDoList  struct {
        TopPvPairDo  []TopPvPairDo `json:"top_pv_pair_do,omitempty"`
    } `json:"top_pv_pair_do_list,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty"`
    

}
*/

// TopImapResultDo 
type TopImapResultDo struct {

    // true表示调用成功
    Success   bool `json:"success,omitempty"`

    // 返回的pv对列表
    TopPvPairDoList   []TopPvPairDo `json:"top_pv_pair_do_list,omitempty"`

    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty"`

}
