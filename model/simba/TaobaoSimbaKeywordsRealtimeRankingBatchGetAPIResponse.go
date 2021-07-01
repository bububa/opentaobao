package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse
获取关键词的新版实时排名 API返回值
taobao.simba.keywords.realtime.ranking.batch.get

根据关键词ID获取关键词的新版实时排名 */
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel
}

// TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel is 获取关键词的新版实时排名 成功返回结果
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_realtime_ranking_batch_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
