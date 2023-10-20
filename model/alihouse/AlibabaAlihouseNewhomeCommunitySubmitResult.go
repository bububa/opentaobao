package alihouse

// AlibabaalihousenewhomecommunitysubmitResult 结构体
type AlibabaalihousenewhomecommunitysubmitResult struct {
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回素材id
	Data *EbbasCommunitySubmitVo `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
