package jst

// HlUserDo 结构体
type HlUserDo struct {
	// 回流信息是否开通买家端展示
	OpenForBuyer string `json:"open_for_buyer,omitempty" xml:"open_for_buyer,omitempty"`
	// 如果为空，则默认是X_TO_SYSTEM,X_WAIT_ALLOCATION,X_OUT_WAREHOUSE
	OpenNodes string `json:"open_nodes,omitempty" xml:"open_nodes,omitempty"`
}
