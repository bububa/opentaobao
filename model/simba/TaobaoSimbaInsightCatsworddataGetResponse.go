package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取类目下关键词的数据 APIResponse
taobao.simba.insight.catsworddata.get

获取给定词在给定类目下的详细数据
*/
type TaobaoSimbaInsightCatsworddataGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaInsightCatsworddataGetResponse
}

type TaobaoSimbaInsightCatsworddataGetResponse struct {
    XMLName xml.Name `xml:"simba_insight_catsworddata_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 关键词在类目下的数据
    
    CatwordDataList   []InsightWordDataUnderCatDTO `json:"catword_data_list,omitempty" xml:"catword_data_list>insight_word_data_under_cat_dto,omitempty"`
    
    
}
