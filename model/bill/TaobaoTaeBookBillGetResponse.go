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
    Response *TaobaoTaeBookBillGetResponse `json:"taobao_tae_book_bill_get_response,omitempty"`
}

type TaobaoTaeBookBillGetResponse struct {

    // 虚拟账户账单
    Bookbill   *TopAcctCashJourDto `json:"bookbill,omitempty"`

}
