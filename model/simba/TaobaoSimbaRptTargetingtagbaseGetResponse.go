package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向基础报表 APIResponse
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表
*/
type TaobaoSimbaRptTargetingtagbaseGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptTargetingtagbaseGetResponse
}

type TaobaoSimbaRptTargetingtagbaseGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_targetingtagbase_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Results   []RptBaseEntityDTO `json:"results,omitempty" xml:"results>rpt_base_entity_dto,omitempty"`
    
    
}
