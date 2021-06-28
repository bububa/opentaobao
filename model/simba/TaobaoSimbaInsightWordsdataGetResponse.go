package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的大盘数据 APIResponse
taobao.simba.insight.wordsdata.get

获取关键词的详细数据
*/
type TaobaoSimbaInsightWordsdataGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaInsightWordsdataGetResponse `json:"simba_insight_wordsdata_get_response,omitempty"` 
    TaobaoSimbaInsightWordsdataGetResponse
}

/* model for simplify = false
type TaobaoSimbaInsightWordsdataGetResponse struct {

    // 关键词大盘数据列表
    
    WordDataList  struct {
        InsightWordDataDTO  []InsightWordDataDTO `json:"insight_word_data_dto,omitempty"`
    } `json:"word_data_list,omitempty"`
    

}
*/

type TaobaoSimbaInsightWordsdataGetResponse struct {

    // 关键词大盘数据列表
    WordDataList   []InsightWordDataDTO `json:"word_data_list,omitempty"`

}
