package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
分销商合约生产 APIResponse
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
type AlibabaWttOrderContractSubscribeAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaWttOrderContractSubscribeResponse `json:"alibaba_wtt_order_contract_subscribe_response,omitempty"` 
    AlibabaWttOrderContractSubscribeResponse
}

/* model for simplify = false
type AlibabaWttOrderContractSubscribeResponse struct {

    // 合约产生陈宫
    
    Issuccess   bool `json:"issuccess,omitempty"`
    

}
*/

type AlibabaWttOrderContractSubscribeResponse struct {

    // 合约产生陈宫
    Issuccess   bool `json:"issuccess,omitempty"`

}
