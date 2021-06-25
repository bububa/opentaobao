package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取推广组实时报表数据 APIResponse
taobao.simba.rtrpt.adgroup.get

获取推广组实时报表数据
*/
type TaobaoSimbaRtrptAdgroupGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRtrptAdgroupGetResponse `json:"taobao_simba_rtrpt_adgroup_get_response,omitempty"`
}

type TaobaoSimbaRtrptAdgroupGetResponse struct {

    // 1111
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
