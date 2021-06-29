package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询单笔虚拟账户明细 APIResponse
taobao.tae.book.bill.get

tae查询单笔虚拟账户明细
*/
type TaobaoTaeBookBillGetAPIResponse struct {
    model.CommonResponse
    TaobaoTaeBookBillGetResponse
}

type TaobaoTaeBookBillGetResponse struct {
    XMLName xml.Name `xml:"tae_book_bill_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 虚拟账户账单
    
    Bookbill   *TopAcctCashJourDto `json:"bookbill,omitempty" xml:"bookbill,omitempty"`

    
}
