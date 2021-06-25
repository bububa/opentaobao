package bill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询虚拟账户明细数据(自研发商家专用) APIResponse
taobao.bill.book.bills.get

查询虚拟账户明细数据
*/
type TaobaoBillBookBillsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBillBookBillsGetResponse `json:"taobao_bill_book_bills_get_response,omitempty"`
}

type TaobaoBillBookBillsGetResponse struct {

    // 是否有下一页
    HasNext   bool `json:"has_next,omitempty"`

    // 虚拟账户账单列表
    Bills   []BookBill `json:"bills,omitempty"`

    // 当前查询的结果数,非总数
    TotalResults   int64 `json:"total_results,omitempty"`

}
