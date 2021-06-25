package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家预存余额检查 APIResponse
alibaba.aliqin.flow.wallet.check.balance

检查商家CRM预存余额是否足够进行活动
*/
type AlibabaAliqinFlowWalletCheckBalanceAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAliqinFlowWalletCheckBalanceResponse `json:"alibaba_aliqin_flow_wallet_check_balance_response,omitempty"`
}

type AlibabaAliqinFlowWalletCheckBalanceResponse struct {

    // 余额是否大于校验值，大于返回true，小于返回false
    Value   string `json:"value,omitempty"`

}
