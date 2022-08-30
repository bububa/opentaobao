package drugtrace

// TopResultModel 结构体
type TopResultModel struct {
	// 返回的批次列表信息
	BlindFileBatchInfoDtoList []BlindFileBatchInfoDto `json:"blind_file_batch_info_dto_list,omitempty" xml:"blind_file_batch_info_dto_list>blind_file_batch_info_dto,omitempty"`
	// 导出的项目和药品目录
	Models []TrialProjectDto `json:"models,omitempty" xml:"models>trial_project_dto,omitempty"`
	// 操作码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 操作说明
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 顶层结构
	Model *AdvanceCodeSearchDto `json:"model,omitempty" xml:"model,omitempty"`
	// 调用成功
	ResponseSuccess bool `json:"response_success,omitempty" xml:"response_success,omitempty"`
}
