package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词按流量细分的数据 APIResponse
taobao.simba.insight.wordssubdata.get

获取关键词按流量进行细分的数据，返回结果中network表示流量的来源，意义如下：1->PC站内，2->PC站外,4->无线站内 5->无线站外
*/
type TaobaoSimbaInsightWordssubdataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightWordssubdataGetResponse
}

type TaobaoSimbaInsightWordssubdataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_wordssubdata_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 关键词按流量细分的数据
    
    WordSubdataList   []InsightWordSubDataDTO `json:"word_subdata_list,omitempty" xml:"word_subdata_list>insight_word_sub_data_dto,omitempty"`
    
    
}
