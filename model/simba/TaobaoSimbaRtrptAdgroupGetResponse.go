package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取推广组实时报表数据 APIResponse
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
type TaobaoSimbaRtrptAdgroupGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRtrptAdgroupGetResponse
}

type TaobaoSimbaRtrptAdgroupGetResponse struct {
    XMLName xml.Name `xml:"simba_rtrpt_adgroup_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 1111
    
    Results   []RtRptResultEntityDTO `json:"results,omitempty" xml:"results>rt_rpt_result_entity_dto,omitempty"`
    
    
}
