package bill

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
tae查询虚拟账户明细数据 APIResponse
taobao.tae.book.bills.get

tae查询虚拟账户明细数据
*/
type TaobaoTaeBookBillsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tae_book_bills_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 虚拟账户账单列表
    
    Bills   []TopAcctCashJourDto `json:"bills,omitempty" xml:"