package category

// TopImapResultDo 结构体
type TopImapResultDo struct {
	// 返回的pv对列表
	TopPvPairDoList []TopPvpairDo `json:"top_pv_pair_do_list,omitempty" xml:"top_pv_pair_do_list>top_pvpair_do,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// true表示调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
