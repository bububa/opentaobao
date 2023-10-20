package idleisv

// AlibabaIdleIsvSpuSearchResult 结构体
type AlibabaIdleIsvSpuSearchResult struct {
	// 候选的品牌型号列表
	SpuList []SpuPvdo `json:"spu_list,omitempty" xml:"spu_list>spu_pvdo,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
