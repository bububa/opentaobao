package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组效果报表数据对象 APIResponse
taobao.simba.rpt.adgroupeffect.get

推广组效果报表数据对象
*/
type TaobaoSimbaRptAdgroupeffectGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptAdgroupeffectGetResponse `json:"taobao_simba_rpt_adgroupeffect_get_response,omitempty"`
}

type TaobaoSimbaRptAdgroupeffectGetResponse struct {

    // 推广组效果报表数据对象
    RptAdgroupEffectList   string `json:"rpt_adgroup_effect_list,omitempty"`

}
