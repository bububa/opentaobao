package bus

import (
	"sync"
)

// TopStandardInsRecommendResponse 结构体
type TopStandardInsRecommendResponse struct {
	// 推荐结果
	ResultList []TopInsProduct `json:"result_list,omitempty" xml:"result_list>top_ins_product,omitempty"`
}

var poolTopStandardInsRecommendResponse = sync.Pool{
	New: func() any {
		return new(TopStandardInsRecommendResponse)
	},
}

// GetTopStandardInsRecommendResponse() 从对象池中获取TopStandardInsRecommendResponse
func GetTopStandardInsRecommendResponse() *TopStandardInsRecommendResponse {
	return poolTopStandardInsRecommendResponse.Get().(*TopStandardInsRecommendResponse)
}

// ReleaseTopStandardInsRecommendResponse 释放TopStandardInsRecommendResponse
func ReleaseTopStandardInsRecommendResponse(v *TopStandardInsRecommendResponse) {
	v.ResultList = v.ResultList[:0]
	poolTopStandardInsRecommendResponse.Put(v)
}
