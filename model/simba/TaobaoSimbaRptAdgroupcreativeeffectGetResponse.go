package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的创意报表效果数据查询(汇总数据，不分类型) APIResponse
taobao.simba.rpt.adgroupcreativeeffect.get

推广组下的创意报表效果数据查询(汇总数据，不分类型)
*/
type TaobaoSimbaRptAdgroupcreativeeffectGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptAdgroupcreativeeffectGetResponse `json:"simba_rpt_adgroupcreativeeffect_get_response,omitempty"` 
    TaobaoSimbaRptAdgroupcreativeeffectGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptAdgroupcreativeeffectGetResponse struct {

    // 推广组下的创意效果数据列表
    
    RptAdgroupcreativeEffectList   string `json:"rpt_adgroupcreative_effect_list,omitempty"`
    

}
*/

type TaobaoSimbaRptAdgroupcreativeeffectGetResponse struct {

    // 推广组下的创意效果数据列表
    RptAdgroupcreativeEffectList   string `json:"rpt_adgroupcreative_effect_list,omitempty"`

}
