package koubeimall

import (
	"sync"
)

// PageResult 结构体
type PageResult struct {
	// 已授权商圈列表data模型
	MallInfoList []MallDto `json:"mall_info_list,omitempty" xml:"mall_info_list>mall_dto,omitempty"`
	// 门店信息模型
	StoreList []StoreDto `json:"store_list,omitempty" xml:"store_list>store_dto,omitempty"`
	// 下一页起始值
	NextStart int64 `json:"next_start,omitempty" xml:"next_start,omitempty"`
	// 分页长度
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 门店评论信息dto
	CommentInfo *CommentInfoDto `json:"comment_info,omitempty" xml:"comment_info,omitempty"`
	// 是否有更多信息
	HasMore bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
}

var poolPageResult = sync.Pool{
	New: func() any {
		return new(PageResult)
	},
}

// GetPageResult() 从对象池中获取PageResult
func GetPageResult() *PageResult {
	return poolPageResult.Get().(*PageResult)
}

// ReleasePageResult 释放PageResult
func ReleasePageResult(v *PageResult) {
	v.MallInfoList = v.MallInfoList[:0]
	v.StoreList = v.StoreList[:0]
	v.NextStart = 0
	v.PageSize = 0
	v.CommentInfo = nil
	v.HasMore = false
	poolPageResult.Put(v)
}
