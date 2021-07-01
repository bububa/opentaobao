package drugtrace

// SaveCodeRelationResultDto 结构体
type SaveCodeRelationResultDto struct {
	// 结果信息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 结果标记
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
}
