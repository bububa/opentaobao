package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量平台用户签约情况查询 API返回值 
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询
*/
type AlibabaAliqinFlowWalletSignAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletSignAPIResponseModel
}

// 流量平台用户签约情况查询 成功返回结果
type AlibabaAliqinFlowWalletSignAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_sign_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 是否成功
    Value   string `json:"value,omitempty" xml:"value,omitempty"`
}
