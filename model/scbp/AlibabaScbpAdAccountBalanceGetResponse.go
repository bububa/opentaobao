package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户余额 APIResponse
alibaba.scbp.ad.account.balance.get

查询推广账户余额
*/
type AlibabaScbpAdAccountBalanceGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_account_balance_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Balance   string `json:"balance,omitempty" xml:"