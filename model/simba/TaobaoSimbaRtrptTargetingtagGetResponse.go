package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群实时报表 APIResponse
taobao.simba.rtrpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRtrptTargetingtagGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRtrptTargetingtagGetResponse `json:"simba_rtrpt_targetingtag_get_response,omitempty"` 
    TaobaoSimbaRtrptTargetingtagGetResponse
}

/* model for simplify = false
type TaobaoSimbaRtrptTargetingtagGetResponse struct {

    // 111
    
    Results  struct {
        RtRptResultEntityDTO  []RtRptResultEntityDTO `json:"rt_rpt_result_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaRtrptTargetingtagGetResponse struct {

    // 111
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
