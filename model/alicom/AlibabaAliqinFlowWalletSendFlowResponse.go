package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量发放 APIResponse
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowWalletSendFlowAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowWalletSendFlowResponse `json:"alibaba_aliqin_flow_wallet_send_flow_response,omitempty"`
}

type AlibabaAliqinFlowWalletSendFlowResponse struct {

    // true为成功
    Value   string `json:"value,omitempty"`

}
