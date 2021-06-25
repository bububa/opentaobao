package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广组下的词效果报表数据查询(明细数据不分类型查询) APIResponse
taobao.simba.rpt.adgroupkeywordeffect.get

推广组下的词效果报表数据查询(明细数据不分类型查询)
*/
type TaobaoSimbaRptAdgroupkeywordeffectGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptAdgroupkeywordeffectGetResponse `json:"taobao_simba_rpt_adgroupkeywordeffect_get_response,omitempty"`
}

type TaobaoSimbaRptAdgroupkeywordeffectGetResponse struct {

    // 词效果数据返回结果
    RptAdgroupkeywordEffectList   string `json:"rpt_adgroupkeyword_effect_list,omitempty"`

}
