package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取账户实时报表数据 APIResponse
taobao.simba.rtrpt.cust.get

获取账户实时报表数据
*/
type TaobaoSimbaRtrptCustGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRtrptCustGetResponse `json:"taobao_simba_rtrpt_cust_get_response,omitempty"`
}

type TaobaoSimbaRtrptCustGetResponse struct {

    // 11
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
