package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词按竞价区间的分布数据 APIResponse
taobao.simba.insight.wordspricedata.get

获取关键词按竞价区间进行细分的数据
*/
type TaobaoSimbaInsightWordspricedataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightWordspricedataGetResponse
}

type TaobaoSimbaInsightWordspricedataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_wordspricedata_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 竞价区间分布数据
    
    WordPricedataList   []InsightWordPriceDistributeDataDTO `json:"word_pricedata_list,omitempty" xml:"word_pricedata_list>insight_word_price_distribute_data_dto,omitempty"`
    
    
}
