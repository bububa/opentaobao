package alihouse

// AlibabaAlihouseExistinghomeStoreLevelQueryResult 结构体
type AlibabaAlihouseExistinghomeStoreLevelQueryResult struct {
	// 等级结构体
	Data []LevelDTOList `json:"data,omitempty" xml:"data>level_dto_list,omitempty"`
	// 错误消息
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 错误编号
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
