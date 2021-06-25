package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取实时余额，”元”为单位 APIResponse
taobao.simba.account.balance.get

获取实时余额，”元”为单位
*/
type TaobaoSimbaAccountBalanceGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaAccountBalanceGetResponse `json:"taobao_simba_account_balance_get_response,omitempty"`
}

type TaobaoSimbaAccountBalanceGetResponse struct {

    // 实时余额，”元”为单位
    Balance   string `json:"balance,omitempty"`

}
