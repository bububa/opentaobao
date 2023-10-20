package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse 获取关键词的新版实时排名 API返回值
// taobao.simba.keywords.realtime.ranking.batch.get
//
// 根据关键词ID获取关键词的新版实时排名
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel is 获取关键词的新版实时排名 成功返回结果
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_realtime_ranking_batch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse
func GetTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse() *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse {
	return poolTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse.Get().(*TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse 将 TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse(v *TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse.Put(v)
}
