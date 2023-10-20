package scs

import (
	"sync"
)

// TaobaoOnebpDkxCampaignCampaignAddResultDto 结构体
type TaobaoOnebpDkxCampaignCampaignAddResultDto struct {
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果
	SolutionResultTopDTO *SolutionResultTopDto `json:"solution_result_top_d_t_o,omitempty" xml:"solution_result_top_d_t_o,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOnebpDkxCampaignCampaignAddResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCampaignCampaignAddResultDto)
	},
}

// GetTaobaoOnebpDkxCampaignCampaignAddResultDto() 从对象池中获取TaobaoOnebpDkxCampaignCampaignAddResultDto
func GetTaobaoOnebpDkxCampaignCampaignAddResultDto() *TaobaoOnebpDkxCampaignCampaignAddResultDto {
	return poolTaobaoOnebpDkxCampaignCampaignAddResultDto.Get().(*TaobaoOnebpDkxCampaignCampaignAddResultDto)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignAddResultDto 释放TaobaoOnebpDkxCampaignCampaignAddResultDto
func ReleaseTaobaoOnebpDkxCampaignCampaignAddResultDto(v *TaobaoOnebpDkxCampaignCampaignAddResultDto) {
	v.Message = ""
	v.SolutionResultTopDTO = nil
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoOnebpDkxCampaignCampaignAddResultDto.Put(v)
}
