package alitripmerchant

// ReplaceRpCode 结构体
type ReplaceRpCode struct {
	// 代替后商品
	ReplaceRpCode string `json:"replace_rp_code,omitempty" xml:"replace_rp_code,omitempty"`
	// 原价商品
	OrigRpCode string `json:"orig_rp_code,omitempty" xml:"orig_rp_code,omitempty"`
}
