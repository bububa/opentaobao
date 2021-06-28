package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量扣减 APIResponse
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口
*/
type AlibabaAliqinFlowWalletConsumeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_wallet_consume_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true为成功
    
    Value   string `json:"value,omitempty" xml:"