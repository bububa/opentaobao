package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询费用科目信息(限自研商家) APIResponse
taobao.bill.accounts.get

查询费用账户信息
*/
type TaobaoBillAccountsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBillAccountsGetResponse `json:"taobao_bill_accounts_get_response,omitempty"`
}

type TaobaoBillAccountsGetResponse struct {

    // 返回的科目信息
    Accounts   []Account `json:"accounts,omitempty"`

    // 返回记录行数
    TotalResults   int64 `json:"total_results,omitempty"`

}
