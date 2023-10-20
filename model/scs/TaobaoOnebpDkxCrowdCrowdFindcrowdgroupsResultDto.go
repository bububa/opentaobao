package scs

import (
	"sync"
)

// TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto 结构体
type TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto struct {
	// 返回结果
	TemplateGroupTopDTOList []TemplateGroupTopDto `json:"template_group_top_d_t_o_list,omitempty" xml:"template_group_top_d_t_o_list>template_group_top_dto,omitempty"`
	// 返回消息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 结果总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回状态码
	ResultCode *ResultCode `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto)
	},
}

// GetTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto() 从对象池中获取TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto
func GetTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto() *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto {
	return poolTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto.Get().(*TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto 释放TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto
func ReleaseTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto(v *TaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto) {
	v.TemplateGroupTopDTOList = v.TemplateGroupTopDTOList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.ResultCode = nil
	v.Success = false
	poolTaobaoOnebpDkxCrowdCrowdFindcrowdgroupsResultDto.Put(v)
}
