package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分销商合约生产 APIResponse
alibaba.wtt.order.contract.subscribe

分销商合约生产
*/
type AlibabaWttOrderContractSubscribeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wtt_order_contract_subscribe_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 合约产生陈宫
    
    Issuccess   bool `json:"issuccess,omitempty" xml:"