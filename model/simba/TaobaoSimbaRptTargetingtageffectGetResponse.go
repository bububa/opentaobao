package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取定向效果报表数据 APIResponse
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
type TaobaoSimbaRptTargetingtageffectGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptTargetingtageffectGetResponse
}

type TaobaoSimbaRptTargetingtageffectGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_targetingtageffect_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 效果数据
    
    Results   []RptEffectEntityDTO `json:"results,omitempty" xml:"results>rpt_effect_entity_dto,omitempty"`
    
    
}
