package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询费用科目信息(限自研商家) APIResponse
taobao.bill.accounts.get

查询费用账户信息
*/
type TaobaoBillAccountsGetAPIResponse struct {
    model.CommonResponse
    TaobaoBillAccountsGetResponse
}

type TaobaoBillAccountsGetResponse struct {
    XMLName xml.Name `xml:"bill_accounts_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的科目信息
    
    Accounts   []Account `json:"accounts,omitempty" xml:"accounts>account,omitempty"`
    
    
    // 返回记录行数
    
    TotalResults   int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`

    
}
