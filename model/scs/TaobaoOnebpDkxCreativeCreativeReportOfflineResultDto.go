package scs

import (
	"sync"
)

// TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto 结构体
type TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto struct {
	// 返回结果
	CreativeBindResultTopDTOList []CreativeBindResultTopDto `json:"creative_bind_result_top_d_t_o_list,omitempty" xml:"creative_bind_result_top_d_t_o_list>creative_bind_result_top_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto)
	},
}

// GetTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto() 从对象池中获取TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto
func GetTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto() *TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto {
	return poolTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto.Get().(*TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto)
}

// ReleaseTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto 释放TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto
func ReleaseTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto(v *TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto) {
	v.CreativeBindResultTopDTOList = v.CreativeBindResultTopDTOList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoOnebpDkxCreativeCreativeReportOfflineResultDto.Put(v)
}
