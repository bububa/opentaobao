package singletreasure

import (
	"sync"
)

// PageResultDto 结构体
type PageResultDto struct {
	// 查询结果
	DataList []ItemDetailInfo `json:"data_list,omitempty" xml:"data_list>item_detail_info,omitempty"`
	// 系统信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 页码
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 系统编码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 返回数量个数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 系统执行成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPageResultDto = sync.Pool{
	New: func() any {
		return new(PageResultDto)
	},
}

// GetPageResultDto() 从对象池中获取PageResultDto
func GetPageResultDto() *PageResultDto {
	return poolPageResultDto.Get().(*PageResultDto)
}

// ReleasePageResultDto 释放PageResultDto
func ReleasePageResultDto(v *PageResultDto) {
	v.DataList = v.DataList[:0]
	v.Message = ""
	v.TotalCount = 0
	v.PageNumber = 0
	v.Code = 0
	v.Size = 0
	v.Success = false
	poolPageResultDto.Put(v)
}
