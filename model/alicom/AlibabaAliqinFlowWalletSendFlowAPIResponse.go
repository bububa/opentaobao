package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放 API返回值 
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowWalletSendFlowAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletSendFlowAPIResponseModel
}

// 流量发放 成功返回结果
type AlibabaAliqinFlowWalletSendFlowAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_send_flow_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // true为成功
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
