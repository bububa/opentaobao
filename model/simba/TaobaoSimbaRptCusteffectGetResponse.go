package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
用户账户报表效果数据查询（只有汇总数据，无分类数据） APIResponse
taobao.simba.rpt.custeffect.get

用户账户报表效果数据查询（只有汇总数据，无分类数据）
*/
type TaobaoSimbaRptCusteffectGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptCusteffectGetResponse `json:"taobao_simba_rpt_custeffect_get_response,omitempty"`
}

type TaobaoSimbaRptCusteffectGetResponse struct {

    // 账户效果数据返回结果
    RptCustEffectList   string `json:"rpt_cust_effect_list,omitempty"`

}