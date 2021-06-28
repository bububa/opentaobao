package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔虚拟账户明细 APIResponse
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细
*/
type TaobaoTaeBookBillGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoTaeBookBillGetResponse `json:"tae_book_bill_get_response,omitempty"` 
    TaobaoTaeBookBillGetResponse
}

/* model for simplify = false
type TaobaoTaeBookBillGetResponse struct {

    // 虚拟账户账单
    
    Bookbill  *struct {
        TopAcctCashJourDto  *TopAcctCashJourDto `json:"top_acct_cash_jour_dto,omitempty"`
    } `json:"bookbill,omitempty"`
    

}
*/

type TaobaoTaeBookBillGetResponse struct {

    // 虚拟账户账单
    Bookbill   *TopAcctCashJourDto `json:"bookbill,omitempty"`

}
