package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tae查询费用科目信息 APIResponse
taobao.tae.accounts.get

tae查询费用科目信息
*/
type TaobaoTaeAccountsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaeAccountsGetResponse `json:"tae_accounts_get_response,omitempty"` 
    TaobaoTaeAccountsGetResponse
}

/* model for simplify = false
type TaobaoTaeAccountsGetResponse struct {

    // 返回的科目信息
    
    Accounts  struct {
        TopAccountDto  []TopAccountDto `json:"top_account_dto,omitempty"`
    } `json:"accounts,omitempty"`
    

    // 返回记录行数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoTaeAccountsGetResponse struct {

    // 返回的科目信息
    Accounts   []TopAccountDto `json:"accounts,omitempty"`

    // 返回记录行数
    TotalResults   int64 `json:"total_results,omitempty"`

}
