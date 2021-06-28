package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的新版实时排名 APIResponse
taobao.simba.keywords.realtime.ranking.batch.get

根据关键词ID获取关键词的新版实时排名
*/
type TaobaoSimbaKeywordsRealtimeRankingBatchGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse `json:"simba_keywords_realtime_ranking_batch_get_response,omitempty"` 
    TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse
}

/* model for simplify = false
type TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse struct {

    // 返回值
    
    Result  *struct {
        TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto  *TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto `json:"taobao_simba_keywords_realtime_ranking_batch_get_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type TaobaoSimbaKeywordsRealtimeRankingBatchGetResponse struct {

    // 返回值
    Result   *TaobaoSimbaKeywordsRealtimeRankingBatchGetResultDto `json:"result,omitempty"`

}
