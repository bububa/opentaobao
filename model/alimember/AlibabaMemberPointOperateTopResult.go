package alimember

// AlibabaMemberPointOperateTopResult 结构体
type AlibabaMemberPointOperateTopResult struct {
	// 返回码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// json格式
	Data *PointOperateResponseDto `json:"data,omitempty" xml:"data,omitempty"`
}
