package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量平台用户签约情况查询 APIResponse
alibaba.aliqin.flow.wallet.sign

流量平台用户签约情况查询
*/
type AlibabaAliqinFlowWalletSignAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowWalletSignResponse `json:"alibaba_aliqin_flow_wallet_sign_response,omitempty"`
}

type AlibabaAliqinFlowWalletSignResponse struct {

    // 是否成功
    Value   string `json:"value,omitempty"`

}
