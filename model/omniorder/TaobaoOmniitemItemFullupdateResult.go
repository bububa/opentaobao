package omniorder

// TaobaoOmniitemItemFullupdateResult 结构体
type TaobaoOmniitemItemFullupdateResult struct {
	// 错误码，code=0表示成功
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemLightPublishResult `json:"data,omitempty" xml:"data,omitempty"`
}
