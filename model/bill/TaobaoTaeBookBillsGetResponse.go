package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tae查询虚拟账户明细数据 APIResponse
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
type TaobaoTaeBookBillsGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaeBookBillsGetResponse `json:"tae_book_bills_get_response,omitempty"` 
    TaobaoTaeBookBillsGetResponse
}

/* model for simplify = false
type TaobaoTaeBookBillsGetResponse struct {

    // 虚拟账户账单列表
    
    Bills  struct {
        TopAcctCashJourDto  []TopAcctCashJourDto `json:"top_acct_cash_jour_dto,omitempty"`
    } `json:"bills,omitempty"`
    

    // 是否有下一页
    
    HasNext   bool `json:"has_next,omitempty"`
    

    // 当前查询的结果数,非总数
    
    TotalResults   int64 `json:"total_results,omitempty"`
    

}
*/

type TaobaoTaeBookBillsGetResponse struct {

    // 虚拟账户账单列表
    Bills   []TopAcctCashJourDto `json:"bills,omitempty"`

    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 当前查询的结果数,非总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
