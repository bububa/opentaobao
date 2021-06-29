package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的新版实时排名 APIResponse
taobao.simba.keywords.realtime.ranking.batch.get

根据关键词ID获取关键词的新版实时排名
*/
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse
}

type TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse struct {
    XMLName xml.Name `xml:"simba_keywords_realtime_ranking_batch_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值
    
    Result   *TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
