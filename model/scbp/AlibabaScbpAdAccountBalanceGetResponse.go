package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询账户余额 APIResponse
alibaba.scbp.ad.account.balance.get

查询推广账户余额
*/
type AlibabaScbpAdAccountBalanceGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdAccountBalanceGetResponse `json:"alibaba_scbp_ad_account_balance_get_response,omitempty"` 
    AlibabaScbpAdAccountBalanceGetResponse
}

/* model for simplify = false
type AlibabaScbpAdAccountBalanceGetResponse struct {

    // result
    
    Balance   string `json:"balance,omitempty"`
    

}
*/

type AlibabaScbpAdAccountBalanceGetResponse struct {

    // result
    Balance   string `json:"balance,omitempty"`

}
