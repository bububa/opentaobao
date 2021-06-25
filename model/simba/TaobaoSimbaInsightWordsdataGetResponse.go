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
    Response *TaobaoSimbaInsightWordsdataGetResponse `json:"taobao_simba_insight_wordsdata_get_response,omitempty"`
}

type TaobaoSimbaInsightWordsdataGetResponse struct {

    // 关键词大盘数据列表
    WordDataList   []InsightWordDataDTO `json:"word_data_list,omitempty"`

}
