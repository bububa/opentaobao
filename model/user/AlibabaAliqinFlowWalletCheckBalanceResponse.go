package user

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家预存余额检查 APIResponse
alibaba.aliqin.flow.wallet.check.balance

检查商家CRM预存余额是否足够进行活动
*/
type AlibabaAliqinFlowWalletCheckBalanceAPIResponse struct {
    model.CommonResponse
    AlibabaAliqinFlowWalletCheckBalanceResponse
}

type AlibabaAliqinFlowWalletCheckBalanceResponse struct {
    XMLName xml.Name `xml:"alibaba_aliqin_flow_wallet_check_balance_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 余额是否大于校验值，大于返回true，小于返回false
    
    Value   string `json:"value,omitempty" xml:"value,omitempty"`

    
}
