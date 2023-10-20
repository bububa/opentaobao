package bill

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBillBookBillsGetAPIResponse 查询虚拟账户明细数据(自研发商家专用) API返回值
// taobao.bill.book.bills.get
//
// 查询虚拟账户明细数据
type TaobaoBillBookBillsGetAPIResponse struct {
	model.CommonResponse
	TaobaoBillBookBillsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBillBookBillsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBillBookBillsGetAPIResponseModel).Reset()
}

// TaobaoBillBookBillsGetAPIResponseModel is 查询虚拟账户明细数据(自研发商家专用) 成功返回结果
type TaobaoBillBookBillsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bill_book_bills_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 虚拟账户账单列表
	Bills []BookBill `json:"bills,omitempty" xml:"bills>book_bill,omitempty"`
	// 当前查询的结果数,非总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBillBookBillsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Bills = m.Bills[:0]
	m.TotalResults = 0
	m.HasNext = false
}

var poolTaobaoBillBookBillsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBillBookBillsGetAPIResponse)
	},
}

// GetTaobaoBillBookBillsGetAPIResponse 从 sync.Pool 获取 TaobaoBillBookBillsGetAPIResponse
func GetTaobaoBillBookBillsGetAPIResponse() *TaobaoBillBookBillsGetAPIResponse {
	return poolTaobaoBillBookBillsGetAPIResponse.Get().(*TaobaoBillBookBillsGetAPIResponse)
}

// ReleaseTaobaoBillBookBillsGetAPIResponse 将 TaobaoBillBookBillsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBillBookBillsGetAPIResponse(v *TaobaoBillBookBillsGetAPIResponse) {
	v.Reset()
	poolTaobaoBillBookBillsGetAPIResponse.Put(v)
}
