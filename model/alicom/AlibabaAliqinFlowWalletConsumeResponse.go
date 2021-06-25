package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量扣减 APIResponse
alibaba.aliqin.flow.wallet.consume

流量钱包流量扣减接口
*/
type AlibabaAliqinFlowWalletConsumeAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowWalletConsumeResponse `json:"alibaba_aliqin_flow_wallet_consume_response,omitempty"`
}

type AlibabaAliqinFlowWalletConsumeResponse struct {

    // true为成功
    Value   string `json:"value,omitempty"`

}
