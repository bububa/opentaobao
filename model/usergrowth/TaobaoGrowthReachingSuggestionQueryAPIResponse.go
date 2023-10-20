package usergrowth

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoGrowthReachingSuggestionQueryAPIResponse 厂商生态推荐信息查询 API返回值
// taobao.growth.reaching.suggestion.query
//
// 厂商生态推荐信息查询
type TaobaoGrowthReachingSuggestionQueryAPIResponse struct {
	model.CommonResponse
	TaobaoGrowthReachingSuggestionQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingSuggestionQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoGrowthReachingSuggestionQueryAPIResponseModel).Reset()
}

// TaobaoGrowthReachingSuggestionQueryAPIResponseModel is 厂商生态推荐信息查询 成功返回结果
type TaobaoGrowthReachingSuggestionQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"growth_reaching_suggestion_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推荐信息列表
	Suggestions []SuggestionDto `json:"suggestions,omitempty" xml:"suggestions>suggestion_dto,omitempty"`
	// 列表维度曝光上报链接
	ExposureUrl string `json:"exposure_url,omitempty" xml:"exposure_url,omitempty"`
	// 是否参竞
	IsOffering bool `json:"is_offering,omitempty" xml:"is_offering,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoGrowthReachingSuggestionQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Suggestions = m.Suggestions[:0]
	m.ExposureUrl = ""
	m.IsOffering = false
}

var poolTaobaoGrowthReachingSuggestionQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoGrowthReachingSuggestionQueryAPIResponse)
	},
}

// GetTaobaoGrowthReachingSuggestionQueryAPIResponse 从 sync.Pool 获取 TaobaoGrowthReachingSuggestionQueryAPIResponse
func GetTaobaoGrowthReachingSuggestionQueryAPIResponse() *TaobaoGrowthReachingSuggestionQueryAPIResponse {
	return poolTaobaoGrowthReachingSuggestionQueryAPIResponse.Get().(*TaobaoGrowthReachingSuggestionQueryAPIResponse)
}

// ReleaseTaobaoGrowthReachingSuggestionQueryAPIResponse 将 TaobaoGrowthReachingSuggestionQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoGrowthReachingSuggestionQueryAPIResponse(v *TaobaoGrowthReachingSuggestionQueryAPIResponse) {
	v.Reset()
	poolTaobaoGrowthReachingSuggestionQueryAPIResponse.Put(v)
}
