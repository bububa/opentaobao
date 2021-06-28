package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取关键词的大盘数据 APIResponse
taobao.simba.insight.wordsdata.get

获取关键词的详细数据
*/
type TaobaoSimbaInsightWordsdataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightWordsdataGetResponse
}

type TaobaoSimbaInsightWordsdataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_wordsdata_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词大盘数据列表
    
    WordDataList   []InsightWordDataDTO `json:"word_data_list,omitempty" xml:"word_data_list>insight_word_data_dto,omitempty"`
    
    
}
