package scs

import (
	"sync"
)

// TaobaoOnebpDkxReportReportCrowdListExpandResultDto 结构体
type TaobaoOnebpDkxReportReportCrowdListExpandResultDto struct {
	// 返回结果
	ReportResultTopDTOList []ReportResultTopDto `json:"report_result_top_d_t_o_list,omitempty" xml:"report_result_top_d_t_o_list>report_result_top_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOnebpDkxReportReportCrowdListExpandResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportCrowdListExpandResultDto)
	},
}

// GetTaobaoOnebpDkxReportReportCrowdListExpandResultDto() 从对象池中获取TaobaoOnebpDkxReportReportCrowdListExpandResultDto
func GetTaobaoOnebpDkxReportReportCrowdListExpandResultDto() *TaobaoOnebpDkxReportReportCrowdListExpandResultDto {
	return poolTaobaoOnebpDkxReportReportCrowdListExpandResultDto.Get().(*TaobaoOnebpDkxReportReportCrowdListExpandResultDto)
}

// ReleaseTaobaoOnebpDkxReportReportCrowdListExpandResultDto 释放TaobaoOnebpDkxReportReportCrowdListExpandResultDto
func ReleaseTaobaoOnebpDkxReportReportCrowdListExpandResultDto(v *TaobaoOnebpDkxReportReportCrowdListExpandResultDto) {
	v.ReportResultTopDTOList = v.ReportResultTopDTOList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoOnebpDkxReportReportCrowdListExpandResultDto.Put(v)
}
