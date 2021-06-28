package category

// TopImapResultDo 
type TopImapResultDo struct {

    // true表示调用成功
    
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    

    // 返回的pv对列表
    
    TopPvPairDoList   []TopPvPairDo `json:"top_pv_pair_do_list,omitempty" xml:"top_pv_pair_do_list,omitempty"`
    

    // 错误信息
    
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
    

}
