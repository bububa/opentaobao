package lsticitem

import (
	"sync"
)

// PagedResultDto 结构体
type PagedResultDto struct {
	// 商品集合
	ContentList []TopLstItemDto `json:"content_list,omitempty" xml:"content_list>top_lst_item_dto,omitempty"`
	// 失败原因
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 失败code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 查询总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页条数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 当前页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 返回成功与否
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolPagedResultDto = sync.Pool{
	New: func() any {
		return new(PagedResultDto)
	},
}

// GetPagedResultDto() 从对象池中获取PagedResultDto
func GetPagedResultDto() *PagedResultDto {
	return poolPagedResultDto.Get().(*PagedResultDto)
}

// ReleasePagedResultDto 释放PagedResultDto
func ReleasePagedResultDto(v *PagedResultDto) {
	v.ContentList = v.ContentList[:0]
	v.ErrorMessage = ""
	v.ErrorCode = ""
	v.Total = 0
	v.Size = 0
	v.Page = 0
	v.Success = false
	poolPagedResultDto.Put(v)
}
