package homeai

// TaobaoHomeaiAlgPredictResult 
type TaobaoHomeaiAlgPredictResult struct {
    // data
    Data   *FeatureWallSuggestionDTO `json:"data,omitempty" xml:"data,omitempty"`
    // code
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // errormsg
    Msg   string `json:"msg,omitempty" xml:"msg,omitempty"`
    // 是否成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
