package drugtrace

// PiatsCodeFlowResponseDto 结构体
type PiatsCodeFlowResponseDto struct {
	// 头部结构
	Header *Header `json:"header,omitempty" xml:"header,omitempty"`
	// 内容体
	ResponseBody *ResponseBody `json:"response_body,omitempty" xml:"response_body,omitempty"`
}
