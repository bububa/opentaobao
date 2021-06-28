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
    // Response *TaobaoSimbaRptTargetingtagbaseGetResponse `json:"simba_rpt_targetingtagbase_get_response,omitempty"` 
    TaobaoSimbaRptTargetingtagbaseGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptTargetingtagbaseGetResponse struct {

    // result
    
    Results  struct {
        RptBaseEntityDTO  []RptBaseEntityDTO `json:"rpt_base_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaRptTargetingtagbaseGetResponse struct {

    // result
    Results   []RptBaseEntityDTO `json:"results,omitempty"`

}
