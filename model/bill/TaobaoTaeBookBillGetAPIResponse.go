package bill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTaeBookBillGetAPIResponse tae查询单笔虚拟账户明细 API返回值
// taobao.tae.book.bill.get
//
// tae查询单笔虚拟账户明细
type TaobaoTaeBookBillGetAPIResponse struct {
	model.CommonResponse
	TaobaoTaeBookBillGetAPIResponseModel
}

// TaobaoTaeBookBillGetAPIResponseModel is tae查询单笔虚拟账户明细 成功返回结果
type TaobaoTaeBookBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tae_book_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 虚拟账户账单
	Bookbill *TopAcctCashJourDto `json:"bookbill,omitempty" xml:"bookbill,omitempty"`
}
