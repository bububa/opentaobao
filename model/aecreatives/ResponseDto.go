package aecreatives

// ResponseDto 结构体
type ResponseDto struct {
	// 返回结果状态描述
	RespMsg string `json:"resp_msg,omitempty" xml:"resp_msg,omitempty"`
	// 返回结果状态码
	RespCode int64 `json:"resp_code,omitempty" xml:"resp_code,omitempty"`
	// 返回结果状态明细
	Result *TrafficProductResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
