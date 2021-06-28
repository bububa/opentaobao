package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
关键词按竞价区间的分布数据 APIResponse
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据
*/
type TaobaoSimbaInsightWordspricedataGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaInsightWordspricedataGetResponse `json:"simba_insight_wordspricedata_get_response,omitempty"` 
    TaobaoSimbaInsightWordspricedataGetResponse
}

/* model for simplify = false
type TaobaoSimbaInsightWordspricedataGetResponse struct {

    // 竞价区间分布数据
    
    WordPricedataList  struct {
        InsightWordPriceDistributeDataDTO  []InsightWordPriceDistributeDataDTO `json:"insight_word_price_distribute_data_dto,omitempty"`
    } `json:"word_pricedata_list,omitempty"`
    

}
*/

type TaobaoSimbaInsightWordspricedataGetResponse struct {

    // 竞价区间分布数据
    WordPricedataList   []InsightWordPriceDistributeDataDTO `json:"word_pricedata_list,omitempty"`

}
