package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量平台用户签约情况查询 APIResponse
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询
*/
type AlibabaAliqinFlowWalletSignAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletSignResponse
}

type AlibabaAliqinFlowWalletSignResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_sign_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`

    
}
