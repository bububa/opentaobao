package drugtrace

// NewCodeMovePackagingResultDto 结构体
type NewCodeMovePackagingResultDto struct {
	// 内层小对象
	NewCodeMovePackagingResultDTOs []CodeMovePackagingResultDto `json:"new_code_move_packaging_result_d_t_os,omitempty" xml:"new_code_move_packaging_result_d_t_os>code_move_packaging_result_dto,omitempty"`
	// 二级码
	SecondaryCode string `json:"secondary_code,omitempty" xml:"secondary_code,omitempty"`
	// 一级码
	PrimaryCodes string `json:"primary_codes,omitempty" xml:"primary_codes,omitempty"`
	// 码拼箱状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 码拼箱信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
}
