package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取创意实时报表数据 APIResponse
taobao.simba.rtrpt.creative.get

获取创意实时报表数据
*/
type TaobaoSimbaRtrptCreativeGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRtrptCreativeGetResponse `json:"taobao_simba_rtrpt_creative_get_response,omitempty"`
}

type TaobaoSimbaRtrptCreativeGetResponse struct {

    // 111
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
