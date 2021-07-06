package omniorder

// TaobaoOmniitemItemPublishResult 结构体
type TaobaoOmniitemItemPublishResult struct {
	// 错误码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data *ItemLightPublishResult `json:"data,omitempty" xml:"data,omitempty"`
}
