package omniorder

// TaobaoOmniitemClassifyItemQueryResult 结构体
type TaobaoOmniitemClassifyItemQueryResult struct {
	// 商品ID集合
	Datas []string `json:"datas,omitempty" xml:"datas>string,omitempty"`
	// 提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
