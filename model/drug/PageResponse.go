package drug

import (
	"sync"
)

// PageResponse 结构体
type PageResponse struct {
	// 结果列表
	Spus []TopAlihealthSpuQuery `json:"spus,omitempty" xml:"spus>top_alihealth_spu_query,omitempty"`
	// 当前页码
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

var poolPageResponse = sync.Pool{
	New: func() any {
		return new(PageResponse)
	},
}

// GetPageResponse() 从对象池中获取PageResponse
func GetPageResponse() *PageResponse {
	return poolPageResponse.Get().(*PageResponse)
}

// ReleasePageResponse 释放PageResponse
func ReleasePageResponse(v *PageResponse) {
	v.Spus = v.Spus[:0]
	v.CurrentPage = 0
	v.TotalPage = 0
	v.TotalCount = 0
	poolPageResponse.Put(v)
}
