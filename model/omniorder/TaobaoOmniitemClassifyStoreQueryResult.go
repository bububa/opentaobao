package omniorder

// TaobaoOmniitemClassifyStoreQueryResult 结构体
type TaobaoOmniitemClassifyStoreQueryResult struct {
	// 业务数据
	Datas []Classify `json:"datas,omitempty" xml:"datas>classify,omitempty"`
	// 提示信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
}
