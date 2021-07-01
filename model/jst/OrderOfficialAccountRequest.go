package jst

// OrderOfficialAccountRequest 结构体
type OrderOfficialAccountRequest struct {
	// 规格，目前只有一个规格“HOV”，表示覆盖 华为/OPPO/vivo
	Spec string `json:"spec,omitempty" xml:"spec,omitempty"`
}
