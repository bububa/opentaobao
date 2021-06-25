package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搜索人群离线报表 APIResponse
taobao.simba.rpt.targetingtag.get

获取搜搜人群实时报表
*/
type TaobaoSimbaRptTargetingtagGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRptTargetingtagGetResponse `json:"taobao_simba_rpt_targetingtag_get_response,omitempty"`
}

type TaobaoSimbaRptTargetingtagGetResponse struct {

    // 111
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
