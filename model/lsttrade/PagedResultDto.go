package lsttrade

import (
	"sync"
)

// PagedResultDto 结构体
type PagedResultDto struct {
	// 返回实体包装类
	ContentList []TopOrderChange2BrandownerDto `json:"content_list,omitempty" xml:"content_list>top_order_change2brandowner_dto,omitempty"`
	// 错误描述
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 总记录条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 页码
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页最大记录数
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 接口调用是否成功 true:调用成功 false:调用失败
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
	v.Page = 0
	v.Size = 0
	v.Success = false
	poolPagedResultDto.Put(v)
}
