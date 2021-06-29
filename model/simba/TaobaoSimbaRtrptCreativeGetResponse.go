package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取创意实时报表数据 APIResponse
taobao.simba.rtrpt.creative.get

获取创意实时报表数据
*/
type TaobaoSimbaRtrptCreativeGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRtrptCreativeGetResponse
}

type TaobaoSimbaRtrptCreativeGetResponse struct {
    XMLName xml.Name `xml:"simba_rtrpt_creative_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 111
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
    
    
}
