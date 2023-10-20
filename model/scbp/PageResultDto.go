package scbp

import (
	"sync"
)

// PageResultDto 结构体
type PageResultDto struct {
	// 返回数据
	ResultList []AdProductDto `json:"result_list,omitempty" xml:"result_list>ad_product_dto,omitempty"`
	// 错误信息
	Key string `json:"key,omitempty" xml:"key,omitempty"`
	// 总数量
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
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
	v.ResultList = v.ResultList[:0]
	v.Key = ""
	v.TotalCount = 0
	poolPageResultDto.Put(v)
}
