package homeai

// TaobaoHomeaiAlgPredictResult 结构体
type TaobaoHomeaiAlgPredictResult struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// errormsg
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// data
	Data *FeatureWallSuggestionDto `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
