package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按流量细分的数据 APIResponse
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外
*/
type TaobaoSimbaInsightWordssubdataGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaInsightWordssubdataGetResponse `json:"simba_insight_wordssubdata_get_response,omitempty"` 
    TaobaoSimbaInsightWordssubdataGetResponse
}

/* model for simplify = false
type TaobaoSimbaInsightWordssubdataGetResponse struct {

    // 关键词按流量细分的数据
    
    WordSubdataList  struct {
        InsightWordSubDataDTO  []InsightWordSubDataDTO `json:"insight_word_sub_data_dto,omitempty"`
    } `json:"word_subdata_list,omitempty"`
    

}
*/

type TaobaoSimbaInsightWordssubdataGetResponse struct {

    // 关键词按流量细分的数据
    WordSubdataList   []InsightWordSubDataDTO `json:"word_subdata_list,omitempty"`

}
