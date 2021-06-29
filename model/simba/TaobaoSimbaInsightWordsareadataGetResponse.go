package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按地域进行细分的数据 APIResponse
taobao.simba.insight.wordsareadata.get

获取关键词按地域细分的详细数据，目前地域只能细化到省级别，返回的结果中包括市，是为了方便以后扩展，目前结果中市的值等于省。
*/
type TaobaoSimbaInsightWordsareadataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightWordsareadataGetResponse
}

type TaobaoSimbaInsightWordsareadataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_wordsareadata_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 地域细分数据
    
    WordAreadataList   []InsightWordsAreaDistributeDataDTO `json:"word_areadata_list,omitempty" xml:"word_areadata_list>insight_words_area_distribute_data_dto,omitempty"`
    
    
}
