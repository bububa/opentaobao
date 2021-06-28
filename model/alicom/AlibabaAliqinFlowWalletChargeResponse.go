package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量直充 APIResponse
alibaba.aliqin.flow.wallet.charge

流量直充
*/
type AlibabaAliqinFlowWalletChargeAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_wallet_charge_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 充值请求
    
    Charge   string `json:"charge,omitempty" xml:"