package drugtrace

import (
	"sync"
)

// PageInfoDto 结构体
type PageInfoDto struct {
	// 返回结果
	ResultList []CodeActiveProcessDto `json:"result_list,omitempty" xml:"result_list>code_active_process_dto,omitempty"`
	// 返回结果对象
	Result []PEntParDto `json:"result,omitempty" xml:"result>p_ent_par_dto,omitempty"`
	// 总条数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 每页几条
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 第几页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

var poolPageInfoDto = sync.Pool{
	New: func() any {
		return new(PageInfoDto)
	},
}

// GetPageInfoDto() 从对象池中获取PageInfoDto
func GetPageInfoDto() *PageInfoDto {
	return poolPageInfoDto.Get().(*PageInfoDto)
}

// ReleasePageInfoDto 释放PageInfoDto
func ReleasePageInfoDto(v *PageInfoDto) {
	v.ResultList = v.ResultList[:0]
	v.Result = v.Result[:0]
	v.TotalNum = 0
	v.PageSize = 0
	v.Page = 0
	poolPageInfoDto.Put(v)
}
