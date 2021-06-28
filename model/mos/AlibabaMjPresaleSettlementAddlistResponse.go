package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
预售结算数据回传 APIResponse
alibaba.mj.presale.settlement.addlist

用于预售活动结算数据的回传。
*/
type AlibabaMjPresaleSettlementAddlistAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMjPresaleSettlementAddlistResponse `json:"alibaba_mj_presale_settlement_addlist_response,omitempty"` 
    AlibabaMjPresaleSettlementAddlistResponse
}

/* model for simplify = false
type AlibabaMjPresaleSettlementAddlistResponse struct {

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type AlibabaMjPresaleSettlementAddlistResponse struct {

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
