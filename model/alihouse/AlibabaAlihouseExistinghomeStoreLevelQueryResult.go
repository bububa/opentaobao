package alihouse

// AlibabaalihouseexistinghomestorelevelqueryResult 结构体
type AlibabaalihouseexistinghomestorelevelqueryResult struct {
	// 等级结构体
	Data []LevelDtolist `json:"data,omitempty" xml:"data>level_dtolist,omitempty"`
	// 错误消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误编号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
