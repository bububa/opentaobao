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
    // Response *TaobaoSimbaRtrptBidwordGetResponse `json:"simba_rtrpt_bidword_get_response,omitempty"` 
    TaobaoSimbaRtrptBidwordGetResponse
}

/* model for simplify = false
type TaobaoSimbaRtrptBidwordGetResponse struct {

    // bidword result
    
    Results  struct {
        RtRptResultEntityDTO  []RtRptResultEntityDTO `json:"rt_rpt_result_entity_dto,omitempty"`
    } `json:"results,omitempty"`
    

}
*/

type TaobaoSimbaRtrptBidwordGetResponse struct {

    // bidword result
    Results   []RtRptResultEntityDTO `json:"results,omitempty"`

}
