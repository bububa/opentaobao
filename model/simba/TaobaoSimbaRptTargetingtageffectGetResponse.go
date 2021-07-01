package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取定向效果报表数据 API返回值 
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
type TaobaoSimbaRptTargetingtageffectGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaRptTargetingtageffectGetResponse
}

// 获取定向效果报表数据 成功返回结果
type TaobaoSimbaRptTargetingtageffectGetResponse struct {
    XMLName xml.Name `xml:"simba_rpt_targetingtageffect_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 效果数据
    Results   []RptEffectEntityDto `json:"results,omitempty" xml:"results>rpt_effect_entity_dto,omitempty"`
}
