package icburfq

import (
	"sync"
)

// RecommendReturnDto 结构体
type RecommendReturnDto struct {
	// 返回推荐RFQ
	RfqList []RecommendRfqDto `json:"rfq_list,omitempty" xml:"rfq_list>recommend_rfq_dto,omitempty"`
	// 返回结果统计
	Pagination *PageView `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

var poolRecommendReturnDto = sync.Pool{
	New: func() any {
		return new(RecommendReturnDto)
	},
}

// GetRecommendReturnDto() 从对象池中获取RecommendReturnDto
func GetRecommendReturnDto() *RecommendReturnDto {
	return poolRecommendReturnDto.Get().(*RecommendReturnDto)
}

// ReleaseRecommendReturnDto 释放RecommendReturnDto
func ReleaseRecommendReturnDto(v *RecommendReturnDto) {
	v.RfqList = v.RfqList[:0]
	v.Pagination = nil
	poolRecommendReturnDto.Put(v)
}
