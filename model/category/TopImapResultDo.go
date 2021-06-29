package category

// TopImapResultDO 
type TopImapResultDO struct {
    // true表示调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回的pv对列表
    TopPvPairDoList   []TopPvPairDO `json:"top_pv_pair_do_list,omitempty" xml:"top_pv_pair_do_list>top_pv_pair_do,omitempty"`
    // 错误信息
    ErrorMsg   string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}
