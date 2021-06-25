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
    Response *TaobaoSimbaRptTargetingtageffectGetResponse `json:"taobao_simba_rpt_targetingtageffect_get_response,omitempty"`
}

type TaobaoSimbaRptTargetingtageffectGetResponse struct {

    // 效果数据
    Results   []RptEffectEntityDTO `json:"results,omitempty"`

}
