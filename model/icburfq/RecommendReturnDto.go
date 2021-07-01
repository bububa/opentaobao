package icburfq

// RecommendReturnDto 结构体
type RecommendReturnDto struct {
	// 返回结果统计
	Pagination *PageView `json:"pagination,omitempty" xml:"pagination,omitempty"`
	// 返回推荐RFQ
	RfqList []RecommendRfqDto `json:"rfq_list,omitempty" xml:"rfq_list>recommend_rfq_dto,omitempty"`
}
