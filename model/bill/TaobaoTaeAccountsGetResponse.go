package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询费用科目信息 APIResponse
taobao.tae.accounts.get

tae查询费用科目信息
*/
type TaobaoTaeAccountsGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaeAccountsGetResponse
}

type TaobaoTaeAccountsGetResponse struct {
    XMLName xml.Name `xml:"tae_accounts_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的科目信息
    
    Accounts   []TopAccountDto `json:"accounts,omitempty" xml:"accounts>top_account_dto,omitempty"`
    
    
    // 返回记录行数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
