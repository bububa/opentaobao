package omniorder

// TaobaoOmniitemItemPublishResult 
type TaobaoOmniitemItemPublishResult struct {
    // 错误码
    Code   string `json:"code,omitempty" xml:"code,omitempty"`
    // data
    Data   *ItemLightPublishResult `json:"data,omitempty" xml:"data,omitempty"`
    // 错误信息
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
