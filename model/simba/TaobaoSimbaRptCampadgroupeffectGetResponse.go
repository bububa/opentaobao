package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型) APIResponse
taobao.simba.rpt.campadgroupeffect.get

推广计划下的推广组报表效果数据查询(只有汇总数据，无分类类型)
*/
type TaobaoSimbaRptCampadgroupeffectGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptCampadgroupeffectGetResponse `json:"taobao_simba_rpt_campadgroupeffect_get_response,omitempty"`
}

type TaobaoSimbaRptCampadgroupeffectGetResponse struct {

    // 推广计划下推广组的效果数据列表
    RptCampadgroupEffectList   string `json:"rpt_campadgroup_effect_list,omitempty"`

}
