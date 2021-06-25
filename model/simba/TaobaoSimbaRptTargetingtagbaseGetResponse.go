package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
定向基础报表 APIResponse
taobao.simba.rpt.targetingtagbase.get

获取定向基础报表
*/
type TaobaoSimbaRptTargetingtagbaseGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptTargetingtagbaseGetResponse `json:"taobao_simba_rpt_targetingtagbase_get_response,omitempty"`
}

type TaobaoSimbaRptTargetingtagbaseGetResponse struct {

    // result
    Results   []RptBaseEntityDTO `json:"results,omitempty"`

}
