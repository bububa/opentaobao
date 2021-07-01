package qimen

// TaobaoQimenCombineitemQueryResponse 结构体
type TaobaoQimenCombineitemQueryResponse struct {
	// 响应结果:success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 奇门仓储字段
	CombItems []CombItem `json:"combItems,omitempty" xml:"combItems>comb_item,omitempty"`
}
