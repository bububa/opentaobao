package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取定向效果报表数据 APIResponse
taobao.simba.rpt.targetingtageffect.get

获取定向效果报表数据
*/
type TaobaoSimbaRptTargetingtageffectGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptTargetingtageffectGetResponse `json:"simba_rpt_targetingtageffect_get_response,omitempty"` 
    TaobaoSimbaRptTargetingtageffectGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptTargetingtageffectGetResponse struct {

    // 效果数据
    
    Results  struct {
        RptEffectEntityDTO  []RptEffectEntityDTO `json:"rpt_effect_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaRptTargetingtageffectGetResponse struct {

    // 效果数据
    Results   []RptEffectEntityDTO `json:"results,omitempty"`

}
