package bill

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoTaeBookBillGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTaeBookBillGetAPIResponseModel).Reset()
}

// TaobaoTaeBookBillGetAPIResponseModel is tae查询单笔虚拟账户明细 成功返回结果
type TaobaoTaeBookBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tae_book_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 虚拟账户账单
	Bookbill *TopAcctCashJourDto `json:"bookbill,omitempty" xml:"bookbill,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTaeBookBillGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Bookbill = nil
}

var poolTaobaoTaeBookBillGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTaeBookBillGetAPIResponse)
	},
}

// GetTaobaoTaeBookBillGetAPIResponse 从 sync.Pool 获取 TaobaoTaeBookBillGetAPIResponse
func GetTaobaoTaeBookBillGetAPIResponse() *TaobaoTaeBookBillGetAPIResponse {
	return poolTaobaoTaeBookBillGetAPIResponse.Get().(*TaobaoTaeBookBillGetAPIResponse)
}

// ReleaseTaobaoTaeBookBillGetAPIResponse 将 TaobaoTaeBookBillGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTaeBookBillGetAPIResponse(v *TaobaoTaeBookBillGetAPIResponse) {
	v.Reset()
	poolTaobaoTaeBookBillGetAPIResponse.Put(v)
}
