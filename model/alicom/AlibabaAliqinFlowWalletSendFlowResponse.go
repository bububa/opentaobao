package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放 APIResponse
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowWalletSendFlowAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_wallet_send_flow_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true为成功
    
    Value   string `json:"value,omitempty" xml:"