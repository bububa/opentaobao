package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取推广词实时报表数据 APIResponse
taobao.simba.rtrpt.bidword.get

获取推广词报表数据
*/
type TaobaoSimbaRtrptBidwordGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaRtrptBidwordGetResponse `json:"taobao_simba_rtrpt_bidword_get_response,omitempty"`
}

type TaobaoSimbaRtrptBidwordGetResponse struct {

    // bidword result
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
