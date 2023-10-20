package simba

// TaobaosimbakeywordsqscoresplitgetResultDto 结构体
type TaobaosimbakeywordsqscoresplitgetResultDto struct {
	// 返回成功/错误码
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回新质量分实体信息
	Result *QscoreSplitDto `json:"result,omitempty" xml:"result,omitempty"`
}
